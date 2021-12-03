package controllers

import (
	"github.com/chrispappas/grafana-kitchen-sink/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.IndentedJSON(http.StatusOK, gin.H{"data": book})
}

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.IndentedJSON(http.StatusOK, gin.H{"data": books})
}
