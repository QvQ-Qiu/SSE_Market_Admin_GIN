// 用于处理Web应用程序中的跨源资源共享（CORS）请求的中间件，不用管，要看也行。模板写法来着，至于为什么，我也不是很懂
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin","*")
		ctx.Writer.Header().Set("Access-Control-Max-Age","86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods","*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers","*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials","true")
		if ctx.Request.Method == http.MethodOptions{
			ctx.AbortWithStatus(200)
		}else{
			ctx.Next()
		}
	}
}