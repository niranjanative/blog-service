package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/niranjanative/blog-service/controllers/articles"
	"github.com/niranjanative/blog-service/controllers/comments"
	"github.com/niranjanative/blog-service/models"
)

func main() {
	router := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	router.GET("/articles", articles.ListArticles)
	router.GET("/articles/:id", articles.GetArticleContent)
	router.POST("/articles", articles.CreateArticle)
	router.GET("/comments/:parent-id", comments.ListComments)
	router.POST("/comments", comments.CreateComment)

	// Run the server
	log.Fatal(router.Run(":3000"))
}
