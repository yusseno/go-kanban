package middelware

import (
	"fmt"
	"go-kanban/security"
	"go-kanban/session"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Fullname   string
	Email      string
	Permission int
	jwt.StandardClaims
}

func AuthzUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			res := map[string]string{"message": "token not empty"}
			c.JSON(401, res)
			c.Abort()
			return
		}

		reqtoken := tokenString[7:]
		// fmt.Println(reqtoken)

		err := security.ValidateToken(reqtoken)
		if err != nil {
			res := map[string]string{"message": "token not valid"}
			c.JSON(401, res)
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(reqtoken, &security.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*security.JWTClaim)

		if !ok || !token.Valid || claims.Is_Role != 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "tidak ada akses"})
			c.Abort()
			return
		}

		ses, _ := session.SessionStore.Get(claims.Email)
		if ses.TokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		send := claims.Email
		c.Set("objek", send)
		c.Next()
	}
}

func AuthzAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			res := map[string]string{"message": "token not empty"}
			c.JSON(401, res)
			c.Abort()
			return
		}

		reqtoken := tokenString[7:]
		// fmt.Println(reqtoken)

		err := security.ValidateToken(reqtoken)
		if err != nil {
			res := map[string]string{"message": "token not valid"}
			c.JSON(401, res)
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(reqtoken, &security.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*security.JWTClaim)
		fmt.Println(claims)

		if !ok || !token.Valid || claims.Is_Role != 1 {
			c.JSON(http.StatusBadGateway, gin.H{"error": "tidak ada akses"})
			c.Abort()
			return
		}

		ses, _ := session.SessionStore.Get(claims.Email)
		if ses.TokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		send := claims.Email
		c.Set("objek", send)
		c.Next()
	}
}
