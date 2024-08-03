package middleware

import (
	"book_inventory/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthValid(c *gin.Context) {
	var tokenstring string
	tokenstring = c.Query("auth")
	if tokenstring == "" {
		tokenstring = c.PostForm("auth")
		if tokenstring == "" {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "token nil"})
			c.Abort()
		}
	}

	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, invalid := token.Method.(*jwt.SigningMethodHMAC); !invalid {
			return nil, fmt.Errorf("Invalid token ", token.Header["alg"])
		}
		return []byte(models.SECRET), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
		c.Next()
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "token is expiry"})
		c.Abort()
	}
}
