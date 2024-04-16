package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"ozinshe/pkg/entity"
	"strings"
)

func (h *Handler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if c.Request.URL.Path == "/auth/sign-in" || c.Request.URL.Path == "/auth/sign-up" || c.Request.URL.Path == "/auth/verifyAccount" {
		//	c.Next()
		//	return
		//}
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			h.log.Printf("Error in AuthMiddleware: %v", "Authorization header is missing")
			h.WriteHTTPResponse(c, http.StatusUnauthorized, "Authorization header is missing")
			//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			h.log.Printf("Error in AuthMiddleware: %v", "Invalid Authorization header format")
			h.WriteHTTPResponse(c, http.StatusUnauthorized, "Invalid Authorization header format")
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		tokenStr := parts[1]

		tkn, err := jwt.ParseWithClaims(tokenStr, &entity.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return entity.JWTKey, nil
		})
		if err != nil {
			if err.Error() == jwt.ErrSignatureInvalid.Error() {
				h.log.Printf("Error in AuthMiddleware: %v", err)
				h.WriteHTTPResponse(c, http.StatusUnauthorized, err.Error())
				return
			}
			h.log.Printf("Error in AuthMiddleware: %v", err)
			h.WriteHTTPResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		if !tkn.Valid {
			h.log.Printf("Error in AuthMiddleware: %v", err)
			h.WriteHTTPResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		decodedClaims := tkn.Claims.(*entity.Claims)
		c.Set("decodedClaims", decodedClaims)
		c.Next()
	}

}

func (h *Handler) AdminRoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.Value("decodedClaims").(*entity.Claims)
		if claims.Role != "admin" {
			h.log.Print("Error in AdminMiddleware: User is not admin!")
			h.WriteHTTPResponse(c, http.StatusUnauthorized, "User is not admin!")
			return
		}
		c.Next()
	}
}
