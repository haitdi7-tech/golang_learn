package main

import (
	"fmt"

	//"github.com/go-sql-driver/mysql"
	//"github.com/jmoiron/sqlx"
	//"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	//fmt.Println("Hello, world!")
	dsn := "root:1003200802@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		fmt.Println("连接失败")
		return
	}

	//创建表：USER,POST,COMMENT
	// db.AutoMigrate(&User{})
	//db.AutoMigrate(&Post{})
	//db.AutoMigrate(&Comment{})
	//db.AutoMigrate(&User{}, &Post{}, &Comment{})

	//创建一个用户一篇文章和一个评论
	// pk := db.Create(&User{
	// 	//Model: gorm.Model{ID: 1},
	// 	Name: "u1",
	// 	Posts: []Post{
	// 		{Title: " 测试标题",
	// 			Detail: "2内容1",
	// 			Comments: []Comment{
	// 				{Detail: "2赞同1"},
	// 			},
	// 		},
	// 	},
	// })

	// pk := db.Create(&Post{
	// 	Title:  "测试标题2",
	// 	Detail: "内容2",
	// 	UserId: 1,
	// 	Comments: []Comment{
	// 		{Detail: "评论1"},
	// 		{Detail: "评论2"},
	// 	},
	// })

	//钩子函数中没有删除的对象的详细信息
	// pk := db.Delete(&Comment{Model: gorm.Model{ID: 2}}).Error
	comment := Comment{}
	pk := db.Clauses(clause.Returning{}).Delete(&comment, 2).Error
	// var comment Comment
	// pk := db.First(&comment, 3).Error
	// pk = db.Delete(&comment, 3).Error
	fmt.Printf("返回的类型是：%v", pk)

	// var u = GetPostInfo(db)
	// fmt.Printf("返回的类型是：%v", *u)

	{
		// user, err := GetAllPostAndComment(db)

		// if err != nil {
		// 	fmt.Print("查询失败：%v \n", err)
		// 	return
		// }

		// for _, post := range user.Posts {
		// 	for _, comment := range post.Comments {
		// 		fmt.Printf("User: %v | Post: %v | Comment: %v", user.Name, post.Detail, comment.Detail)
		// 	}
		// }
	}
}

type Employee struct {
	id         string `db:"id"`
	name       string `db:"name"`
	department string `db:"department"`
	salary     string `db:"salary"`
}

// 使用SQLx查询数据
//1)编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中
//2)编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

// func SearchEmployees() {
// 	var epls []Employee
// 	db, error := sqlx.Connect("Mysql", "user:password@tcp(127.0.0.1:3306)/test?parsetime=true")
// 	if error != nil {
// 		fmt.Println("连接失败")
// 		return
// 	}
// 	defer db.Close()
// 	db.Select(&epls, "Select id,name,department,salary form employees where department = ?", "技术部")
// 	var maxSalaryE Employee
// 	query := `
// 			select id,name,department,salary
// 			form employees
// 			where salary = (selec max(salary) form employees)
// 			limit 1
// 			`
// 	err := db.Select(&maxSalaryE, query)
// 	if err != nil {
// 		fmt.Printf("查询失败： %v \n", err)
// 		return
// 	}
// }

// id 、 title 、 author 、 price
type Book struct {
	id     int
	title  string
	author string
	price  float64
}

//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全
// func SearchBooks() {
// 	var bks []Book
// 	db, error := sqlx.Connect("Mysql", "user:password@tcp(127.0.0.1:3306)/test?parsetime=true")
// 	if error != nil {
// 		fmt.Println("连接失败")
// 		return
// 	}
// 	query := `
// 		select id,title,author,price
// 		from books
// 		where price > ?
// 	`
// 	if err := db.Ping(); err != nil {
// 		fmt.Printf("数据库连接失败 %v \n", err)
// 		return
// 	}
// 	err := db.Select(&bks, query, 50)
// 	if err != nil {
// 		fmt.Printf("查询失败：%v \n", err)
// 		return
// 	}
// }

type User struct {
	gorm.Model
	Name      string
	Posts     []Post
	PostCount int //文章数量
}

type Post struct {
	gorm.Model
	Title        string //标题
	Detail       string //内容
	UserId       uint
	Comments     []Comment
	CommentState string //评论状态：0=无评论
}

type Comment struct {
	gorm.Model
	Detail string
	PostID uint
}

// 获取某一用户的所有文章和评论信息
func GetAllPostAndComment(db *gorm.DB, userId uint) ([]Post, error) {

	var posts []Post
	err := db.Model(&User{}).Where("user_id = ?", userId).Preload(clause.Associations).Find(&posts).Error
	return posts, err
}

// 获取评论最多的文章信息
func GetPostInfo(db *gorm.DB) *Post {
	var postInfo Post
	subQuery := db.Select("Count(1) as count,post_id").Group("post_id").Order("count desc").Limit(1).Table("Comments")
	db.Joins("Join (?) c on posts.id = c.post_id", subQuery).Preload("Comments").Find(&postInfo) //表名必须是结构名称+s,不然找不到
	return &postInfo
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。

func (post *Post) BeforeCreate(db *gorm.DB) error {
	// subQuery := db.Model(&Post{}).Select("count(*) as count").Where("user_id = ?", post.UserId)
	// err := db.Model(&User{Model: gorm.Model{ID: post.UserId}}).Update("post_count", subQuery).Error
	var user User
	if err := db.First(&user, post.UserId).Error; err != nil {
		return err
	}

	return db.Model(&user).Update("post_count", user.PostCount+1).Error

}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (comment *Comment) AfterDelete(db *gorm.DB) error {

	//var comments []Comment
	// if err := db.Select("count(*) as count").Where("post_id = ?", comment.PostID).Find(&comments).Error; err != nil {
	// 	return err
	// }
	var count int64
	if err := db.Model(&Comment{}).Where("post_id = ?", comment.PostID).Count(&count).Error; err != nil {
		return err
	}
	// if len(comments) == 0 {
	// 	err := db.Model(&Post{}).Where().Update("comment_state", "无评论").Error
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	if count == 0 {
		return db.Model(&Post{}).Where("id = ?", comment.PostID).Update("comment_state", "无评论").Error
	}

	return db.Debug().Model(&Post{}).Where("id = ?", comment.PostID).Update("comment_state", "有评论").Error

}
