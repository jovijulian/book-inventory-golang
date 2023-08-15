package auth

import (
	"book-inventory-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"net/url"
	"time"
)

func HomeHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func LoginGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"content": "",
	})
}

func LoginPostHandler(c *gin.Context) {
	var credential models.Login
	err := c.Bind(&credential)
	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"content": err.Error(),
		})
	}

	if credential.Username != models.USER || credential.Password != models.PASSWORD {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"content": "Email atau Password salah",
		})
		return
	}

	claim := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		Issuer:    "Book Inventory",
		IssuedAt:  time.Now().Unix(),
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := sign.SignedString([]byte(models.SECRET))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"error": err.Error(),
		})
		c.Abort()
	}

	q := url.Values{}
	q.Set("auth", token)
	location := url.URL{Path: "/books", RawQuery: q.Encode()}
	c.Redirect(http.StatusMovedPermanently, location.RequestURI())
}
