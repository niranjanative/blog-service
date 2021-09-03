package comments

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/niranjanative/blog-service/controllers"
	"github.com/niranjanative/blog-service/models"
	"github.com/niranjanative/blog-service/models/comments"
)

// ListComments finds all comments with given parent-id and parent-type
// GET /comments
func ListComments(c *gin.Context) {
	var (
		commentsList []comments.Comment
		parentID     uint
	)

	pageNumber, pageSize, er := controllers.Pagination(c)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	parentType := c.Query(controllers.ParentTypeQueryParam)
	if parentType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Required QueryParam parent-type"})
		return
	} else if parentType != controllers.CommentQueryParamValue && parentType != controllers.ArticleQueryParamValue {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for QueryParam " + controllers.ParentTypeQueryParam})
		return
	}

	parentIDstring := c.Params.ByName(controllers.ParentIDPathParam)
	if parentIDstring != "" {
		ID, err := strconv.Atoi(parentIDstring)
		if err == nil {
			parentID = uint(ID)
			if err = models.DB.Model(&comments.Comment{}).Where("parent_id = ?", parentID).Where("parent_type = ?", parentType).
				Limit(pageSize).Offset(pageNumber).Find(&commentsList).Error; err != nil || len(commentsList) == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Records not found!"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": commentsList})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Required PathParam " + controllers.ParentTypeQueryParam})
}

// CreateComment creates comments on given parent
// POST /comments
func CreateComment(c *gin.Context) {

	var input comments.Comment
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
