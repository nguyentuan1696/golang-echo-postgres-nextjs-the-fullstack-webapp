package docs

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"thichlab-backend-docs/infrastructure/cache"
	"thichlab-backend-docs/infrastructure/search"
	"thichlab-backend-docs/module/docs/controller"
	"thichlab-backend-docs/module/docs/repository"
	"thichlab-backend-docs/module/docs/service"
)

var mDocsController *controller.DocsController

func Initialize(e *echo.Echo, dbContext *sql.DB, sqlxDBContext *sqlx.DB, cache cache.Client, search search.Client) {

	DocsRepository := repository.NewDocsRepository(dbContext, sqlxDBContext)
	DocsService := service.NewDocsService(DocsRepository, cache, search)

	mDocsController = &controller.DocsController{
		DocsService: DocsService,
	}

	initRoute(e)

}

func initRoute(e *echo.Echo) {
	route := e.Group("/docs/api/v1/")
	privateRoute := e.Group("/private/docs/api/v1/")

	//Docs:Public

	route.GET("search", mDocsController.DocsSearch)
	route.GET("search/most-keyword", mDocsController.GetMostSearchedKeywords)

	route.GET("tag", mDocsController.GetTagList)
	route.GET("tag/:id", mDocsController.GetPostListByTag)
	route.GET("tag/info", mDocsController.GetInfoTag)

	route.GET("post/:id", mDocsController.GetPostDetail)
	route.GET("post/same-category", mDocsController.GetPostsInSameCategory)
	route.GET("post/latest", mDocsController.GetPostListLatestCreated)

	route.GET("category", mDocsController.GetListCategory)
	route.GET("category", mDocsController.GetListCategory)
	route.GET("category", mDocsController.GetListCategory)
	route.GET("category/post", mDocsController.GetPostListByCategory)
	route.GET("category/post-total", mDocsController.CountPostsByCategory)
	route.GET("category/info", mDocsController.GetInfoCategory)
	route.GET("category/child-list", mDocsController.GetChildCategoryList)
	route.GET("category/hierarchy", mDocsController.GetCategoryListByHierarchy)

	//Docs:Private
	privateRoute.POST("post", mDocsController.DocsPostCreate)
	privateRoute.DELETE("post", mDocsController.DocsPostDelete)
	privateRoute.GET("post/:id", mDocsController.DocsPostGet)
	privateRoute.GET("post", mDocsController.DocsPostList)
	privateRoute.PUT("post", mDocsController.DocsPostUpdate)

	privateRoute.POST("tag", mDocsController.DocsTagCreate)
	privateRoute.DELETE("tag", mDocsController.DocsTagDelete)
	privateRoute.GET("tag/:id", mDocsController.DocsTagGet)
	privateRoute.GET("tag", mDocsController.DocsTagList)
	privateRoute.PUT("tag", mDocsController.DocsTagUpdate)

	privateRoute.POST("category", mDocsController.DocsCategoryCreate)
	privateRoute.DELETE("category", mDocsController.DocsCategoryDelete)
	privateRoute.GET("category/:id", mDocsController.DocsCategoryGet)
	privateRoute.GET("category", mDocsController.DocsCategoryList)
	privateRoute.PUT("category", mDocsController.DocsCategoryUpdate)

	privateRoute.POST("post-tag", mDocsController.DocsPostTagCreate)
	privateRoute.DELETE("post-tag/:id", mDocsController.DocsPostTagDelete)
	privateRoute.GET("post-tag/:id", mDocsController.DocsPostTagGet)
	privateRoute.GET("post-tag", mDocsController.DocsPostTagList)
	privateRoute.PUT("post-tag", mDocsController.DocsPostTagUpdate)

	privateRoute.POST("source", mDocsController.DocsSourceCreate)
	privateRoute.DELETE("source/:id", mDocsController.DocsSourceDelete)
	privateRoute.GET("source/:id", mDocsController.DocsSourceGet)
	privateRoute.GET("source", mDocsController.DocsSourceList)
	privateRoute.PUT("source", mDocsController.DocsSourceUpdate)
}
