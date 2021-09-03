package controllers

import "os"

const (
	DefaultPageNumber = 0
	DefaultPageSize   = 20

	ArticleQueryParamValue = "article"
	CommentQueryParamValue = "comment"

	ParentTypeQueryParam = "parent-type"

	ParentIDPathParam  = "parent-id"
	ArticleIDPathParam = "id"

	DatabaseName = "DATABASE_NAME"
	UserName     = "USER_NAME"
	Password     = "PASSWORD"
	Host         = "HOST"
	Port         = "PORT"
)

var (
	DBName     = os.Getenv(DatabaseName)
	DBUserName = os.Getenv(UserName)
	DBPassword = os.Getenv(Password)
	DBHost     = os.Getenv(Host)
	DBPort     = os.Getenv(Port)
)
