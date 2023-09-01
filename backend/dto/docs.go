package dto

type CategoryListByHierarchy struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Slug        string `json:"slug" db:"slug"`
	Parent      string `json:"parent" db:"parent"`
}

type ParentCategoryList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Slug        string `json:"slug" db:"slug"`
}

type CategoryListByHierarchyRes struct {
	Parent string                    `json:"parent"`
	Child  []CategoryListByHierarchy `json:"child"`
}

type GetChildCategoryList struct {
	Id    string `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Slug  string `json:"slug" db:"slug"`
}

type PostsByDocSearch struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type PostListLatestCreated struct {
	Id    string `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Slug  string `json:"slug" db:"slug"`
}

type PostsInSameCategory struct {
	Id    string `json:"id" db:"id"`
	Title string `json:"title" dbL:"title"`
	Slug  string `json:"slug" db:"slug"`
}

type TagInfo struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type CategoryInfo struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type CategoryList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Slug        string `json:"slug" db:"slug"`
	Parent      string `json:"parent" db:"parent"`
}

type PostsByDocSearchRes struct {
	Total int                `json:"total"`
	Data  []PostsByDocSearch `json:"data"`
}

type PostListByTag struct {
	Id    string `json:"id,omitempty" db:"id"`
	Title string `json:"title,omitempty" db:"title"`
	Slug  string `json:"slug,omitempty" db:"slug"`
}

type TagListByPost struct {
	TagId    int    `json:"tag_id" db:"tag_id"`
	SlugTag  string `json:"slug_tag" db:"slug_tag"`
	TitleTag string `json:"title_tag" db:"title_tag"`
	PostId   string `json:"post_id" db:"post_id"`
}

type PostListByCategoryRes struct {
	ListPost  *[]PostListByCategory `json:"list_post"`
	TotalPage *int                  `json:"total_page"`
}

type PostListByCategory struct {
	IdPost       string `json:"id_post" db:"id_post"`
	TitlePost    string `json:"title_post" db:"title_post"`
	SlugPost     string `json:"slug_post" db:"slug_post"`
	CategoryPost string `json:"category_post" db:"category_post"`
}

type SourceByPost struct {
	SourceId    int    `json:"source_id,omitempty" db:"source_id"`
	SourceTitle string `json:"source_title,omitempty" db:"source_title"`
	SourceUrl   string `json:"source_url" db:"source_url"`
}

type Post struct {
	Id           string `json:"id,omitempty" db:"id"`
	Title        string `json:"title,omitempty" db:"title"`
	Description  string `json:"description,omitempty" db:"description"`
	Content      string `json:"content,omitempty" db:"content"`
	Slug         string `json:"slug,omitempty" db:"slug"`
	Status       int8   `json:"status,omitempty" db:"status"` // 1:PostStatusPublished 2:PostStatusHidden
	CategoryId   int    `json:"category_id" db:"category_id"`
	CategoryName string `json:"category_name" db:"category_name"`
	CategorySlug string `json:"category_slug" db:"category_slug"`
	CreatedAt    int64  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt    int64  `json:"updated_at,omitempty" db:"updated_at"`
}

type Tag struct {
	Id          int    `json:"id,omitempty" db:"id"`
	Title       string `json:"title,omitempty" db:"title"`
	Slug        string `json:"slug,omitempty" db:"slug"`
	Description string `json:"description,omitempty" db:"description"`
	CreatedAt   int64  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   int64  `json:"updated_at,omitempty" db:"updated_at"`
}

type Category struct {
	Id          int    `json:"id,omitempty" db:"id"`
	Title       string `json:"title,omitempty" db:"title"`
	Slug        string `json:"slug,omitempty" db:"slug"`
	Parent      string `json:"parent" db:"parent"`
	Description string `json:"description,omitempty" db:"description"`
	CreatedAt   int64  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   int64  `json:"updated_at,omitempty" db:"updated_at"`
}

type Source struct {
	Id        int    `json:"id,omitempty" db:"id"`
	Title     string `json:"title,omitempty" db:"title"`
	Url       string `json:"url,omitempty" db:"url"`
	PostId    string `json:"post_id,omitempty" db:"post_id"`
	CreatedAt int64  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt int64  `json:"updated_at,omitempty" db:"updated_at"`
}

type PostTag struct {
	Id        int    `json:"id" db:"id"`
	PostId    string `json:"post_id" db:"post_id"`
	TagId     []int  `json:"tag_id" db:"tag_id"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
	UpdatedAt int64  `json:"updated_at" db:"updated_at"`
}

type PostTagUpdate struct {
	Id        int    `json:"id" db:"id"`
	PostId    string `json:"post_id" db:"post_id"`
	TagId     int    `json:"tag_id" db:"tag_id"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
	UpdatedAt int64  `json:"updated_at" db:"updated_at"`
}
