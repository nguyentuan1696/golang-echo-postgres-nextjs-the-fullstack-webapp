package constant

import "time"

const (
	CachePrefix string = "backend:thichlab:docs:"
)

const (
	CacheExpiresInOneHour  = time.Hour
	CacheExpiresInOneDay   = CacheExpiresInOneHour * 24
	CacheExpiresInOneMonth = CacheExpiresInOneDay * 30
)

const (
	CacheTagList            = CachePrefix + "tag-list"
	CachePostListByTag      = CachePrefix + "post-list-tag"
	CachePostDetail         = CachePrefix + "post-detail"
	CachePostListByCategory = CachePrefix + "post-list-category"
	CacheCategoryParentList = CachePrefix + "category-parent-list"
	CacheCategoryChildList  = CachePrefix + "category-child-list"
	CacheCategoryInfo       = CachePrefix + "category-info"
	CacheTagInfo            = CachePrefix + "tag-info"
	CachePostListLatest     = CachePrefix + "post-list-latest"
	CachePostInSameCategory = CachePrefix + "post-same-category"
)

const (
	CacheSearchPost = CachePrefix + "search-post"
)

const (
	CacheSearchedKeywordsRankings = CachePrefix + "searched-keywords-ranking"
)
