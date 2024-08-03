package main

import (
	apps "book_inventory/app"
	"book_inventory/auth"
	"book_inventory/db"
	"book_inventory/middleware"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	conn := db.InitDB()

	router := gin.Default()

	// load file static
	router.Static("/static", "./static")

	// load template html
	router.LoadHTMLGlob("templates/*")

	handler := apps.New(conn)

	// home
	router.GET("/", auth.HomeHandler)

	// login
	router.GET("/login", auth.LoginGethandler)
	router.POST("/login", auth.LoginPostHandler)

	//  get all books
	router.GET("/books", middleware.AuthValid, handler.GetBooks)
	router.GET("/book/:id", middleware.AuthValid, handler.GetBooksById)

	// add book
	router.GET("/addBooks", middleware.AuthValid, handler.AddBooks)
	router.POST("/book", middleware.AuthValid, handler.PostBooks)

	// update book
	router.GET("/updateBook/:id", middleware.AuthValid, handler.UpdateBooks)
	router.POST("/updateBook/:id", middleware.AuthValid, handler.PutBooks)

	// delete book
	router.POST("/deleteBook/:id", middleware.AuthValid, handler.DeleteBooks)

	router.Run()
}
