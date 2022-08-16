package model

import "github.com/go-programming-tour-book/blog-service/pkg/app"

type TagSwagger struct {
	List []*Tag
	Pager *app.Pager
}
