package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuedun/ginode-mongo/middleware"
	"github.com/yuedun/ginode-mongo/pkg/component"
	"github.com/yuedun/ginode-mongo/pkg/page"
	"github.com/yuedun/ginode-mongo/pkg/user"
	"github.com/yuedun/ginode-mongo/pkg/website"
)

/**
 * 路由注册
 */
func Register(router *gin.Engine) {
	userRouter := router.Group("/user")
	//user路由注册,可以给各个group加中间件
	userRouter.Use(middleware.Logger())
	{
		//userRouter.POST("/login", user.Login)
		userRouter.POST("/login", middleware.Jwt().LoginHandler)
		userRouter.GET("/refresh_token", middleware.Jwt().RefreshHandler) // 刷新token
		userRouter.GET("/logout", middleware.Jwt().LogoutHandler)
		userRouter.GET("/info/:id", middleware.Jwt().MiddlewareFunc(), user.GetUserInfo) //单独给某个路由添加中间件
		userRouter.POST("/", user.CreateUser)
		userRouter.GET("/list", user.GetUserList)
		userRouter.PUT("/update/:id", user.UpdateUser)
		userRouter.DELETE("/del/:id", user.DeleteUser)
	}
	//website路由注册
	websiteRouter := router.Group("/website")
	websiteRouter.Use(middleware.Jwt().MiddlewareFunc())
	{
		websiteRouter.GET("/", website.WebsiteList)
		websiteRouter.POST("/create", website.Create)
		websiteRouter.PUT("/update", website.Update)
		websiteRouter.DELETE("/delete/:id", website.Delete)
	}
	//page路由注册
	pageRouter := router.Group("/page")
	pageRouter.Use(middleware.Jwt().MiddlewareFunc())
	{
		pageRouter.GET("/list", page.List)
		pageRouter.POST("/create/:websiteId", page.Create)
		pageRouter.PUT("/update", page.Update)
		pageRouter.DELETE("/delete/:id", page.Delete)
		pageRouter.GET("/getPageComponents/:id", page.GetPageComponents)
		pageRouter.PUT("/updatePageComponents/:id", page.UpdatePageComponents)
		pageRouter.GET("/copyPage/:id/:url", page.CopyPage)
	}
	//component路由注册
	componentRouter := router.Group("/component")
	componentRouter.Use(middleware.Jwt().MiddlewareFunc())
	{
		componentRouter.GET("/", component.ComponentList)
		componentRouter.GET("/getComponent/:id", component.GetComponent)
		componentRouter.POST("/create", component.Create)
		componentRouter.PUT("/update", component.Update)
		componentRouter.DELETE("/delete/:id", component.Delete)
	}
	//website路由注册
	websiteAPI := router.Group("/api")
	{
		websiteSubAPI := websiteAPI.Group("/website")
		websiteSubAPI.GET("/", website.WebsiteList)
		websiteSubAPI.GET("/get-website", website.GetWebsite)
		websiteSubAPI.POST("/create", website.Create)
		websiteSubAPI.PUT("/update", website.Update)
		websiteSubAPI.DELETE("/delete/:id", website.Delete)
	}
}
