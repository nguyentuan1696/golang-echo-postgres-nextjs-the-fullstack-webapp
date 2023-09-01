package service

import (
	"context"
	"thichlab-backend-docs/dto"
	"thichlab-backend-docs/infrastructure/cache"
	"thichlab-backend-docs/infrastructure/search"
	"thichlab-backend-docs/module/docs/repository"
)

type DocsService struct {
	DocsRepository repository.IDocsRepository
	Cache          cache.Client
	Search         search.Client
}

func NewDocsService(repository repository.IDocsRepository, cache cache.Client, search search.Client) IDocsService {
	docsService := DocsService{
		Cache:  cache,
		Search: search,
	}
	docsService.DocsRepository = repository
	return &docsService
}

type IDocsService interface {
	DocsPostCreate(ctx context.Context, p *dto.Post) error
	DocsPostDelete(ctx context.Context, id, categorySlug string) error
	DocsPostGet(ctx context.Context, id string) (*dto.Post, error)
	DocsPostList(ctx context.Context, status, pageSize, pageIndex int) (*[]dto.Post, error)
	DocsPostUpdate(ctx context.Context, condition *dto.Post) error

	DocsTagCreate(ctx context.Context, p *dto.Tag) error
	DocsTagDelete(ctx context.Context, id int, slug string) error
	DocsTagGet(ctx context.Context, id int) (*dto.Tag, error)
	DocsTagList(ctx context.Context, pageSize, pageIndex int) (*[]dto.Tag, error)
	DocsTagUpdate(ctx context.Context, condition *dto.Tag) error

	DocsCategoryCreate(ctx context.Context, p *dto.Category) error
	DocsCategoryDelete(ctx context.Context, id int, slug string) error
	DocsCategoryGet(ctx context.Context, id int) (*dto.Category, error)
	DocsCategoryList(ctx context.Context, pageSize, pageIndex int) (*[]dto.Category, error)
	DocsCategoryUpdate(ctx context.Context, condition *dto.Category) error

	DocsPostTagCreate(ctx context.Context, p []*dto.PostTag) error
	DocsPostTagDelete(ctx context.Context, id int) error
	DocsPostTagGet(ctx context.Context, id int) (*dto.PostTag, error)
	DocsPostTagList(ctx context.Context, pageSize, pageIndex int) (*[]dto.PostTag, error)
	DocsPostTagUpdate(ctx context.Context, condition *dto.PostTagUpdate) error

	DocsSourceCreate(ctx context.Context, p *dto.Source) error
	DocsSourceDelete(ctx context.Context, id int, postId string) error
	DocsSourceGet(ctx context.Context, id int) (*dto.Source, error)
	DocsSourceList(ctx context.Context, pageSize, pageIndex int) (*[]dto.Source, error)
	DocsSourceUpdate(ctx context.Context, condition *dto.Source) error

	DocsGetTagList(ctx context.Context, pageSize, pageIndex int) (*[]dto.Tag, error)
	DocsGetPostListByTag(ctx context.Context, tag string, pageSize, pageIndex int) (*[]dto.PostListByTag, error)
	DocsGetSourceByPost(ctx context.Context, id string, pageSize, pageIndex int) (*[]dto.SourceByPost, error)
	DocsGetPostDetail(ctx context.Context, id string) (*map[string]interface{}, error)
	GetPostListByCategory(ctx context.Context, parentCate, childCate string, pageSize, pageIndex int) *dto.PostListByCategoryRes
	DocsSearch(ctx context.Context, query string, pageSize, pageIndex int) (*dto.PostsByDocSearchRes, error)
	GetCategoryList(ctx context.Context, parent string) (*[]dto.CategoryList, error)
	CountPostsByCategory(ctx context.Context, parentCate, childCate string) *int
	GetInfoCategoryBySlug(ctx context.Context, parentCate, childCate string) (*dto.CategoryInfo, error)
	GetTagInfoBySlug(ctx context.Context, slug string) (*dto.TagInfo, error)
	GetPostsInSameCategory(ctx context.Context, id string, pageSize, pageIndex int) (*[]dto.PostsInSameCategory, error)
	GetPostListLatestCreated(ctx context.Context, pageSize, pageIndex int) (*[]dto.PostListLatestCreated, error)
	GetChildCategoryList(ctx context.Context, parentCate string, pageSize, pageIndex int) (*[]dto.GetChildCategoryList, error)
	GetMostSearchedKeywords(ctx context.Context, pageSize, pageIndex int64) []string
	IndexSearchKeywords(ctx context.Context, query string) error
	IsValidQuerySearch(query string) bool
	GetCategoryListByHierarchy(ctx context.Context) (*[]dto.CategoryListByHierarchyRes, error)
}
