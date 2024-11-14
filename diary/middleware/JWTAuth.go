package middleware

import (
	"diary/controllers"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// JWT 中介軟體：驗證請求中的 JWT
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 從 Authorization 標頭中獲取 token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// 認證標頭應該是 "Bearer <token>"
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		// 提取 token
		tokenStr := bearerToken[1]

		// 解析並驗證 token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// 確保使用的是正確的簽名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// 返回密鑰
			return controllers.JWTSecret, nil
		})

		// 檢查 token 是否有效
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 從 token 中獲取用戶資料並存到上下文中
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			// 可以將用戶名、角色等信息存入上下文中
			c.Set("username", claims["username"])
		}

		// 繼續處理請求
		c.Next()
	}
}
