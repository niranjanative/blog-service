package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/niranjanative/blog-service/controllers"
	"github.com/niranjanative/blog-service/models/articles"
	"github.com/niranjanative/blog-service/models/comments"
)

var DB *gorm.DB

func ConnectDatabase() {

	if controllers.DBName == "" {
		panic("No DatabaseName provided to connect to database!")
	}
	if controllers.DBUserName == "" {
		panic("No UserName provided to connect to database!")
	}
	if controllers.DBPassword == "" {
		panic("No Password provided to connect to database!")
	}
	if controllers.DBHost == "" {
		panic("No Host provided to connect to database!")
	}
	if controllers.DBPort == "" {
		panic("No Port provided to connect to database!")
	}

	// dataSourceName for db
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", controllers.DBUserName,
		controllers.DBPassword, controllers.DBHost, controllers.DBPort, controllers.DBName)
	database, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to database!" + err.Error())
	}

	database.AutoMigrate(&articles.Article{}, &comments.Comment{})

	DB = database
}
