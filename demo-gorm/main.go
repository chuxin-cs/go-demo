package main

import (
	"fmt"
	"github.com/chuxin-cs/demo-gorm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	fmt.Println("demo gorm......")
	// 由于这里没有密码所以省略password=号后面没有值
	dsn := "root:123456@tcp(localhost:3306)/spring-boot-demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	// 创建项目
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&models.User{})
	// 插入一条数据
	db.Create(&Product{Code: "D42", Price: 100})

	// 创建
	user := models.User{
		Name:     "chuxin",
		Email:    "chuxin@qq.com",
		Password: "123456",
	}
	db.Create(&user)

	// 查询
	var userList models.User
	db.First(&userList)
	log.Printf("First user: %#v\n", user)

	db.Find(&userList, "email = ?", "chuxin@qq.com")
	log.Printf("Find user: %#v\n", user)

	// 获取全部记录
	var users []models.User
	db.Find(&users)
	// 循环取出所有的users遍历打印ID
	for _, u := range users {
		log.Printf("User ID: %d\n", u.ID)
	}

}
