package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"thichlab-backend-docs/dto"
	"thichlab-backend-docs/infrastructure/repository"
)

type DocsRepository struct {
	Postgres repository.PostgresRepository
}

func NewDocsRepository(dbContext *sql.DB, sqlxDBContext *sqlx.DB) IDocsRepository {
	docsRepository := DocsRepository{}
	docsRepository.Postgres.SetDbContext(dbContext, sqlxDBContext)
	return &docsRepository
}

type IDocsRepository interface {
	DBDocsPostCreate(ctx context.Context, p *dto.Post) error
	DBDocsPostDelete(ctx context.Context, id string) error
	DBDocsPostGet(ctx context.Context, id string) (*dto.Post, error)
	DBDocsPostList(ctx context.Context, status, pageSize, pageIndex int) (*[]dto.Post, error)
	DBDocsPostUpdate(ctx context.Context, condition *dto.Post) error

	DBDocsTagCreate(ctx context.Context, p *dto.Tag) error
	DBDocsTagDelete(ctx context.Context, id int) error
	DBDocsTagGet(ctx context.Context, id int) (*dto.Tag, error)
	DBDocsTagList(ctx context.Context, pageSize, pageIndex int) (*[]dto.Tag, error)
	DBDocsTagUpdate(ctx context.Context, condition *dto.Tag) error

	DBDocsCategoryCreate(ctx context.Context, p *dto.Category) error
	DBDocsCategoryDelete(ctx context.Context, id int) error
	DBDocsCategoryGet(ctx context.Context, id int) (*dto.Category, error)
	DBDocsCategoryList(ctx context.Context, pageSize, pageIndex int) (*[]dto.Category, error)
	DBDocsCategoryUpdate(ctx context.Context, condition *dto.Category) error

	DBDocsPostTagCreate(ctx context.Context, p []*dto.PostTag) error
	DBDocsPostTagDelete(ctx context.Context, id int) error
	DBDocsPostTagGet(ctx context.Context, id int) (*dto.PostTag, error)
	DBDocsPostTagList(ctx context.Context, pageSize, pageIndex int) (*[]dto.PostTag, error)
	DBDocsPostTagUpdate(ctx context.Context, condition *dto.PostTagUpdate) error

	DBDocsSourceCreate(ctx context.Context, p *dto.Source) error
	DBDocsSourceDelete(ctx context.Context, id int) error
	DBDocsSourceGet(ctx context.Context, id int) (*dto.Source, error)
	DBDocsSourceList(ctx context.Context, pageSize, pageIndex int) (*[]dto.Source, error)
	DBDocsSourceUpdate(ctx context.Context, condition *dto.Source) error

	DBGetPostDetail(ctx context.Context, id string) (*dto.Post, error)
	DBGetTagList(ctx context.Context, pageSize, pageIndex int) (*[]dto.Tag, error)
	DBGetPostListByTag(ctx context.Context, tag string, pageSize, pageIndex int) (*[]dto.PostListByTag, error)
	DBGetSourceByPost(ctx context.Context, id string, pageSize, pageIndex int) (*[]dto.SourceByPost, error)
	DBGetPostListByCategory(ctx context.Context, parentCate, childCate string, pageSize, pageIndex int) *[]dto.PostListByCategory
	DBGetTagListByPost(ctx context.Context, id string) (*[]dto.TagListByPost, error)
	DBGetCategoryList(ctx context.Context, parent string) (*[]dto.CategoryList, error)
	DBCountPostsByCategory(ctx context.Context, parentCate, childCate string) *int
	DBGetInfoCategoryBySlug(ctx context.Context, parentCate, childCate string) (*dto.CategoryInfo, error)
	DBGetTagInfoBySlug(ctx context.Context, slug string) (*dto.TagInfo, error)
	DBGetPostsInSameCategory(ctx context.Context, id string, pageSize, pageIndex int) (*[]dto.PostsInSameCategory, error)
	DBGetPostListLatestCreated(ctx context.Context, pageSize, pageIndex int) (*[]dto.PostListLatestCreated, error)
	DBGetChildCategoryList(ctx context.Context, parentCate string, pageSize, pageIndex int) (*[]dto.GetChildCategoryList, error)
	DBGetCategoryListByHierarchy(ctx context.Context) (*[]dto.CategoryListByHierarchy, error)
	DBGetParentCategoryList(ctx context.Context) (*[]dto.ParentCategoryList, error)
}
