package repository

import (
	"context"
	"fmt"
	"strings"
	"thichlab-backend-docs/constant"
	"thichlab-backend-docs/dto"
	"thichlab-backend-docs/infrastructure/util"
)

//------------------------------------
// Docs:Public
//------------------------------------

func (repository *DocsRepository) DBGetParentCategoryList(ctx context.Context) (*[]dto.ParentCategoryList, error) {
	r := make([]dto.ParentCategoryList, 0)
	query := fmt.Sprintf("select c.id, c.title, c.description, c.parent, c.slug from category c where c.parent = ''")
	err := repository.Postgres.SQLxDBContext.SelectContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return &r, nil
}

func (repository *DocsRepository) DBGetCategoryListByHierarchy(ctx context.Context) (*[]dto.CategoryListByHierarchy, error) {
	r := make([]dto.CategoryListByHierarchy, 0)
	query := fmt.Sprintf("select c.id, c.title, c.description, c.parent, c.slug from category c")
	err := repository.Postgres.SQLxDBContext.SelectContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return &r, nil
}

func (repository *DocsRepository) DBGetChildCategoryList(ctx context.Context, parentCate string, pageSize, pageIndex int) (*[]dto.GetChildCategoryList, error) {
	r := make([]dto.GetChildCategoryList, 0)
	query := fmt.Sprintf("select c.id, c.title, c.slug from category c where c.parent = '%s' order by c.created_at limit %d offset %d", parentCate, pageSize, pageIndex*pageSize-pageSize)
	err := repository.Postgres.SQLxDBContext.SelectContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return &r, nil
}

func (repository *DocsRepository) DBGetPostListLatestCreated(ctx context.Context, pageSize, pageIndex int) (*[]dto.PostListLatestCreated, error) {
	r := make([]dto.PostListLatestCreated, 0)
	query := fmt.Sprintf("select p.id, p.title, p.slug from post p where p.status = %d order by p.created_at desc limit %d offset %d", constant.PostStatusPublished, pageSize, pageIndex*pageSize-pageSize)
	err := repository.Postgres.SQLxDBContext.SelectContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return &r, nil
}

func (repository *DocsRepository) DBGetPostsInSameCategory(ctx context.Context, id string, pageSize, pageIndex int) (*[]dto.PostsInSameCategory, error) {
	r := make([]dto.PostsInSameCategory, 0)
	query := fmt.Sprintf("select p.id, p.title, p.slug from post p where p.category_id in (select category_id from post where id = '%s') order by random() limit %d offset %d", id, pageSize, pageIndex*pageSize-pageSize)
	err := repository.Postgres.SQLxDBContext.SelectContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return &r, nil
}

func (repository *DocsRepository) DBGetTagInfoBySlug(ctx context.Context, slug string) (*dto.TagInfo, error) {
	r := &dto.TagInfo{}
	query := fmt.Sprintf("select t.id, t.title, t.description from tag t where t.slug = '%s'", slug)
	err := repository.Postgres.SQLxDBContext.GetContext(ctx, r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return r, nil
}

func (repository *DocsRepository) DBGetInfoCategoryBySlug(ctx context.Context, parentCate, childCate string) (*dto.CategoryInfo, error) {

	var conditions []string
	var conditionQuery string
	var r dto.CategoryInfo

	query := "select c.title, c.description from category c"

	if parentCate != constant.StringEmpty && childCate == constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("c.slug = '%s'", parentCate))
	}

	if childCate != constant.StringEmpty && parentCate != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("c.slug = '%s'", childCate))
	}

	if len(conditions) > constant.ValueEmpty {
		conditionQuery = ` WHERE ` + strings.Join(conditions, " AND ")
		query += conditionQuery
	}

	err := repository.Postgres.SQLxDBContext.GetContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, nil
	}
	return &r, nil
}

func (repository *DocsRepository) DBCountPostsByCategory(ctx context.Context, parentCate, childCate string) *int {
	var total int

	query := "select count(p.id) as total_post from category c inner join post p on c.id = p.category_id"
	var conditions []string
	var conditionQuery string

	if parentCate != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("c.slug in (select c.slug from category c where parent = '%s')", parentCate))
	}

	if childCate != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("c.slug = '%s'", childCate))
	}

	if len(conditions) > constant.ValueEmpty {
		conditionQuery = ` WHERE ` + strings.Join(conditions, " AND ")
		query += conditionQuery
	}

	err := repository.Postgres.SQLxDBContext.GetContext(ctx, &total, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil
	}
	return &total
}

func (repository *DocsRepository) DBGetCategoryList(ctx context.Context, parent string) (*[]dto.CategoryList, error) {
	r := make([]dto.CategoryList, 0)
	query := fmt.Sprintf("select c.id , c.title, c.description, c.slug, c.parent from category c where  parent = '%s'", parent)
	err := repository.Postgres.SQLxDBContext.SelectContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return &r, nil

}

func (repository *DocsRepository) DBGetTagListByPost(ctx context.Context, id string) (*[]dto.TagListByPost, error) {
	r := make([]dto.TagListByPost, 0)
	query := fmt.Sprintf("select t.id as tag_id, t.slug as slug_tag, t.title as title_tag, p.id as post_id from post_tag pt inner join tag t on pt.tag_id = t.id inner join post p on pt.post_id = p.id where pt.post_id = '%s'", id)
	err := repository.Postgres.SQLxDBContext.SelectContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return &r, nil
}

func (repository *DocsRepository) DBGetPostListByCategory(ctx context.Context, parentCate, childCate string, pageSize, pageIndex int) *[]dto.PostListByCategory {
	r := make([]dto.PostListByCategory, 0)
	query := "select p.id as id_post, p.title as title_post, p.slug as slug_post, c.title as category_post from category c inner join post p on c.id = p.category_id"

	var conditions []string
	var conditionQuery string

	if parentCate != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("c.slug in (select c.slug from category c where parent = '%s')", parentCate))
	}

	if childCate != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("c.slug = '%s'", childCate))
	}

	if len(conditions) > constant.ValueEmpty {
		conditionQuery = ` WHERE ` + strings.Join(conditions, " AND ")
		query += conditionQuery
	}

	query += fmt.Sprintf(" ORDER BY p.created_at DESC LIMIT %v OFFSET %v", pageSize, pageIndex*pageSize-pageSize)

	err := repository.Postgres.SQLxDBContext.SelectContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil
	}
	return &r

}

func (repository *DocsRepository) DBGetSourceByPost(ctx context.Context, id string, pageSize, pageIndex int) (*[]dto.SourceByPost, error) {
	r := make([]dto.SourceByPost, 0)
	query := fmt.Sprintf("select s.id as source_id, s.title source_title, s.url as source_url from post p join source s on p.id = s.post_id where p.id = '%s' and p.status = %d ORDER BY s.created_at DESC LIMIT %d OFFSET %d", id, constant.PostStatusPublished, pageSize, pageIndex*pageSize-pageSize)
	err := repository.Postgres.SQLxDBContext.SelectContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return &r, nil
}

func (repository *DocsRepository) DBGetPostListByTag(ctx context.Context, tag string, pageSize, pageIndex int) (*[]dto.PostListByTag, error) {
	r := make([]dto.PostListByTag, 0)
	query := fmt.Sprintf("select p.id, p.title, p.slug from post p join post_tag pt on p.id = pt.post_id join tag t on pt.tag_id = t.id where t.slug = '%s' and p.status = %d ORDER BY p.created_at DESC LIMIT %d OFFSET %d", tag, constant.PostStatusPublished, pageSize, pageIndex*pageSize-pageSize)
	err := repository.Postgres.SQLxDBContext.SelectContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return &r, nil
}

func (repository *DocsRepository) DBGetTagList(ctx context.Context, pageSize, pageIndex int) (*[]dto.Tag, error) {
	r := make([]dto.Tag, 0)
	query := fmt.Sprintf("SELECT id, title, description, slug FROM %s ORDER BY created_at DESC LIMIT %v OFFSET %v", constant.TableTag, pageSize, pageIndex*pageSize-pageSize)
	err := repository.Postgres.SQLxDBContext.SelectContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return &r, nil
}

func (repository *DocsRepository) DBGetPostDetail(ctx context.Context, id string) (*dto.Post, error) {
	r := &dto.Post{}
	query := fmt.Sprintf("SELECT id, title, description, content, slug FROM %s WHERE id = '%s' and status = %d",
		constant.TablePost, id, constant.PostStatusPublished,
	)
	err := repository.Postgres.SQLxDBContext.GetContext(ctx, r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return r, nil
}

//------------------------------------
// Docs:Post
//------------------------------------

func (repository *DocsRepository) DBDocsPostCreate(ctx context.Context, p *dto.Post) error {

	query := fmt.Sprintf("INSERT INTO %s (id, title, description, content, slug, status, created_at, updated_at, category_id, category_name, category_slug) VALUES ( '%s', '%s', '%s', '%s', '%s', %d, %v, %v, %d, '%s', '%s')",
		constant.TablePost, p.Id, p.Title, p.Description, p.Content, p.Slug, p.Status, util.NowUnixTimeMillisecond(), util.NowUnixTimeMillisecond(), p.CategoryId, p.CategoryName, p.CategorySlug,
	)

	_, err := repository.Postgres.SQLxDBContext.ExecContext(ctx, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return err
	}

	return nil
}

func (repository *DocsRepository) DBDocsPostDelete(ctx context.Context, id string) error {

	// If you delete a post, you will also delete the post by id in the post, post_tag, source tables
	query1 := fmt.Sprintf("DELETE FROM %s WHERE id = '%s'", constant.TablePost, id)
	query2 := fmt.Sprintf("DELETE FROM %s WHERE post_id = '%s'", constant.TablePostTag, id)
	query3 := fmt.Sprintf("DELETE FROM %s WHERE post_id = '%s'", constant.TableSource, id)

	tx, _ := repository.Postgres.SQLxDBContext.Begin()

	_, err := tx.ExecContext(ctx, query3)
	if err != nil {
		repository.Postgres.HandleError(err, query3)
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	_, err = tx.ExecContext(ctx, query2)
	if err != nil {
		repository.Postgres.HandleError(err, query1)
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	_, err = tx.ExecContext(ctx, query1)
	if err != nil {
		repository.Postgres.HandleError(err, query1)
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
	}

	return nil
}

func (repository *DocsRepository) DBDocsPostGet(ctx context.Context, id string) (*dto.Post, error) {
	r := &dto.Post{}
	query := fmt.Sprintf("SELECT id, title, description, content, slug, status, created_at, updated_at, category_id, category_name, category_slug FROM %s WHERE id = '%s'",
		constant.TablePost, id,
	)
	err := repository.Postgres.SQLxDBContext.GetContext(ctx, r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return r, nil
}

func (repository *DocsRepository) DBDocsPostList(ctx context.Context, status, pageSize, pageIndex int) (*[]dto.Post, error) {
	r := make([]dto.Post, 0)

	query := "SELECT id, title, description, content, slug, status, created_at, updated_at, category_id, category_name, category_slug FROM " + constant.TablePost

	var conditions []string
	var conditionQuery string

	if status > constant.ValueEmpty {
		conditions = append(conditions, fmt.Sprintf("status = %d", status))
	}

	if len(conditions) > constant.ValueEmpty {
		conditionQuery = ` WHERE ` + strings.Join(conditions, " AND ")
		query += conditionQuery
	}

	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT %v OFFSET %v", pageSize, pageIndex*pageSize-pageSize)

	err := repository.Postgres.SQLxDBContext.SelectContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return &r, nil
}

func (repository *DocsRepository) DBDocsPostUpdate(ctx context.Context, condition *dto.Post) error {
	var conditions []string

	if condition.Title != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("title = '%s'", condition.Title))
	}

	if condition.Description != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("description = '%s'", condition.Title))
	}

	if condition.Content != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("content = '%s'", condition.Content))
	}

	if condition.Slug != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("slug = '%s'", condition.Slug))
	}

	if condition.Status <= constant.ValueEmpty {
		conditions = append(conditions, fmt.Sprintf("status = %d'", condition.Status))
	}

	if condition.UpdatedAt <= constant.ValueEmpty {
		conditions = append(conditions, fmt.Sprintf("updated_at = %v", util.NowUnixTimeMillisecond()))
	}

	if condition.CategoryName != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("category_name = '%s'", condition.CategoryName))
	}

	if condition.CategorySlug != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("category_slug = '%s'", condition.CategorySlug))
	}

	if condition.CategoryId <= constant.ValueEmpty {
		conditions = append(conditions, fmt.Sprintf("category_id = %d", condition.CategoryId))
	}

	query := fmt.Sprintf("UPDATE %s", constant.TablePost)
	if len(conditions) != constant.ValueEmpty {
		query += " SET " + strings.Join(conditions, ", ")
	}

	query += fmt.Sprintf(" WHERE id = '%s'", condition.Id)

	_, err := repository.Postgres.SQLxDBContext.ExecContext(ctx, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return err
	}

	return nil
}

//------------------------------------
// Docs:Tag
//------------------------------------

func (repository *DocsRepository) DBDocsTagCreate(ctx context.Context, p *dto.Tag) error {

	query := fmt.Sprintf("INSERT INTO %s (title, description, slug, created_at, updated_at) VALUES (  '%s', '%s', '%s', %v, %v)",
		constant.TableTag, p.Title, p.Description, p.Slug, util.NowUnixTimeMillisecond(), util.NowUnixTimeMillisecond(),
	)

	_, err := repository.Postgres.SQLxDBContext.ExecContext(ctx, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return err
	}

	return nil
}

func (repository *DocsRepository) DBDocsTagDelete(ctx context.Context, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %v", constant.TableTag, id)
	_, err := repository.Postgres.SQLxDBContext.ExecContext(ctx, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return err
	}
	return nil
}

func (repository *DocsRepository) DBDocsTagGet(ctx context.Context, id int) (*dto.Tag, error) {
	r := &dto.Tag{}
	query := fmt.Sprintf("SELECT id, title, description, slug, created_at, updated_at FROM %s WHERE id = %v",
		constant.TableTag, id,
	)
	err := repository.Postgres.SQLxDBContext.GetContext(ctx, r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return r, nil
}

func (repository *DocsRepository) DBDocsTagList(ctx context.Context, pageSize, pageIndex int) (*[]dto.Tag, error) {
	r := make([]dto.Tag, 0)

	query := fmt.Sprintf("SELECT id, title, description, slug, created_at, updated_at FROM %s ORDER BY created_at DESC LIMIT %v OFFSET %v", constant.TableTag, pageSize, pageIndex*pageSize-pageSize)

	err := repository.Postgres.SQLxDBContext.SelectContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return &r, nil
}

func (repository *DocsRepository) DBDocsTagUpdate(ctx context.Context, condition *dto.Tag) error {
	var conditions []string

	if condition.Title != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("title = '%s'", condition.Title))
	}

	if condition.Description != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("description = '%s'", condition.Title))
	}

	if condition.Slug != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("slug = '%s'", condition.Slug))
	}

	if condition.UpdatedAt <= constant.ValueEmpty {
		conditions = append(conditions, fmt.Sprintf("updated_at = %v", util.NowUnixTimeMillisecond()))
	}

	query := fmt.Sprintf("UPDATE %s", constant.TableTag)
	if len(conditions) != constant.ValueEmpty {
		query += " SET " + strings.Join(conditions, ", ")
	}

	query += fmt.Sprintf(" WHERE id = %v", condition.Id)

	_, err := repository.Postgres.SQLxDBContext.ExecContext(ctx, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return err
	}

	return nil
}

//------------------------------------
// Docs:Category
//------------------------------------

func (repository *DocsRepository) DBDocsCategoryCreate(ctx context.Context, p *dto.Category) error {

	query := fmt.Sprintf("INSERT INTO %s (title, description, slug, parent, created_at, updated_at) VALUES (  '%s', '%s', '%s', '%s', %v, %v)",
		constant.TableCategory, p.Title, p.Description, p.Slug, p.Parent, util.NowUnixTimeMillisecond(), util.NowUnixTimeMillisecond(),
	)

	_, err := repository.Postgres.SQLxDBContext.ExecContext(ctx, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return err
	}

	return nil
}

func (repository *DocsRepository) DBDocsCategoryDelete(ctx context.Context, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %v", constant.TableCategory, id)
	_, err := repository.Postgres.SQLxDBContext.ExecContext(ctx, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return err
	}
	return nil
}

func (repository *DocsRepository) DBDocsCategoryGet(ctx context.Context, id int) (*dto.Category, error) {
	r := &dto.Category{}
	query := fmt.Sprintf("SELECT id, title, description, slug, parent, created_at, updated_at FROM %s WHERE id = %v",
		constant.TableCategory, id,
	)
	err := repository.Postgres.SQLxDBContext.GetContext(ctx, r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return r, nil
}

func (repository *DocsRepository) DBDocsCategoryList(ctx context.Context, pageSize, pageIndex int) (*[]dto.Category, error) {
	r := make([]dto.Category, 0)

	query := fmt.Sprintf("SELECT id, title, description, slug, parent, created_at, updated_at FROM %s ORDER BY created_at DESC LIMIT %v OFFSET %v", constant.TableCategory, pageSize, pageIndex*pageSize-pageSize)

	err := repository.Postgres.SQLxDBContext.SelectContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return &r, nil
}

func (repository *DocsRepository) DBDocsCategoryUpdate(ctx context.Context, condition *dto.Category) error {
	var conditions []string

	if condition.Title != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("title = '%s'", condition.Title))
	}

	if condition.Description != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("description = '%s'", condition.Title))
	}

	if condition.Slug != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("slug = '%s'", condition.Slug))
	}

	if condition.Parent != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("parent = '%s'", condition.Parent))
	}

	if condition.UpdatedAt <= constant.ValueEmpty {
		conditions = append(conditions, fmt.Sprintf("updated_at = %v", util.NowUnixTimeMillisecond()))
	}

	query := fmt.Sprintf("UPDATE %s", constant.TableCategory)
	if len(conditions) != constant.ValueEmpty {
		query += " SET " + strings.Join(conditions, ", ")
	}

	query += fmt.Sprintf(" WHERE id = %v", condition.Id)

	_, err := repository.Postgres.SQLxDBContext.ExecContext(ctx, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return err
	}

	return nil
}

//------------------------------------
// Docs:PostTag
//------------------------------------

func (repository *DocsRepository) DBDocsPostTagCreate(ctx context.Context, p []*dto.PostTag) error {

	query := `INSERT INTO post_tag(post_id, tag_id, created_at, updated_at) VALUES %s`

	var values []string
	for _, v := range p {
		values = append(values, fmt.Sprintf("('%s', %d,  %d, %d)",
			v.PostId, v.TagId, util.NowUnixTimeMillisecond(), util.NowUnixTimeMillisecond()))
	}
	query = fmt.Sprintf(query, strings.Join(values, ","))

	_, err := repository.Postgres.SQLxDBContext.ExecContext(ctx, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return err
	}

	return nil
}

func (repository *DocsRepository) DBDocsPostTagDelete(ctx context.Context, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %v", constant.TablePostTag, id)
	_, err := repository.Postgres.SQLxDBContext.ExecContext(ctx, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return err
	}
	return nil
}

func (repository *DocsRepository) DBDocsPostTagGet(ctx context.Context, id int) (*dto.PostTag, error) {
	r := &dto.PostTag{}
	query := fmt.Sprintf("SELECT id, post_id, tag_id, created_at, updated_at FROM %s WHERE id = %v",
		constant.TablePostTag, id,
	)
	err := repository.Postgres.SQLxDBContext.GetContext(ctx, r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return r, nil
}

func (repository *DocsRepository) DBDocsPostTagList(ctx context.Context, pageSize, pageIndex int) (*[]dto.PostTag, error) {
	r := make([]dto.PostTag, 0)

	query := fmt.Sprintf("SELECT id, post_id, tag_id, created_at, updated_at FROM %s ORDER BY created_at DESC LIMIT %v OFFSET %v", constant.TablePostTag, pageSize, pageIndex*pageSize-pageSize)

	err := repository.Postgres.SQLxDBContext.SelectContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return &r, nil
}

func (repository *DocsRepository) DBDocsPostTagUpdate(ctx context.Context, condition *dto.PostTagUpdate) error {
	var conditions []string

	if condition.PostId != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("post_id = %v", condition.PostId))
	}

	if condition.TagId >= constant.ValueEmpty {
		conditions = append(conditions, fmt.Sprintf("tag_id = %v", condition.TagId))
	}

	if condition.UpdatedAt <= constant.ValueEmpty {
		conditions = append(conditions, fmt.Sprintf("updated_at = %v", util.NowUnixTimeMillisecond()))
	}

	query := fmt.Sprintf("UPDATE %s", constant.TablePostTag)
	if len(conditions) != constant.ValueEmpty {
		query += " SET " + strings.Join(conditions, ", ")
	}

	query += fmt.Sprintf(" WHERE id = %v", condition.Id)

	_, err := repository.Postgres.SQLxDBContext.ExecContext(ctx, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return err
	}

	return nil
}

//------------------------------------
// Docs:Source
//------------------------------------

func (repository *DocsRepository) DBDocsSourceCreate(ctx context.Context, p *dto.Source) error {

	query := fmt.Sprintf("INSERT INTO %s (title, url, post_id, created_at, updated_at) VALUES ( '%s', '%s', '%s' , %v, %v)",
		constant.TableSource, p.Title, p.Url, p.PostId, util.NowUnixTimeMillisecond(), util.NowUnixTimeMillisecond(),
	)

	_, err := repository.Postgres.SQLxDBContext.ExecContext(ctx, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return err
	}

	return nil
}

func (repository *DocsRepository) DBDocsSourceDelete(ctx context.Context, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %v", constant.TableSource, id)
	_, err := repository.Postgres.SQLxDBContext.ExecContext(ctx, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return err
	}
	return nil
}

func (repository *DocsRepository) DBDocsSourceGet(ctx context.Context, id int) (*dto.Source, error) {
	r := &dto.Source{}
	query := fmt.Sprintf("SELECT id, title, url, post_id, created_at, updated_at FROM %s WHERE id = %v",
		constant.TableSource, id,
	)
	err := repository.Postgres.SQLxDBContext.GetContext(ctx, r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return r, nil
}

func (repository *DocsRepository) DBDocsSourceList(ctx context.Context, pageSize, pageIndex int) (*[]dto.Source, error) {
	r := make([]dto.Source, 0)

	query := fmt.Sprintf("SELECT id, title, url, post_id, created_at, updated_at FROM %s ORDER BY created_at DESC LIMIT %v OFFSET %v", constant.TableSource, pageSize, pageIndex*pageSize-pageSize)

	err := repository.Postgres.SQLxDBContext.SelectContext(ctx, &r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return &r, nil
}

func (repository *DocsRepository) DBDocsSourceUpdate(ctx context.Context, condition *dto.Source) error {
	var conditions []string

	if condition.Title != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("title = '%s'", condition.Title))
	}

	if condition.Url != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("url = '%s'", condition.Url))
	}

	if condition.PostId != constant.StringEmpty {
		conditions = append(conditions, fmt.Sprintf("post_id = '%s'", condition.PostId))
	}

	if condition.UpdatedAt <= constant.ValueEmpty {
		conditions = append(conditions, fmt.Sprintf("updated_at = %v", util.NowUnixTimeMillisecond()))
	}

	query := fmt.Sprintf("UPDATE %s", constant.TableSource)
	if len(conditions) != constant.ValueEmpty {
		query += " SET " + strings.Join(conditions, ", ")
	}

	query += fmt.Sprintf(" WHERE id = %v", condition.Id)

	_, err := repository.Postgres.SQLxDBContext.ExecContext(ctx, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return err
	}

	return nil
}
