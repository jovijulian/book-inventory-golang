package middleware

import (
	"book-inventory-golang/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	_ "time"
)

func AuthValid(c *gin.Context) {
	var tokenString string
	tokenString = c.Query("auth")
	if tokenString == "" {
		tokenString = c.PostForm("auth")
		if tokenString == "" {
			c.HTML(http.StatusMovedPermanently, "login.html", gin.H{
				"content": "Silahkan login terlebih dahulu",
			})
		}
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, invalid := token.Method.(*jwt.SigningMethodHMAC); !invalid {
			return nil, fmt.Errorf("Invalid token, alg: %v", token.Header["alg"])
		}
		return []byte(models.SECRET), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"content": "Silahkan login kembali",
			})
			c.Abort()
			return
		}

		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				c.HTML(http.StatusUnauthorized, "login.html", gin.H{
					"content": "Silahkan login kembali",
				})
				c.Abort()
				return
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"content": "Silahkan login terlebih dahulu",
		})
		c.Abort()
		return
	}

	if token.Valid {
		fmt.Println("Token Verified")
		c.Next()
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"content": "Silahkan login terlebih dahulu",
		})
		c.Abort()
	}
}
