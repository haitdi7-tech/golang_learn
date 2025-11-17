package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zhaiht/blog_backend/config"
	"github.com/zhaiht/blog_backend/model"
	"github.com/zhaiht/blog_backend/util"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 用户注册
func Register(c *gin.Context) {
	//获取注册信息
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Warnf("Invalid register request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	//检查用户名是否已经存在
	var existingUser model.User
	if err := config.DB.Where("user_name = ?", req.Username).First(&existingUser).Error; err == nil {
		logrus.Warnf("Username %s already exists", req.Username)
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	//检查邮箱是否已经存在
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		logrus.Warnf("Email %s already exists", req.Email)
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	//密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Warnf("Failed to hash password:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register user"})
		return
	}

	//创建用户
	user := model.User{
		UserName: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		logrus.Errorf("Failed to create user:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	logrus.Infof("User %s reqistered sucessfully", req.Username)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User register created successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.UserName,
			"email":    user.Email,
		},
	})
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Warnf("Invalid login request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var user model.User
	if err := config.DB.Where("user_name = ?", req.Username).First(&user).Error; err != nil {
		logrus.Warnf("User %s not found", req.Username)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	//验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		logrus.Warnf("Invalid password for user %s", req.Username)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	//生成JWT
	token, err := util.GenerateToken(user.ID)
	if err != nil {
		logrus.Errorf("Failed to generate token:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}

	logrus.Infof("User %s logged in successfully", req.Username)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.UserName,
			"email":    user.Email,
		},
	})

}
