// 返回token中间件，不用管，要看也行。模板写法来着，至于为什么，我也不是很懂
package middleware

import (
	"net/http"
	"sse_market_admin/common"
	"sse_market_admin/model"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		isWs := false //ws需要query获得token
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			tokenString = ctx.Query("token")
			isWs = true
		}
		//若不是ws连接并且token为空或token格式不正确
		if !isWs && (tokenString == "" || len(tokenString) <= 7 || !strings.HasPrefix(tokenString, "Bearer ")) {
			// 如果请求头中没有Authorization信息，或者信息不合法，则视为游客访问。
			ctx.Set("user", nil)
			ctx.Next()
			return
		}
		//当ws连接并且token小于7
		if isWs && len(tokenString) <= 7 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//若不是ws连接，表示有bearer
		if !isWs {
			tokenString = tokenString[7:]
		}
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//验证通过后获取claims中的userId
		userId := claims.UserID
		if userId == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		db := common.GetDB()
		user := model.User{}
		db.Where("userID = ?", userId).First(&user)
		//用户不存在
		if user.UserID == 0 {
			// 如果用户不存在，则将user设置为nil，表示当前用户为游客。
			ctx.Abort()
		} else {
			// 用户存在，将user的信息写入上下文
			ctx.Set("user", user)
		}
		ctx.Next()
	}
}

func AuthMiddleware_admin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		//若token为空或token格式不正确
		if tokenString == "" || len(tokenString) <= 7 || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			return
		}
		tokenString = tokenString[7:]
		token, claims_admin, err := common.ParseToken_admin(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//验证通过后获取claims中的userId
		adminId := claims_admin.AdminID
		if adminId == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		db := common.GetDB()
		admin := model.Admin{}
		db.Where("adminID = ?", adminId).First(&admin)
		//用户不存在
		if admin.AdminID == 0 {
			// 如果用户不存在，则将user设置为nil，表示当前用户为游客。
			ctx.Set("admin", nil)
		} else {
			// 用户存在，将user的信息写入上下文
			ctx.Set("admin", admin)
		}
		ctx.Next()
	}
}
