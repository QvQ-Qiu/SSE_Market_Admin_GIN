package route

import (
	"sse_market_admin/controller"
	"sse_market_admin/middleware"

	"github.com/gin-gonic/gin"
)

// 创建路由

func CollectRoute(r *gin.Engine) *gin.Engine {

	r.Use(middleware.CORSMiddleware())
	// 这里把路由分成了两组，其中auth是需要token验证的，也就是需要用户登录，noauth是不需要token的，也就是不需要用户登录。
	auth := r.Group("")
	// noauth := r.Group("")

	auth.Use(middleware.AuthMiddleware())

	// 给管理员设置一个新的路由分组
	adminAuth := r.Group("")
	adminAuth.Use(middleware.AuthMiddleware_admin())
	//adminAuth.POST("/api1/auth/passUsers", controller.PassUsers)
	adminAuth.POST("/api1/auth/addAdmin", controller.AddAdmin)
	adminAuth.POST("/api1/auth/changePassword", controller.ChangeAdminPassword)
	adminAuth.POST("/api1/auth/deleteUser", controller.DeleteUser)
	adminAuth.POST("/api1/auth/deleteAdmin", controller.DeleteAdmin)
	adminAuth.POST("/api1/auth/showUsers", controller.ShowFilterUsers)
	r.POST("/api1/auth/adminLogin", controller.AdminLogin)
	adminAuth.GET("/api1/auth/admininfo", controller.AdminInfo)
	adminAuth.GET("/api1/auth/getSues", controller.GetSues)
	adminAuth.POST("/api1/auth/noViolation", controller.NoViolation)
	adminAuth.POST("/api1/auth/violation", controller.Violation)
	adminAuth.POST("/api1/auth/adminPost", controller.AdminPost)
	adminAuth.POST("/api1/auth/adminBrowse", controller.AdminGetPost)
	adminAuth.POST("/api1/auth/adminDeletePost", controller.AdminDeletePost)
	adminAuth.POST("/api1/auth/markHQPost", controller.MarkHQPost)
	adminAuth.POST("/api1/auth/removeHQPost", controller.RemoveHQPost)
	adminAuth.POST("/api1/auth/muteUser", controller.MuteUser)
	adminAuth.POST("/api1/auth/release", controller.Release)
	adminAuth.POST("/api1/auth/getKey", controller.GetKey)
	adminAuth.POST("/api1/auth/addKey", controller.AddKey)
	adminAuth.POST("/api1/auth/getfeedback", controller.Getfeedback)
	return r
}
