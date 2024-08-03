package app

import (
	"book_inventory/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{DB: db}
}

func (h *Handler) GetBooks(c *gin.Context) {
	var books []models.Books

	h.DB.Find(&books)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Home Page",
		"payload": books,
		"auth":    c.Query("auth"),
	})
}

func (h *Handler) GetBooksById(c *gin.Context) {
	bookId := c.Param("id")

	var books models.Books
	if h.DB.Find(&books, "id=?", bookId).RecordNotFound() {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.HTML(http.StatusOK, "books.html",
		gin.H{
			"title":   books.Title,
			"payload": books,
			"auth":    c.Query("auth"),
		},
	)
}

func (h *Handler) AddBooks(c *gin.Context) {
	c.HTML(http.StatusOK, "formBook.html", gin.H{
		"title": "Add Book",
		"auth":  c.Query("auth"),
	})

}

func (h *Handler) PostBooks(c *gin.Context) {
	var books models.Books

	c.Bind(&books)
	h.DB.Create(&books)
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/books?auth=%s", c.PostForm("auth")))
}

func (h *Handler) UpdateBooks(c *gin.Context) {
	var books models.Books

	bookId := c.Param("id")
	if h.DB.Find(&books, "id=?", bookId).RecordNotFound() {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "not found"})
	}

	c.HTML(http.StatusOK, "formBook.html", gin.H{
		"title":   "Add Book",
		"payload": books,
		"auth":    c.Query("auth"),
	})

}

func (h *Handler) PutBooks(c *gin.Context) {
	var books models.Books

	bookId := c.Param("id")
	if h.DB.Find(&books, "id=?", bookId).RecordNotFound() {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "not found"})
	}

	var reqBook = books
	c.Bind(&reqBook)

	h.DB.Model(&books).Where("id=?", bookId).Update(reqBook)

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/book/%s?auth=%s", bookId, c.PostForm("auth")))
}

func (h *Handler) DeleteBooks(c *gin.Context) {
	var books models.Books

	bookId := c.Param("id")
	h.DB.Delete(&books, "id=?", bookId)

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/books?auth=%s", c.PostForm("auth")))
}
