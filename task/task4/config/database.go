package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/zhaiht/blog_backend/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 初始化数据库
func InitDB() {

	//加载环境变量
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error Loading .env file")
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		logrus.Fatal("DB_DSN is not set in .env file")
	}

	//dsn := "root:1003200802@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		//panic("failed to connect database")
		logrus.Fatalf("failed to connect to database : %v", err)
	}

	if err := DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}); err != nil {
		logrus.Fatalf("Failed to migrate database: %v", err)
	}

	logrus.Info("Database connected and migated successfully")

}
