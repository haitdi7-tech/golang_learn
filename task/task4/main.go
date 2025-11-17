package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/zhaiht/blog_backend/config"
	"github.com/zhaiht/blog_backend/controller"
	"github.com/zhaiht/blog_backend/middleware"
)

func main() {
	//加载环境变量
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error Loading .evn file")
	}

	//初始化数据库
	config.InitDB()
	defer func() {
		sqlDB, err := config.DB.DB()
		if err != nil {
			logrus.Errorf("Failed to get DB instance: %v", err)
		} else {
			sqlDB.Close()
		}
	}()

	//设置日志格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	//初始化Gin
	r := gin.Default()

	//设置公共路由
	public := r.Group("/api")
	{
		//用户注册和登录
		public.POST("/register", controller.Register)
		public.POST("/login", controller.Login)

		//获取文章
		public.GET("/posts", controller.GetAllPosts)
		public.GET("/posts/:postId/comments", controller.GetAllComments)
		public.GET("/posts/post/:id", controller.GetPost)

	}

	//需要认证的路由
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		//文章管理
		protected.POST("/posts", controller.CreatePost)
		protected.PUT("/posts/:id", controller.UpdatePost)
		protected.DELETE("/posts/:id", controller.DeletePost)

		//评论管理
		protected.POST("/posts/:postId/comments", controller.CreateComment)
	}

	//启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logrus.Infof("Server is running on port %s", port)
	r.Run(":" + port)
}

