package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zhaiht/blog_backend/config"
	"github.com/zhaiht/blog_backend/model"
	"gorm.io/gorm"
)

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

func CreateComment(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		logrus.Warn("User ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	postID := c.Param("postId")
	var post model.Post
	if err := config.DB.First(&post, postID).Error; err != nil {
		logrus.Warnf("Post %s not found: %v", postID, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Warnf("Invalid create comment request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment := model.Comment{
		Content: req.Content,
		UserId:  userID.(uint),
		PostId:  post.ID,
	}

	if err := config.DB.Create(&comment).Error; err != nil {
		logrus.Errorf("Failed to create comment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	//加载用户信息
	config.DB.Preload("Users", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "UserName")
	}).First(&comment, comment.ID)

	c.JSON(http.StatusCreated, comment)

}

// 获取文章的所有评论
func GetAllComments(c *gin.Context) {
	postId := c.Param("postId")
	var comments []model.Comment
	if err := config.DB.Where("post_id = ?", postId).Preload("User",
		func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "UserName")
		}).Find(&comments).Error; err != nil {
		logrus.Errorf("Failed to get comments for post %s: %v", postId, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
