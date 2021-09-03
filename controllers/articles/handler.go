package articles

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/niranjanative/blog-service/controllers"
	"github.com/niranjanative/blog-service/models"
	"github.com/niranjanative/blog-service/models/articles"
)

// ListArticles Finds all articles without content
// GET /articles
func ListArticles(c *gin.Context) {
	var articleList []articles.Article

	pageNumber, pageSize, er := controllers.Pagination(c)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	if err := models.DB.Model(&articles.Article{}).Select("id, nickname, created_at, title").Limit(pageSize).
		Offset(pageNumber).Find(&articleList).Error; err != nil || len(articleList) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Records not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": articleList})
}

// GetArticleContent Finds particular article wi	th content
// GET /articles:id
func GetArticleContent(c *gin.Context) {
	var (
		article articles.Article
		ID      uint
	)

	IDstring := c.Params.ByName(controllers.ArticleIDPathParam)
	if IDstring != "" {
		IDint, err := strconv.Atoi(IDstring)
		if err == nil {
			ID = uint(IDint)
			if err = models.DB.Model(&articles.Article{}).Where("id = ?", ID).Find(&article).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": article})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PathParam " + controllers.ArticleIDPathParam})
}

// CreateArticle creates article
// POST /articles
func CreateArticle(c *gin.Context) {

	var input articles.Article
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}
