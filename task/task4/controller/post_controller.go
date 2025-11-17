package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zhaiht/blog_backend/config"
	"github.com/zhaiht/blog_backend/model"
	"gorm.io/gorm"
)

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// 创建文章
func CreatePost(c *gin.Context) {
	//只有被认证的用户才能创建文章
	userID, exists := c.Get("userID")
	if !exists {
		logrus.Warn("User ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Warnf("Invalid create post request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	post := model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserId:  userID.(uint),
	}

	if err := config.DB.Create(&post).Error; err != nil {
		logrus.Errorf("Failed to create post: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create post",
		})
		return
	}

	logrus.Infof("Post created by user: %d", userID)
	c.JSON(http.StatusCreated, post)
}

func GetAllPosts(c *gin.Context) {
	var posts []model.Post

	if err := config.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "UserName")
	}).Find(&posts).Error; err != nil {
		logrus.Errorf("Failed to get posts: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// 获取单篇文章
func GetPost(c *gin.Context) {
	id := c.Param("id")

	var post model.Post
	if err := config.DB.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "UserName")
	}).Preload("Comments.User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "UserName")
	}).Find(&post, id).Error; err != nil {
		logrus.Warnf("Post %s not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}
	c.JSON(http.StatusOK, post)
}

// 更新文章
func UpdatePost(c *gin.Context) {
	userID, exists := c.Get("userID")

	if !exists {
		logrus.Warn("User ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorizde"})
		return
	}

	id := c.Param("id")
	var post model.Post
	//查询文章
	if err := config.DB.First(&post, id).Error; err != nil {
		logrus.Warnf("Post %s not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	//检查是否是文章作者
	if post.UserId != userID.(uint) {
		logrus.Warnf("User %d tried to update post %s which then do not own", userID, id)
		c.JSON(http.StatusForbidden, gin.H{"error": "You dont have permession to update this post"})
		return
	}

	var req UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Warnf("Invalid update post request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//文章更新
	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Content != "" {
		post.Content = req.Content
	}

	if err := config.DB.Save(&post).Error; err != nil {
		logrus.Errorf("Failed to update post %s %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	logrus.Infof("Post %s updated by user %d", id, userID)
	c.JSON(http.StatusOK, post)
}

// 删除文章
func DeletePost(c *gin.Context) {
	userID, exists := c.Get("userID")

	if !exists {
		logrus.Warn("User ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorizde"})
		return
	}

	id := c.Param("id")
	var post model.Post
	//查询文章
	if err := config.DB.First(&post, id).Error; err != nil {
		logrus.Warnf("Post %s not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	//检查是否是文章作者
	if post.UserId != userID.(uint) {
		logrus.Warnf("User %d tried to update post %s which then do not own", userID, id)
		c.JSON(http.StatusForbidden, gin.H{"error": "You dont have permession to update this post"})
		return
	}

	if err := config.DB.Delete(&post).Error; err != nil {
		logrus.Errorf("Failed to delete post %s %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	logrus.Infof("Post %s deleted by user %d", id, userID)
	c.JSON(http.StatusOK, gin.H{"messgae": "Post deleted successfully"})

}
