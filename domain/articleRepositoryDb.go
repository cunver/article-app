package domain

import (
	"article-app/config"
	"context"
	"errors"
	"fmt"
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ArticleRepositoryDb struct {
	client *mongo.Database
}

func NewArticleRepositoryDb() ArticleRepositoryDb {
	return ArticleRepositoryDb{config.GetMongoDBConnection()}
}

func (d ArticleRepositoryDb) FindAll(currentPage uint32) (ArticleQueryResult, int, error) {
	articlesCollection := GetArticlesCollection()

	count, err := articlesCollection.CountDocuments(context.TODO(), bson.D{}, nil)
	if err != nil {
		return ArticleQueryResult{}, ERROR_COUNT_QUERY_FAILED, err
	}

	opts := options.Find().SetSkip(int64(currentPage-1) * int64(ARTICLE_PER_PAGE))
	opts.SetLimit(int64(ARTICLE_PER_PAGE))
	opts.SetSort(bson.D{{Key: "_id", Value: 1}})

	cursor, err := articlesCollection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return ArticleQueryResult{}, ERROR_QUERY_FAILED, err
	}

	var articles []Article
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return ArticleQueryResult{}, ERROR_QUERY_RESULT_MAPPING_FAILED, err
	}
	articles = getArticlesFromResultReturn(results)
	totalCount := uint32(count)
	var totalPage uint32
	if totalCount == 0 {
		totalPage = totalCount
	} else if math.Mod(float64(totalCount), float64(ARTICLE_PER_PAGE)) == 0 {
		totalPage = (totalCount / ARTICLE_PER_PAGE)
	} else {
		totalPage = (totalCount / ARTICLE_PER_PAGE) + 1
	}
	articleQueryResult := ArticleQueryResult{
		TotalCount:  totalCount,
		TotalPage:   totalPage,
		CurrentPage: currentPage,
		PerPage:     ARTICLE_PER_PAGE,
		Keyword:     "",
		Data:        articles,
	}

	return articleQueryResult, SUCCESS_OK, err
}

func (d ArticleRepositoryDb) FindOne(id string) (Article, int, error) {
	articlesCollection := GetArticlesCollection()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Article{}, ERROR_OBJECT_ID_NOT_VALID, err
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	singleResult := articlesCollection.FindOne(context.TODO(), filter)
	var article Article
	err = singleResult.Decode(&article) // gets ErrNoDocuments if no result
	if err != nil {
		return Article{}, ERROR_RECORD_NOT_FOUND, errors.New("No article found with id:" + id)
	}
	return article, SUCCESS_OK, nil
}

func getArticlesFromResultReturn(results []bson.M) []Article {
	var articles []Article
	var tempArticle Article
	fmt.Printf("Result size %v", len(results))
	for _, result := range results {
		fmt.Println(result)
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, &tempArticle)
		articles = append(articles, tempArticle)
	}
	return articles
}

func (d ArticleRepositoryDb) InsertOne(article *Article) (string, error) {
	articlesCollection := GetArticlesCollection()
	res, err := articlesCollection.InsertOne(context.TODO(), article)
	if err != nil {
		return "", err
		//fmt.Println("Article inserted to mongodb with Id: "+res.InsertedID)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (d ArticleRepositoryDb) FindByText(searchText string, currentPage uint32) (ArticleQueryResult, error) {

	fmt.Printf("Searching mongodb for searchText:" + searchText + "-")
	articlesCollection := GetArticlesCollection()

	filter := bson.M{
		"$text": bson.M{
			"$search": searchText,
		},
	}

	count, err := articlesCollection.CountDocuments(context.TODO(), filter, nil)
	if err != nil {
		fmt.Printf("Search article count error : %v", err.Error())
	}

	opts := options.Find().SetSkip(int64(currentPage-1) * int64(ARTICLE_PER_PAGE))
	opts.SetLimit(int64(ARTICLE_PER_PAGE))
	opts.SetSort(bson.D{{Key: "_id", Value: 1}})
	cursor, err := articlesCollection.Find(context.TODO(), filter, opts, nil)
	if err != nil {
		fmt.Printf("Search article find error : %v", err.Error())
	}

	// Get a list of all returned documents and print them out.
	// See the mongo.Cursor documentation for more examples of using cursors.
	var articles []Article
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Printf("Search article error : %v", err.Error())
	} else {
		articles = getArticlesFromResultReturn(results)

	}

	totalCount := uint32(count)
	var totalPage uint32
	if totalCount == 0 {
		totalPage = totalCount
	} else if math.Mod(float64(totalCount), float64(ARTICLE_PER_PAGE)) == 0 {
		totalPage = (totalCount / ARTICLE_PER_PAGE)
	} else {
		totalPage = (totalCount / ARTICLE_PER_PAGE) + 1
	}

	articleQueryResult := ArticleQueryResult{
		TotalCount:  totalCount,
		TotalPage:   totalPage,
		CurrentPage: currentPage,
		PerPage:     ARTICLE_PER_PAGE,
		Keyword:     searchText,
		Data:        articles,
	}

	return articleQueryResult, err
}

func GetArticlesCollection() *mongo.Collection {
	return config.GetMongoDBConnection().Collection("articles")
}
