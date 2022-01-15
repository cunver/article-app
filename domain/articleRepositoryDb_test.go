package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

var tests = []struct {
	totalCount        uint32
	articlePerPage    uint32
	expectedTotalPage uint32
}{
	{12, 4, 3},
	{12, 12, 1},
	{12, 11, 2},
	{12, 14, 1},
	{12, 1, 12},
	{0, 2, 0},
}

func TestGetTotalPage(t *testing.T) {
	//when
	for _, e := range tests {
		res := getTotalPage(e.totalCount, e.articlePerPage)
		assert.Equal(t, res, e.expectedTotalPage)
	}
}

func TestGetArticlesFromResultReturn(t *testing.T) {
	var bsonObj bson.M = bson.M{"ObjectID": "61db01c1a44382e1f4161b55", "articleId": 1, "body": "Go Programming"}
	var bsonObj2 bson.M = bson.M{"ObjectID": "61db01d7a44382e1f4161b63", "articleId": 2, "body": "Thread Management in Go"}
	var bsonMap []bson.M = []bson.M{bsonObj, bsonObj2}

	articles := getArticlesFromResultReturn(bsonMap)
	assert.NotEmpty(t, articles, "Cannot converted into articles from bson string")
	assert.Equal(t, len(articles), 2, "Expected 2 articles but got %v", len(articles))
}
