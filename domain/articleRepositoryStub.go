package domain

import (
	"article-app/config"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleRepositoryStub struct {
	articles []Article
}

func NewArticleRepositoryStub() ArticleRepositoryStub {
	articles := generateTestArticle()
	return ArticleRepositoryStub{articles}
}

func (s ArticleRepositoryStub) FindAll(currentPage uint32) (ArticleQueryResult, int, error) {
	return s.getStubArticleQueryResult(), SUCCESS_OK, nil
}

func (s ArticleRepositoryStub) FindOne(id string) (Article, int, error) {
	return s.articles[0], SUCCESS_OK, nil
}

func (s ArticleRepositoryStub) InsertOne(article *Article) (string, error) {
	fmt.Printf("Now inserting article to mock db with Title: %v", article.Title)
	return "61db01c1a44382e1f4161b55", nil
}

func (s ArticleRepositoryStub) FindByText(searchText string, currentPage uint32) (ArticleQueryResult, int, error) {
	articleQueryResult := s.getStubArticleQueryResult()
	articleQueryResult.CurrentPage = currentPage
	articleQueryResult.Keyword = searchText
	return articleQueryResult, SUCCESS_OK, nil
}

func generateTestArticle() []Article {
	return []Article{
		{Id: primitive.NewObjectID(), Title: "Best Practices in Go Programming", Intro: "10 Best Practices tips for go programming language", Body: "1. Naming Conventions 2. Modularity", PostDate: "12.01.2022"},
		{Id: primitive.NewObjectID(), Title: "Thread Management in Go", Intro: "Learn how to use goroutines", Body: "Creating an OS Thread or switching from one to another can be costly for your programs in terms of memory and performance. Go aims to get advantages as much as possible from the cores. It has been designed with concurrency in mind from the beginning.", PostDate: "12.01.2022"},
		{Id: primitive.NewObjectID(), Title: "Restful API with GO", Intro: "Creating RESTFUL API using Golang and MongoDB", Body: "Hello, Gophers welcome back, In our previous tutorial we integrated Postgres to our Go app. In this tutorial, we are going to integrate MongoDB using the Go Mongo DB driver.", PostDate: "12.01.2022"},
	}
}

func (s ArticleRepositoryStub) getStubArticleQueryResult() ArticleQueryResult {
	ARTICLE_PER_PAGE := config.GetMaxRecordPerPage()
	count := uint32(len(s.articles))
	articleQueryResult := ArticleQueryResult{
		TotalCount:  count,
		TotalPage:   getTotalPage(count, ARTICLE_PER_PAGE),
		CurrentPage: 1,
		PerPage:     ARTICLE_PER_PAGE,
		Keyword:     "",
		Data:        s.articles,
	}
	return articleQueryResult
}
