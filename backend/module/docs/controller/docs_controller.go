package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"thichlab-backend-docs/dto"
	"thichlab-backend-docs/gerror"
	"thichlab-backend-docs/infrastructure/response"
	"thichlab-backend-docs/infrastructure/util"
)

//----------------------------------
// Docs:Search
//----------------------------------

func (controller *DocsController) DocsSearch(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.DocsSearch(ctx, c.QueryParam("query"), util.ToInt(c.QueryParam("pageIndex")), util.ToInt(c.QueryParam("pageSize")))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) GetMostSearchedKeywords(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	// TODO: run cronjob delete key after a month
	// TODO: There will be cases of spam keywords, the solution is to transfer it to the management cms, select keywords to display on the web.
	r := controller.DocsService.GetMostSearchedKeywords(ctx, util.ToInt64(c.QueryParam("pageSize")), util.ToInt64(c.QueryParam("pageIndex")))
	return controller.StatusOkResponse(c, r)

}

//----------------------------------
// Docs:Public
//----------------------------------

func (controller *DocsController) GetCategoryListByHierarchy(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.GetCategoryListByHierarchy(ctx)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) DocsUploadImage(c echo.Context) error {

	return nil
}

func (controller *DocsController) GetChildCategoryList(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.GetChildCategoryList(ctx, c.QueryParam("parentCate"), util.ToInt(c.QueryParam("pageSize")), util.ToInt(c.QueryParam("pageIndex")))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) GetPostListLatestCreated(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.GetPostListLatestCreated(ctx, util.ToInt(c.QueryParam("pageSize")), util.ToInt(c.QueryParam("pageIndex")))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}
	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) GetPostsInSameCategory(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.GetPostsInSameCategory(ctx, c.QueryParam("id"), util.ToInt(c.QueryParam("pageSize")), util.ToInt(c.QueryParam("pageIndex")))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}
	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) GetInfoTag(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.GetTagInfoBySlug(ctx, c.QueryParam("slug"))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) GetInfoCategory(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.GetInfoCategoryBySlug(ctx, c.QueryParam("parentCate"), c.QueryParam("childCate"))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) CountPostsByCategory(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r := controller.DocsService.CountPostsByCategory(ctx, c.QueryParam("parentCate"), c.QueryParam("childCate"))

	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) GetListCategory(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.GetCategoryList(ctx, c.QueryParam("parent"))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) GetPostListByCategory(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r := controller.DocsService.GetPostListByCategory(ctx, c.QueryParam("parentCate"), c.QueryParam("childCate"), util.ToInt(c.QueryParam("pageSize")), util.ToInt(c.QueryParam("pageIndex")))

	return controller.StatusOkResponse(c, r)

}

func (controller *DocsController) GetPostListByTag(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.DocsGetPostListByTag(ctx, c.Param("id"), util.ToInt(c.QueryParam("pageSize")), util.ToInt(c.QueryParam("pageIndex")))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) GetTagList(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.DocsGetTagList(ctx, util.ToInt(c.QueryParam("pageSize")), util.ToInt(c.QueryParam("pageIndex")))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) GetPostDetail(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.DocsGetPostDetail(ctx, c.Param("id"))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	//TODO: Recommend System

	return controller.StatusOkResponse(c, r)
}

//----------------------------------
// Docs:Post
//----------------------------------

func (controller *DocsController) DocsPostCreate(c echo.Context) error {

	var err error
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	p := new(dto.Post)

	err = c.Bind(&p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), util.FuncName())
		return controller.StatusBadRequestResponse(c, message, resp)
	}

	err = controller.DocsService.DocsPostCreate(ctx, p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorSaveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, nil)
}

func (controller *DocsController) DocsPostDelete(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := controller.DocsService.DocsPostDelete(ctx, c.QueryParam("id"), c.QueryParam("categorySlug"))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorSaveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, nil)
}

func (controller *DocsController) DocsPostGet(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.DocsPostGet(ctx, c.Param("id"))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusBadRequestResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) DocsPostList(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.DocsPostList(ctx, util.ToInt(c.QueryParam("status")), util.ToInt(c.QueryParam("pageSize")), util.ToInt(c.QueryParam("pageIndex")))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}
	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) DocsPostUpdate(c echo.Context) error {
	var err error
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var p *dto.Post

	err = c.Bind(&p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), util.FuncName())
		return controller.StatusBadRequestResponse(c, message, resp)
	}

	err = controller.DocsService.DocsPostUpdate(ctx, p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorSaveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, nil)
}

//----------------------------------
// Docs:Tag
//----------------------------------

func (controller *DocsController) DocsTagCreate(c echo.Context) error {

	var err error
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	p := new(dto.Tag)

	err = c.Bind(&p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), util.FuncName())
		return controller.StatusBadRequestResponse(c, message, resp)
	}

	err = controller.DocsService.DocsTagCreate(ctx, p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorSaveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, nil)
}

func (controller *DocsController) DocsTagDelete(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := controller.DocsService.DocsTagDelete(ctx, util.ToInt(c.QueryParam("id")), c.QueryParam("slug"))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorSaveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, nil)
}

func (controller *DocsController) DocsTagGet(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.DocsTagGet(ctx, util.ToInt(c.Param("id")))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusBadRequestResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) DocsTagList(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.DocsTagList(ctx, util.ToInt(c.QueryParam("pageSize")), util.ToInt(c.QueryParam("pageIndex")))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}
	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) DocsTagUpdate(c echo.Context) error {
	var err error
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var p *dto.Tag

	err = c.Bind(&p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), util.FuncName())
		return controller.StatusBadRequestResponse(c, message, resp)
	}

	err = controller.DocsService.DocsTagUpdate(ctx, p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorSaveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, nil)
}

//----------------------------------
// Docs:Category
//----------------------------------

func (controller *DocsController) DocsCategoryCreate(c echo.Context) error {

	var err error
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	p := new(dto.Category)

	err = c.Bind(&p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), util.FuncName())
		return controller.StatusBadRequestResponse(c, message, resp)
	}

	err = controller.DocsService.DocsCategoryCreate(ctx, p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorSaveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, nil)
}

func (controller *DocsController) DocsCategoryDelete(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := controller.DocsService.DocsCategoryDelete(ctx, util.ToInt(c.QueryParam("id")), c.QueryParam("slug"))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorSaveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, nil)
}

func (controller *DocsController) DocsCategoryGet(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.DocsCategoryGet(ctx, util.ToInt(c.Param("id")))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusBadRequestResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) DocsCategoryList(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.DocsCategoryList(ctx, util.ToInt(c.QueryParam("pageSize")), util.ToInt(c.QueryParam("pageIndex")))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}
	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) DocsCategoryUpdate(c echo.Context) error {
	var err error
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var p *dto.Category

	err = c.Bind(&p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), util.FuncName())
		return controller.StatusBadRequestResponse(c, message, resp)
	}

	err = controller.DocsService.DocsCategoryUpdate(ctx, p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorSaveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, nil)
}

//----------------------------------
// Docs:PostTag
//----------------------------------

func (controller *DocsController) DocsPostTagCreate(c echo.Context) error {

	var err error
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var p []*dto.PostTag

	err = c.Bind(&p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), util.FuncName())
		return controller.StatusBadRequestResponse(c, message, resp)
	}

	err = controller.DocsService.DocsPostTagCreate(ctx, p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorSaveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, nil)
}

func (controller *DocsController) DocsPostTagDelete(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := controller.DocsService.DocsPostTagDelete(ctx, util.ToInt(c.Param("id")))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorSaveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, nil)
}

func (controller *DocsController) DocsPostTagGet(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.DocsPostTagGet(ctx, util.ToInt(c.Param("id")))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusBadRequestResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) DocsPostTagList(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.DocsPostTagList(ctx, util.ToInt(c.QueryParam("pageSize")), util.ToInt(c.QueryParam("pageIndex")))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}
	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) DocsPostTagUpdate(c echo.Context) error {
	var err error
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var p *dto.PostTagUpdate

	err = c.Bind(&p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), util.FuncName())
		return controller.StatusBadRequestResponse(c, message, resp)
	}

	err = controller.DocsService.DocsPostTagUpdate(ctx, p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorSaveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, nil)
}

//----------------------------------
// Docs:Source
//----------------------------------

func (controller *DocsController) DocsSourceCreate(c echo.Context) error {

	var err error
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	p := new(dto.Source)
	err = c.Bind(&p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), util.FuncName())
		return controller.StatusBadRequestResponse(c, message, resp)
	}

	err = controller.DocsService.DocsSourceCreate(ctx, p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorSaveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, nil)
}

func (controller *DocsController) DocsSourceDelete(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := controller.DocsService.DocsSourceDelete(ctx, util.ToInt(c.Param("id")), c.QueryParam("postId"))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorSaveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, nil)
}

func (controller *DocsController) DocsSourceGet(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.DocsSourceGet(ctx, util.ToInt(c.Param("id")))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusBadRequestResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) DocsSourceList(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r, err := controller.DocsService.DocsSourceList(ctx, util.ToInt(c.QueryParam("pageSize")), util.ToInt(c.QueryParam("pageIndex")))
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorRetrieveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}
	return controller.StatusOkResponse(c, r)
}

func (controller *DocsController) DocsSourceUpdate(c echo.Context) error {
	var err error
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var p *dto.Source
	err = c.Bind(&p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), util.FuncName())
		return controller.StatusBadRequestResponse(c, message, resp)
	}

	err = controller.DocsService.DocsSourceUpdate(ctx, p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorSaveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, nil)
}
