package service

import (
	"context"
	"encoding/json"
	"fmt"
	goaway "github.com/TwiN/go-away"
	"strings"
	"sync"
	"thichlab-backend-docs/constant"
	"thichlab-backend-docs/dto"
	"thichlab-backend-docs/infrastructure/logger"
	"thichlab-backend-docs/infrastructure/util"
)

//------------------------------------
// Docs:Search
//------------------------------------

func (service *DocsService) IsValidQuerySearch(query string) bool {
	formattedQuery := strings.TrimSpace(query)

	// check query
	if formattedQuery == constant.StringEmpty {
		return true
	}

	// check bad word english
	isProfaneEnglish := goaway.IsProfane(formattedQuery)
	if isProfaneEnglish {
		return true
	}

	// check is domain
	isDomainUrl := util.IsValidDomainUrl(formattedQuery)
	if isDomainUrl {
		return true
	}

	// check bad word vietnamese
	isBadWordVietnamese := util.CheckBadWordVietnamese(formattedQuery)
	if isBadWordVietnamese {
		return true
	}

	return false
}

func (service *DocsService) DocsSearch(ctx context.Context, query string, pageIndex, pageSize int) (*dto.PostsByDocSearchRes, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	if pageSize <= constant.ValueEmpty || pageIndex <= constant.ValueEmpty {
		pageSize = constant.PageSizeDefault
		pageIndex = constant.PageIndexDefault
	}

	if service.IsValidQuerySearch(query) {
		// logger.Error("DocsService:DocsSearch:IsValidQuerySearch:Query - ERROR: %v", query)
		return nil, nil
	}

	var r *dto.PostsByDocSearchRes
	var err error

	var wg sync.WaitGroup
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		r, err = service.Search.SearchDocs(util.ToSearchFormat(query), pageIndex, pageSize)
		if err != nil {
			logger.Error("DocsService:GetPostListByCategory: - ERROR: %v", err)
			return
		}
	}(&wg)

	// Ranking query
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err = service.IndexSearchKeywords(context.Background(), query)
		if err != nil {
			logger.Error("DocsService:GetPostListByCategory:IndexSearchKeywords - ERROR: %v", err)
		}
	}(&wg)
	wg.Wait()

	return r, nil

}

func (service *DocsService) IndexSearchKeywords(ctx context.Context, query string) error {

	err := service.Cache.ZIncrBy(ctx, constant.CacheSearchedKeywordsRankings, 1, util.ToLower(strings.TrimSpace(query)))
	if err != nil {
		logger.Error("DocsService:GetPostListByCategory:IndexSearchKeywords - ERROR: %v", err)
		return err
	}

	return nil
}

func (service *DocsService) GetMostSearchedKeywords(ctx context.Context, pageSize, pageIndex int64) []string {
	return service.Cache.GetHighestScore(ctx, constant.CacheSearchedKeywordsRankings, pageSize, pageIndex)
}

//------------------------------------
// Docs:Public
//------------------------------------

func (service *DocsService) GetCategoryListByHierarchy(ctx context.Context) (*[]dto.CategoryListByHierarchyRes, error) {

	r, err := service.DocsRepository.DBGetCategoryListByHierarchy(ctx)
	if err != nil {
		logger.Error("DocsService:GetChildCategoryList: - ERROR: %v", err)
		return nil, err
	}

	mapCategory := make(map[string][]dto.CategoryListByHierarchy)
	for _, v := range *r {

		if _, ok := mapCategory[v.Parent]; ok {
			mapCategory[v.Parent] = append(mapCategory[v.Parent], v)
		} else {
			mapCategory[v.Parent] = append(mapCategory[v.Parent], v)
		}
	}

	categoryRes := make([]dto.CategoryListByHierarchyRes, 0)
	for k, v := range mapCategory {
		if k != "" {
			item := dto.CategoryListByHierarchyRes{
				Parent: k,
				Child:  v,
			}
			categoryRes = append(categoryRes, item)
		}

	}

	return &categoryRes, nil

}

func (service *DocsService) GetChildCategoryList(ctx context.Context, parentCate string, pageSize, pageIndex int) (*[]dto.GetChildCategoryList, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	if pageSize <= constant.ValueEmpty || pageIndex <= constant.ValueEmpty {
		pageSize = constant.PageSizeDefault
		pageIndex = constant.PageIndexDefault
	}

	r, err := service.DocsRepository.DBGetChildCategoryList(ctx, parentCate, pageSize, pageIndex)
	if err != nil {
		logger.Error("DocsService:GetChildCategoryList: - ERROR: %v", err)
		return nil, err
	}

	return r, nil
}

func (service *DocsService) GetPostListLatestCreated(ctx context.Context, pageSize, pageIndex int) (*[]dto.PostListLatestCreated, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	if pageSize <= constant.ValueEmpty || pageIndex <= constant.ValueEmpty {
		pageSize = constant.PageSizeDefault
		pageIndex = constant.PageIndexDefault
	}

	// TODO: Cần handle clear cache
	keyCache := fmt.Sprintf("%s:%d:%d", constant.CachePostListLatest, pageSize, pageIndex)
	existed := service.Cache.Exists(ctx, keyCache)
	if existed == constant.ValueEmpty {
		r, err := service.DocsRepository.DBGetPostListLatestCreated(ctx, pageSize, pageIndex)
		if err != nil {
			logger.Error("DocsService:GetTagInfoBySlug: - ERROR: %v", err)
			return nil, err
		}

		go service.Cache.Set(context.Background(), keyCache, r, constant.CacheExpiresInOneMonth)
		return r, nil
	} else {
		result := make([]dto.PostListLatestCreated, 0)
		r := service.Cache.Get(ctx, keyCache)
		err := json.Unmarshal([]byte(r), &result)
		if err != nil {
			logger.Error("DocsService:GetTagInfoBySlug: - Unmarshall with ERROR: %v", err)
		}
		return &result, nil
	}

}

func (service *DocsService) GetPostsInSameCategory(ctx context.Context, id string, pageSize, pageIndex int) (*[]dto.PostsInSameCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	if pageSize <= constant.ValueEmpty || pageIndex <= constant.ValueEmpty {
		pageSize = constant.PageSizeDefault
		pageIndex = constant.PageIndexDefault
	}

	keyCache := fmt.Sprintf("%s:%s", constant.CachePostInSameCategory, id)
	existed := service.Cache.Exists(ctx, keyCache)
	if existed == constant.ValueEmpty {

		r, err := service.DocsRepository.DBGetPostsInSameCategory(ctx, id, pageSize, pageIndex)
		if err != nil {
			logger.Error("DocsService:GetTagInfoBySlug: - ERROR: %v", err)
			return nil, err
		}

		go service.Cache.Set(context.Background(), keyCache, r, constant.CacheExpiresInOneMonth)

		return r, nil
	} else {
		result := new([]dto.PostsInSameCategory)
		r := service.Cache.Get(ctx, keyCache)
		err := json.Unmarshal([]byte(r), &result)
		if err != nil {
			logger.Error("DocsService:GetTagInfoBySlug: - Unmarshall with ERROR: %v", err)
		}
		return result, nil
	}

}

func (service *DocsService) GetTagInfoBySlug(ctx context.Context, slug string) (*dto.TagInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	keyCache := fmt.Sprintf("%s:%s", constant.CacheTagInfo, slug)
	existed := service.Cache.Exists(ctx, keyCache)
	if existed == constant.ValueEmpty {
		r, err := service.DocsRepository.DBGetTagInfoBySlug(ctx, slug)
		if err != nil {
			logger.Error("DocsService:GetTagInfoBySlug: - ERROR: %v", err)
			return nil, err
		}

		go service.Cache.Set(context.Background(), keyCache, r, constant.CacheExpiresInOneMonth)
		return r, nil
	} else {
		result := new(dto.TagInfo)
		r := service.Cache.Get(ctx, keyCache)
		err := json.Unmarshal([]byte(r), &result)
		if err != nil {
			logger.Error("DocsService:GetTagInfoBySlug: - Unmarshall with ERROR: %v", err)
		}
		return result, nil
	}

}

func (service *DocsService) GetInfoCategoryBySlug(ctx context.Context, parentCate, childCate string) (*dto.CategoryInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	// Cache Parent category
	if parentCate != constant.StringEmpty && childCate == constant.StringEmpty {
		keyCache := fmt.Sprintf("%s:%s", constant.CacheCategoryInfo, parentCate)
		existed := service.Cache.Exists(ctx, keyCache)
		if existed == constant.ValueEmpty {
			r, err := service.DocsRepository.DBGetInfoCategoryBySlug(ctx, parentCate, childCate)
			if err != nil {
				logger.Error("DocsService:GetInfoCategoryBySlug: - ERROR: %v", err)
				return nil, err
			}

			go service.Cache.Set(context.Background(), keyCache, r, constant.CacheExpiresInOneMonth)
			return r, nil
		} else {
			result := new(dto.CategoryInfo)
			r := service.Cache.Get(ctx, keyCache)
			err := json.Unmarshal([]byte(r), &result)
			if err != nil {
				logger.Error("DocsService:GetInfoCategoryBySlug: - Unmarshall with ERROR: %v", err)
			}
			return result, nil
		}
	}

	// Cache Child category
	if childCate != constant.StringEmpty && parentCate != constant.StringEmpty {
		keyCache := fmt.Sprintf("%s:%s", constant.CacheCategoryInfo, childCate)
		existed := service.Cache.Exists(ctx, keyCache)
		if existed == constant.ValueEmpty {
			r, err := service.DocsRepository.DBGetInfoCategoryBySlug(ctx, parentCate, childCate)
			if err != nil {
				logger.Error("DocsService:GetInfoCategoryBySlug: - ERROR: %v", err)
				return nil, err
			}

			go service.Cache.Set(context.Background(), keyCache, r, constant.CacheExpiresInOneMonth)
			return r, nil
		} else {
			result := new(dto.CategoryInfo)
			r := service.Cache.Get(ctx, keyCache)
			err := json.Unmarshal([]byte(r), &result)
			if err != nil {
				logger.Error("DocsService:GetInfoCategoryBySlug: - Unmarshall with ERROR: %v", err)
			}
			return result, nil
		}
	}

	return nil, nil
}

func (service *DocsService) CountPostsByCategory(ctx context.Context, parentCate, childCate string) *int {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	r := service.DocsRepository.DBCountPostsByCategory(ctx, parentCate, childCate)

	return r
}

func (service *DocsService) GetCategoryList(ctx context.Context, parent string) (*[]dto.CategoryList, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	r, err := service.DocsRepository.DBGetCategoryList(ctx, parent)
	if err != nil {
		logger.Error("DocsService:GetCategoryList: - ERROR: %v", err)
		return nil, err
	}

	// cache parent category
	if parent == constant.StringEmpty {
		keyCache := fmt.Sprintf("%s:%s", constant.CacheCategoryParentList, parent)
		existed := service.Cache.Exists(ctx, keyCache)
		if existed == constant.ValueEmpty {
			r, err := service.DocsRepository.DBGetCategoryList(ctx, parent)
			if err != nil {
				logger.Error("DocsService:GetCategoryList: - ERROR: %v", err)
				return nil, err
			}

			go service.Cache.Set(context.Background(), keyCache, r, constant.CacheExpiresInOneMonth)
			return r, nil
		} else {
			result := new([]dto.CategoryList)
			r := service.Cache.Get(ctx, keyCache)
			err := json.Unmarshal([]byte(r), &result)
			if err != nil {
				logger.Error("DocsService:GetCategoryList: - Unmarshall with ERROR: %v", err)
			}
			return result, nil
		}
	}

	// cache child category
	if parent != constant.StringEmpty {
		keyCache := fmt.Sprintf("%s:%s", constant.CacheCategoryChildList, parent)
		existed := service.Cache.Exists(ctx, keyCache)
		if existed == constant.ValueEmpty {
			r, err := service.DocsRepository.DBGetCategoryList(ctx, parent)
			if err != nil {
				logger.Error("DocsService:GetCategoryList: - ERROR: %v", err)
				return nil, err
			}

			go service.Cache.Set(context.Background(), keyCache, r, constant.CacheExpiresInOneMonth)
			return r, nil
		} else {
			result := new([]dto.CategoryList)
			r := service.Cache.Get(ctx, keyCache)
			err := json.Unmarshal([]byte(r), &result)
			if err != nil {
				logger.Error("DocsService:GetCategoryList: - Unmarshall with ERROR: %v", err)
			}
			return result, nil
		}
	}

	return r, nil

}

func (service *DocsService) GetPostListByCategory(ctx context.Context, parentCate, childCate string, pageSize, pageIndex int) *dto.PostListByCategoryRes {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	if pageSize <= constant.ValueEmpty || pageIndex <= constant.ValueEmpty {
		pageSize = constant.PageSizeDefault
		pageIndex = constant.PageIndexDefault
	}

	result := dto.PostListByCategoryRes{}

	// Handle node route example.com/category
	// Cache Parent category
	if parentCate != constant.StringEmpty && childCate == constant.StringEmpty {
		keyCache := fmt.Sprintf("%s:%s:%d:%d", constant.CachePostListByCategory, parentCate, pageSize, pageIndex)
		existed := service.Cache.Exists(ctx, keyCache)
		if existed == constant.ValueEmpty {

			channel := make(chan *[]dto.PostListByCategory)
			channel1 := make(chan *int)
			defer close(channel)
			defer close(channel1)

			// get post db
			go func() {
				channel <- service.DocsRepository.DBGetPostListByCategory(ctx, parentCate, childCate, pageSize, pageIndex)
			}()

			// count post db
			go func() {
				channel1 <- service.DocsRepository.DBCountPostsByCategory(ctx, parentCate, childCate)
			}()

			for i := 0; i < 2; i++ {
				// Await both of these values
				select {
				case msg1, ok1 := <-channel:
					if ok1 {
						result.ListPost = msg1
					}
				case msg2, ok2 := <-channel1:
					if ok2 {
						result.TotalPage = msg2
					}

				}
			}

			// Wait for results from 2 go routines then set cache
			go service.Cache.Set(context.Background(), keyCache, result, constant.CacheExpiresInOneMonth)

			return &result
		} else {
			r := service.Cache.Get(ctx, keyCache)
			err := json.Unmarshal([]byte(r), &result)
			if err != nil {
				logger.Error("DocsService:GetPostListByCategory: - Unmarshall with ERROR: %v", err)
			}
			return &result

		}
	}

	// Handle leaf route example.com/category/subCategory
	// Cache Child category
	if childCate != constant.StringEmpty && parentCate != constant.StringEmpty {

		keyCache := fmt.Sprintf("%s:%s:%d:%d", constant.CachePostListByCategory, childCate, pageSize, pageIndex)
		existed := service.Cache.Exists(ctx, keyCache)
		if existed == constant.ValueEmpty {

			channel := make(chan *[]dto.PostListByCategory)
			channel1 := make(chan *int)
			defer close(channel)
			defer close(channel1)

			// get post db
			go func() {
				channel <- service.DocsRepository.DBGetPostListByCategory(ctx, parentCate, childCate, pageSize, pageIndex)
			}()

			// count post db
			go func() {
				channel1 <- service.DocsRepository.DBCountPostsByCategory(ctx, parentCate, childCate)
			}()

			for i := 0; i < 2; i++ {
				// Await both of these values
				select {
				case msg1, ok1 := <-channel:
					if ok1 {
						result.ListPost = msg1
					}
				case msg2, ok2 := <-channel1:
					if ok2 {
						result.TotalPage = msg2
					}

				}
			}

			// Wait for results from 2 go routines then set cache
			go service.Cache.Set(context.Background(), keyCache, result, constant.CacheExpiresInOneMonth)

			return &result
		} else {
			r := service.Cache.Get(ctx, keyCache)
			err := json.Unmarshal([]byte(r), &result)
			if err != nil {
				logger.Error("DocsService:GetPostListByCategory: - Unmarshall with ERROR: %v", err)
			}

			return &result

		}
	}

	return nil

}

func (service *DocsService) DocsGetSourceByPost(ctx context.Context, id string, pageSize, pageIndex int) (*[]dto.SourceByPost, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	if pageSize <= constant.ValueEmpty || pageIndex <= constant.ValueEmpty {
		pageSize = constant.PageSizeDefault
		pageIndex = constant.PageIndexDefault
	}

	r, err := service.DocsRepository.DBGetSourceByPost(ctx, id, pageSize, pageIndex)
	if err != nil {
		logger.Error("DocsService:DocsGetSourceByPost: - ERROR: %v", err)
		return nil, err
	}

	return r, nil
}

func (service *DocsService) DocsGetPostListByTag(ctx context.Context, tag string, pageSize, pageIndex int) (*[]dto.PostListByTag, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	if pageSize <= constant.ValueEmpty || pageIndex <= constant.ValueEmpty {
		pageSize = constant.PageSizeDefault
		pageIndex = constant.PageIndexDefault
	}

	// Cache
	var result *[]dto.PostListByTag
	keyCache := fmt.Sprintf("%s:%s:%d:%d", constant.CachePostListByTag, tag, pageSize, pageIndex)
	existed := service.Cache.Exists(ctx, keyCache)
	if existed == constant.ValueEmpty {
		r, err := service.DocsRepository.DBGetPostListByTag(ctx, tag, pageSize, pageIndex)
		if err != nil {
			logger.Error("DocsService:DocsGetPostListByTag:DBGetPostListByTag - ERROR: %v", err)
			return nil, err
		}

		go service.Cache.Set(context.Background(), keyCache, r, constant.CacheExpiresInOneMonth)
		return r, nil

	} else {
		r := service.Cache.Get(ctx, keyCache)
		err := json.Unmarshal([]byte(r), &result)
		if err != nil {
			logger.Error("DocsService:DocsGetPostListByTag: - Unmarshall with ERROR: %v", err)
		}
		return result, err
	}

}

func (service *DocsService) DocsGetTagList(ctx context.Context, pageSize, pageIndex int) (*[]dto.Tag, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	if pageSize <= constant.ValueEmpty || pageIndex <= constant.ValueEmpty {
		pageSize = constant.PageSizeDefault
		pageIndex = constant.PageIndexDefault
	}

	// Cache
	var result *[]dto.Tag
	existed := service.Cache.Exists(ctx, constant.CacheTagList)
	if existed == constant.ValueEmpty {
		r, err := service.DocsRepository.DBGetTagList(ctx, pageSize, pageIndex)
		if err != nil {
			logger.Error("DocsService:DocsGetTagList: - ERROR: %v", err)
			return nil, err
		}

		go service.Cache.Set(context.Background(), constant.CacheTagList, r, constant.CacheExpiresInOneMonth)
		return r, nil

	} else {
		r := service.Cache.Get(ctx, constant.CacheTagList)
		err := json.Unmarshal([]byte(r), &result)
		if err != nil {
			logger.Error("DocsService:DocsGetTagList: - Unmarshall with ERROR: %v", err)

		}
		return result, err
	}
}

func (service *DocsService) DocsGetPostDetail(ctx context.Context, id string) (*map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	result := map[string]interface{}{}

	// Cache
	keyCache := fmt.Sprintf("%s:%s", constant.CachePostDetail, id)
	existed := service.Cache.Exists(ctx, keyCache)
	if existed == constant.ValueEmpty {

		// Get post detail
		resulPostDetail, err := service.DocsRepository.DBGetPostDetail(ctx, id)
		if err != nil {
			logger.Error("DocsService:DocsGetPostDetail: - ERROR: %v", err)
			return nil, err
		}

		// Get source by post
		resultSourceByPost, err := service.DocsGetSourceByPost(ctx, id, 10, 1)
		if err != nil {
			logger.Error("DocsService:DocsGetSourceByPost: - ERROR: %v", err)
			return nil, err
		}

		// Get tag by post
		resultTagListByPost, err := service.DocsRepository.DBGetTagListByPost(ctx, id)
		if err != nil {
			logger.Error("DocsService:DocsGetSourceByPost: - ERROR: %v", err)
			return nil, err
		}

		// Set result
		result["post_detail"] = resulPostDetail
		result["source_detail"] = resultSourceByPost
		result["tag_list_by_post"] = resultTagListByPost

		go service.Cache.Set(context.Background(), keyCache, result, constant.CacheExpiresInOneMonth)
		return &result, nil

	} else {
		r := service.Cache.Get(ctx, keyCache)
		err := json.Unmarshal([]byte(r), &result)
		if err != nil {
			logger.Error("DocsService:DocsGetPostDetail: - Unmarshall with ERROR: %v", err)
			return nil, err
		}
		return &result, err
	}
}

//------------------------------------
// Docs:Post
//------------------------------------

func (service *DocsService) DocsPostCreate(ctx context.Context, p *dto.Post) error {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	p.Id = util.NewNanoIdString()
	p.Title = util.ToTrimSpace(p.Title)
	p.Description = util.ToTrimSpace(p.Description)
	p.Slug = util.ToTrimSpace(p.Slug)
	p.Content = util.ToTrimSpace(p.Content)

	// If you do not enter the slug, then use the title as the slug
	if p.Slug == constant.StringEmpty {
		p.Slug = util.ToSlugFromTitleWithId(p.Title, p.Id)
	} else {
		p.Slug = util.ToSlugWithId(p.Slug, p.Id)
	}

	// one channel to share communicate two goroutine, 1 for insert, 2 for index caching
	signalInsertOk := make(chan bool)

	// Add counter goroutine
	var wg sync.WaitGroup
	wg.Add(5)

	// Create data into the database
	go func(wg *sync.WaitGroup, signalChan *chan bool) {
		defer wg.Done()
		defer close(signalInsertOk)
		err := service.DocsRepository.DBDocsPostCreate(ctx, p)
		if err != nil {
			// if failed on insert, then say to channel is failed
			*signalChan <- false
			logger.Error("DocsService:DocsPostCreate: - ERROR: %v", err)
			return
		}
		*signalChan <- true

	}(&wg, &signalInsertOk)

	// Index post for search service
	go func(wg *sync.WaitGroup, signalChan *chan bool) {
		defer wg.Done()
		// go routine will wait for insert failed or success
		signal := <-signalInsertOk
		if !signal {
			return
		}
		err := service.Search.CreateSearchDocument(p.Id, p.Title, p.Content, p.Slug)
		if err != nil {
			logger.Error("DocsService:DocsPostCreate:CreateSearchDocument: - ERROR: %v", err)
			return
		}
	}(&wg, &signalInsertOk)

	// Clear cache post list by category
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		keyCache := fmt.Sprintf("%s:%s:", constant.CachePostListByCategory, p.CategorySlug)
		err := service.Cache.DelKeys(context.Background(), keyCache)
		if err != nil {
			logger.Error("DocsService:DocsPostCreate:ClearCachePostListByCategory - ERROR: %v", err)
			return
		}
	}(&wg)

	// Clear cache post list latest
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		keyCache := fmt.Sprintf("%s", constant.CachePostListLatest)
		err := service.Cache.DelKeys(context.Background(), keyCache)
		if err != nil {
			logger.Error("DocsService:DocsPostCreate:ClearCachePostListLatest - ERROR: %v", err)
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		keyCache := fmt.Sprintf("%s", constant.CachePostInSameCategory)
		err := service.Cache.DelKeys(context.Background(), keyCache)
		if err != nil {
			logger.Error("DocsService:DocsPostCreate:ClearCachePostListLatest - ERROR: %v", err)
			return
		}
	}(&wg)

	wg.Wait()
	return nil
}

func (service *DocsService) DocsPostDelete(ctx context.Context, id, categorySlug string) error {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	// Add counter goroutine
	var wg sync.WaitGroup
	wg.Add(4)

	// Delete DB Post
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err := service.DocsRepository.DBDocsPostDelete(ctx, id)
		if err != nil {
			logger.Error("DocsService:DocsPostDelete: - ERROR: %v", err)
			return
		}
	}(&wg)

	// Remove CachePostDetail
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		keyCache := fmt.Sprintf("%s:%s", constant.CachePostDetail, id)
		service.Cache.Del(context.Background(), keyCache)
	}(&wg)

	// Remove index search
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		keyIndex := fmt.Sprintf("%s:%s", constant.CacheSearchPost, id)
		err := service.Search.DelDoc(keyIndex)
		if err != nil {
			logger.Error("DocsService:DocsPostDelete:RemoveIndex - ERROR: %v", err)
		}
	}(&wg)

	// Clear cache list post by category
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		keyCache := fmt.Sprintf("%s:%s:", constant.CachePostListByCategory, categorySlug)
		service.Cache.DelKeys(context.Background(), keyCache)
	}(&wg)

	wg.Wait()
	return nil
}

func (service *DocsService) DocsPostGet(ctx context.Context, id string) (*dto.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	r, err := service.DocsRepository.DBDocsPostGet(ctx, id)
	if err != nil {
		logger.Error("DocsService:DocsPostGet: - ERROR: %v", err)
		return nil, err
	}

	return r, nil
}

func (service *DocsService) DocsPostList(ctx context.Context, status, pageSize, pageIndex int) (*[]dto.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	if pageSize <= constant.ValueEmpty || pageIndex <= constant.ValueEmpty {
		pageSize = constant.PageSizeDefault
		pageIndex = constant.PageIndexDefault
	}

	r, err := service.DocsRepository.DBDocsPostList(ctx, status, pageSize, pageIndex)
	if err != nil {
		logger.Error("DocsService:DocsPostList: - ERROR: %v", err)
		return nil, err
	}

	return r, nil
}

func (service *DocsService) DocsPostUpdate(ctx context.Context, condition *dto.Post) error {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	err := service.DocsRepository.DBDocsPostUpdate(ctx, condition)
	if err != nil {
		logger.Error("DocsService:DocsPostUpdate: - ERROR: %v", err)
		return err
	}

	// Remove cache
	keyCache := fmt.Sprintf("%s:%s", constant.CachePostDetail, condition.Id)
	go service.Cache.Del(context.Background(), keyCache)

	// Remove index
	// Delete the index, if the index is deleted successfully, update the index with new information
	keyIndex := fmt.Sprintf("%s:%s", constant.CacheSearchPost, condition.Id)
	go func() {
		err := service.Search.DelDoc(keyIndex)
		if err != nil {
			logger.Error("DocsService:DocsPostDelete:RemoveIndex - ERROR: %v", err)
		}
		if err == nil {
			if condition.Status == constant.PostStatusPublished {
				go func() {
					err = service.Search.CreateSearchDocument(condition.Id, condition.Title, condition.Content, condition.Slug)
					if err != nil {
						logger.Error("DocsService:CreateSearchDocument: - ERROR: %v", err)
						return
					}
				}()
			}
		}
	}()

	// Clear cache list post by category
	keyCacheListPostByCategory := fmt.Sprintf("%s:%s:", constant.CachePostListByCategory, condition.CategorySlug)
	go service.Cache.DelKeys(context.Background(), keyCacheListPostByCategory)

	return nil
}

//------------------------------------
// Docs:Tag
//------------------------------------

func (service *DocsService) DocsTagCreate(ctx context.Context, p *dto.Tag) error {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	p.Title = util.ToTrimSpace(p.Title)

	// If you do not enter the slug, then use the title as the slug
	if p.Slug == constant.StringEmpty {
		p.Slug = util.ToSlugFromTitleWithoutId(p.Title)
	}

	err := service.DocsRepository.DBDocsTagCreate(ctx, p)
	if err != nil {
		logger.Error("DocsService:DocsTagCreate: - ERROR: %v", err)
		return err
	}

	// Remove cache
	go func() {
		errCache := service.Cache.DelKeys(context.Background(), constant.CachePostListByTag)
		if errCache != nil {
			logger.Error("DocsService:DocsTagCreate: - Cache with ERROR: %v", err)
			return
		}
	}()

	return nil
}

func (service *DocsService) DocsTagDelete(ctx context.Context, id int, slug string) error {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	err := service.DocsRepository.DBDocsTagDelete(ctx, id)
	if err != nil {
		logger.Error("DocsService:DocsTagDelete: - ERROR: %v", err)
		return err
	}

	// Remove cache
	go service.Cache.Del(context.Background(), constant.CacheTagList)

	// Remove cache tag info
	keyCacheTagInfo := fmt.Sprintf("%s:%s", constant.CacheTagInfo, slug)
	go service.Cache.Del(context.Background(), keyCacheTagInfo)

	return nil
}

func (service *DocsService) DocsTagGet(ctx context.Context, id int) (*dto.Tag, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	r, err := service.DocsRepository.DBDocsTagGet(ctx, id)
	if err != nil {
		logger.Error("DocsService:DocsTagGet: - ERROR: %v", err)
		return nil, err
	}

	return r, nil
}

func (service *DocsService) DocsTagList(ctx context.Context, pageSize, pageIndex int) (*[]dto.Tag, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	if pageSize <= constant.ValueEmpty || pageIndex <= constant.ValueEmpty {
		pageSize = constant.PageSizeDefault
		pageIndex = constant.PageIndexDefault
	}

	r, err := service.DocsRepository.DBDocsTagList(ctx, pageSize, pageIndex)
	if err != nil {
		logger.Error("DocsService:DocsTagList: - ERROR: %v", err)
		return nil, err
	}

	return r, nil
}

func (service *DocsService) DocsTagUpdate(ctx context.Context, condition *dto.Tag) error {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	err := service.DocsRepository.DBDocsTagUpdate(ctx, condition)
	if err != nil {
		logger.Error("DocsService:DocsTagUpdate: - ERROR: %v", err)
		return err
	}

	// Remove cache
	keyCache := fmt.Sprintf("%s:%s", constant.CacheTagList, condition.Slug)
	go service.Cache.Del(context.Background(), keyCache)

	// Remove cache tag info
	keyCacheTagInfo := fmt.Sprintf("%s:%s", constant.CacheTagInfo, condition.Slug)
	go service.Cache.Del(context.Background(), keyCacheTagInfo)

	return nil
}

//------------------------------------
// Docs:Category
//------------------------------------

func (service *DocsService) DocsCategoryCreate(ctx context.Context, p *dto.Category) error {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	if p.Slug == constant.StringEmpty {
		p.Slug = util.ToSlug(p.Title)
	}

	err := service.DocsRepository.DBDocsCategoryCreate(ctx, p)
	if err != nil {
		logger.Error("DocsService:DocsCategoryCreate: - ERROR: %v", err)
		return err
	}
	return nil
}

func (service *DocsService) DocsCategoryDelete(ctx context.Context, id int, slug string) error {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	err := service.DocsRepository.DBDocsCategoryDelete(ctx, id)
	if err != nil {
		logger.Error("DocsService:DocsCategoryDelete: - ERROR: %v", err)
		return err
	}

	// Remove cache category info
	keyCache := fmt.Sprintf("%s:%s", constant.CacheCategoryInfo, slug)
	go service.Cache.Del(context.Background(), keyCache)
	return nil
}

func (service *DocsService) DocsCategoryGet(ctx context.Context, id int) (*dto.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	r, err := service.DocsRepository.DBDocsCategoryGet(ctx, id)
	if err != nil {
		logger.Error("DocsService:DocsCategoryGet: - ERROR: %v", err)
		return nil, err
	}

	return r, nil
}

func (service *DocsService) DocsCategoryList(ctx context.Context, pageSize, pageIndex int) (*[]dto.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	if pageSize <= constant.ValueEmpty || pageIndex <= constant.ValueEmpty {
		pageSize = constant.PageSizeDefault
		pageIndex = constant.PageIndexDefault
	}

	r, err := service.DocsRepository.DBDocsCategoryList(ctx, pageSize, pageIndex)
	if err != nil {
		logger.Error("DocsService:DocsCategoryList: - ERROR: %v", err)
		return nil, err
	}

	return r, nil
}

func (service *DocsService) DocsCategoryUpdate(ctx context.Context, condition *dto.Category) error {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	err := service.DocsRepository.DBDocsCategoryUpdate(ctx, condition)
	if err != nil {
		logger.Error("DocsService:DocsCategoryUpdate: - ERROR: %v", err)
		return err
	}

	// Remove cache
	keyCache := fmt.Sprintf("%s:%s", constant.CachePostListByCategory, condition.Slug)
	go service.Cache.Del(context.Background(), keyCache)

	// TODO: remove category list cache

	// Remove cache category info
	keyCacheCateInfo := fmt.Sprintf("%s:%s", constant.CacheCategoryInfo, condition.Slug)
	go service.Cache.Del(context.Background(), keyCacheCateInfo)

	return nil

}

//------------------------------------
// Docs:PostTag
//------------------------------------

func (service *DocsService) DocsPostTagCreate(ctx context.Context, p []*dto.PostTag) error {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	err := service.DocsRepository.DBDocsPostTagCreate(ctx, p)
	if err != nil {
		logger.Error("DocsService:DocsPostTagCreate: - ERROR: %v", err)
		return err
	}
	return nil
}

func (service *DocsService) DocsPostTagDelete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	err := service.DocsRepository.DBDocsPostTagDelete(ctx, id)
	if err != nil {
		logger.Error("DocsService:DocsPostTagDelete: - ERROR: %v", err)
		return err
	}
	return nil
}

func (service *DocsService) DocsPostTagGet(ctx context.Context, id int) (*dto.PostTag, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	r, err := service.DocsRepository.DBDocsPostTagGet(ctx, id)
	if err != nil {
		logger.Error("DocsService:DocsPostTagGet: - ERROR: %v", err)
		return nil, err
	}

	return r, nil
}

func (service *DocsService) DocsPostTagList(ctx context.Context, pageSize, pageIndex int) (*[]dto.PostTag, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	if pageSize <= constant.ValueEmpty || pageIndex <= constant.ValueEmpty {
		pageSize = constant.PageSizeDefault
		pageIndex = constant.PageIndexDefault
	}

	r, err := service.DocsRepository.DBDocsPostTagList(ctx, pageSize, pageIndex)
	if err != nil {
		logger.Error("DocsService:DocsPostTagList: - ERROR: %v", err)
		return nil, err
	}

	return r, nil
}

func (service *DocsService) DocsPostTagUpdate(ctx context.Context, condition *dto.PostTagUpdate) error {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	err := service.DocsRepository.DBDocsPostTagUpdate(ctx, condition)
	if err != nil {
		logger.Error("DocsService:DocsPostTagUpdate: - ERROR: %v", err)
		return err
	}

	return nil
}

//------------------------------------
// Docs:Source
//------------------------------------

func (service *DocsService) DocsSourceCreate(ctx context.Context, p *dto.Source) error {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	err := service.DocsRepository.DBDocsSourceCreate(ctx, p)
	if err != nil {
		logger.Error("DocsService:DocsSourceCreate: - ERROR: %v", err)
		return err
	}

	// Remove cache
	keyCache := fmt.Sprintf("%s:%s", constant.CachePostDetail, p.PostId)
	go service.Cache.Del(context.Background(), keyCache)

	return nil
}

func (service *DocsService) DocsSourceDelete(ctx context.Context, id int, postId string) error {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	err := service.DocsRepository.DBDocsSourceDelete(ctx, id)
	if err != nil {
		logger.Error("DocsService:DocsSourceDelete: - ERROR: %v", err)
		return err
	}

	// Remove cache
	keyCache := fmt.Sprintf("%s:%s", constant.CacheTagList, postId)
	go service.Cache.Del(context.Background(), keyCache)

	return nil
}

func (service *DocsService) DocsSourceGet(ctx context.Context, id int) (*dto.Source, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	r, err := service.DocsRepository.DBDocsSourceGet(ctx, id)
	if err != nil {
		logger.Error("DocsService:DocsSourceGet: - ERROR: %v", err)
		return nil, err
	}

	return r, nil
}

func (service *DocsService) DocsSourceList(ctx context.Context, pageSize, pageIndex int) (*[]dto.Source, error) {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	if pageSize <= constant.ValueEmpty || pageIndex <= constant.ValueEmpty {
		pageSize = constant.PageSizeDefault
		pageIndex = constant.PageIndexDefault
	}

	r, err := service.DocsRepository.DBDocsSourceList(ctx, pageSize, pageIndex)
	if err != nil {
		logger.Error("DocsService:DocsSourceList: - ERROR: %v", err)
		return nil, err
	}

	return r, nil
}

func (service *DocsService) DocsSourceUpdate(ctx context.Context, condition *dto.Source) error {
	ctx, cancel := context.WithTimeout(ctx, constant.TimeoutRequestDefault)
	defer cancel()

	err := service.DocsRepository.DBDocsSourceUpdate(ctx, condition)
	if err != nil {
		logger.Error("DocsService:DocsSourceUpdate: - ERROR: %v", err)
		return err
	}

	// Remove cache
	keyCache := fmt.Sprintf("%s:%s", constant.CachePostDetail, condition.PostId)
	go service.Cache.Del(context.Background(), keyCache)

	return nil
}
