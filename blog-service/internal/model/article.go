package model

import "github.com/go-programming-tour-book/blog-service/pkg/app"

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}
