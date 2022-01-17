package domain

import (
	"article-app/config"
	"context"
	"errors"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ARTICLES_COLLECTION_NAME string = "articles"
const TIME_FORMAT_POST_DATE string = "2006-01-02 15:04:05"

type ArticleRepositoryDb struct {
	client *mongo.Database
}

func NewArticleRepositoryDb() ArticleRepositoryDb {
	return ArticleRepositoryDb{config.GetMongoDBConnection()}
}

func (d ArticleRepositoryDb) FindAll(currentPage uint32) (ArticleQueryResult, int, error) {
	return d.findAllByText("", currentPage)
}

func (d ArticleRepositoryDb) FindByText(searchText string, currentPage uint32) (ArticleQueryResult, int, error) {
	return d.findAllByText(searchText, currentPage)
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

func (d ArticleRepositoryDb) InsertOne(article *Article) (string, error) {
	articlesCollection := GetArticlesCollection()
	article.PostDate = time.Now().Format(TIME_FORMAT_POST_DATE)
	res, err := articlesCollection.InsertOne(context.TODO(), article)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (d ArticleRepositoryDb) findAllByText(searchText string, currentPage uint32) (ArticleQueryResult, int, error) {
	ARTICLE_PER_PAGE := config.GetMaxRecordPerPage()
	articlesCollection := GetArticlesCollection()

	var filter interface{}
	if len(searchText) > 0 {
		filter = bson.M{
			"$text": bson.M{
				"$search": searchText,
			},
		}
	} else {
		filter = bson.D{}
	}
	count, err := articlesCollection.CountDocuments(context.TODO(), filter, nil)
	if err != nil {
		return ArticleQueryResult{}, ERROR_COUNT_QUERY_FAILED, err
	}

	opts := getFindOptions(currentPage, config.GetMaxRecordPerPage())

	cursor, err := articlesCollection.Find(context.TODO(), filter, opts)
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

	articleQueryResult := ArticleQueryResult{
		TotalCount:  totalCount,
		TotalPage:   getTotalPage(totalCount, config.GetMaxRecordPerPage()),
		CurrentPage: currentPage,
		PerPage:     ARTICLE_PER_PAGE,
		Keyword:     searchText,
		Data:        articles,
	}

	return articleQueryResult, SUCCESS_OK, err
}

func getArticlesFromResultReturn(results []bson.M) []Article {
	var articles []Article
	var tempArticle Article
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, &tempArticle)
		articles = append(articles, tempArticle)
	}
	if articles == nil {
		articles = []Article{}
	}
	return articles
}

func getFindOptions(currentPage uint32, articlePerPage uint32) *options.FindOptions {
	opts := options.Find().SetSkip(int64(currentPage-1) * int64(articlePerPage))
	opts.SetLimit(int64(articlePerPage))
	opts.SetSort(bson.D{{Key: "_id", Value: 1}})
	return opts
}

func getTotalPage(totalCount uint32, articlePerPage uint32) uint32 {
	var totalPage uint32
	if totalCount == 0 {
		totalPage = totalCount
	} else if math.Mod(float64(totalCount), float64(articlePerPage)) == 0 {
		totalPage = (totalCount / articlePerPage)
	} else {
		totalPage = (totalCount / articlePerPage) + 1
	}
	return totalPage
}

func GetArticlesCollection() *mongo.Collection {
	return config.GetMongoDBConnection().Collection(ARTICLES_COLLECTION_NAME)
}
