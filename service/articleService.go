package service

import (
	"article-app/domain"
)

type ArticleService interface {
	GetArticles(currentPage uint32) (domain.ArticleQueryResult, int, error)
	PublishArticle(article *domain.Article) (string, error)
	SearchArticle(searchText string, currentPage uint32) (domain.ArticleQueryResult, error)
	GetArticleById(id string) (domain.Article, int, error)
}

/*there may be more implementation in future*/
type DefaultArticleService struct {
	repo domain.ArticleRepository
}

func NewArticleService(repository domain.ArticleRepository) DefaultArticleService {
	return DefaultArticleService{repository}
}

func (s DefaultArticleService) GetArticles(currentPage uint32) (domain.ArticleQueryResult, int, error) {
	return s.repo.FindAll(currentPage)
}

func (s DefaultArticleService) GetArticleById(id string) (domain.Article, int, error) {
	return s.repo.FindOne(id)
}

func (s DefaultArticleService) PublishArticle(article *domain.Article) (string, error) {
	return s.repo.InsertOne(article)
}

func (s DefaultArticleService) SearchArticle(searchText string, currrentPage uint32) (domain.ArticleQueryResult, error) {
	return s.repo.FindByText(searchText, currrentPage)
}
