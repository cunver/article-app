package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title"`
	Intro    string             `json:"intro"`
	Body     string             `json:"body"`
	PostDate string             `json:"postDate"`
}

type ArticleQueryResult struct {
	TotalCount  uint32    `json:"totalCount"`
	TotalPage   uint32    `json:"totalPage"`
	CurrentPage uint32    `json:"currentPage"`
	PerPage     uint32    `json:"perPage"`
	Keyword     string    `json:"keyword"`
	Data        []Article `json:"data"`
}

type ArticleRepository interface {
	FindAll(currentPage uint32) (ArticleQueryResult, int, error)
	FindOne(id string) (Article, int, error)
	InsertOne(article *Article) (string, error)
	FindByText(searchText string, currentPage uint32) (ArticleQueryResult, int, error)
}

const SUCCESS_OK int = 0
const ERROR_RECORD_NOT_FOUND int = 1
const ERROR_OBJECT_ID_NOT_VALID int = 2
const ERROR_COUNT_QUERY_FAILED int = 3
const ERROR_QUERY_FAILED int = 4
const ERROR_QUERY_RESULT_MAPPING_FAILED = 5

func init() {

}
