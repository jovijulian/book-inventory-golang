package main

import (
	apps "book-inventory-golang/app"
	"book-inventory-golang/auth"
	"book-inventory-golang/db"
	"book-inventory-golang/middleware"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	conn := db.InitDb()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	handler := apps.New(conn)

	// home
	router.GET("/", auth.HomeHandler)

	// login
	router.GET("/login", auth.LoginGetHandler)
	router.POST("/login", auth.LoginPostHandler)

	// get all books
	router.GET("/books", middleware.AuthValid, handler.GetBooks)
	router.GET("/book/:id", middleware.AuthValid, handler.GetBookById)

	// add book
	router.GET("/addBook", middleware.AuthValid, handler.AddBook)
	router.POST("/book", middleware.AuthValid, handler.PostBook)

	// update book
	router.GET("/updateBook/:id", middleware.AuthValid, handler.UpdateBook)
	router.POST("/updateBook/:id", middleware.AuthValid, handler.PutBook)

	// delete book
	router.POST("/deleteBook/:id", middleware.AuthValid, handler.DeleteBook)

	router.Run()
}
