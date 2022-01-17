package app

import (
	"article-app/config"
	"article-app/domain"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetArticles(t *testing.T) {
	w := executeRequestOnRouter(getGetArticlesRequest())
	logFailIfNotSuccess(w, t)
	articleQueryResult := getArticleQueryResult(w)
	logFailIfInvalidQueryResult(articleQueryResult, t)
}

func TestPublishArticle(t *testing.T) {
	req := getPublishArticleRequest()
	w := executeRequestOnRouter(req)
	logFailIfNotSuccess(w, t)
	logFailIfInvalidPublishResponse(w, t)
}

func TestSearchArticleReturnSomeArticles(t *testing.T) {
	w := executeRequestOnRouter(getSearchArticleRequest("%20Go%20"))
	logFailIfNotSuccess(w, t)
	articleQueryResult := getArticleQueryResult(w)
	logFailIfInvalidQueryResult(articleQueryResult, t)
}

func TestGetArticlesForPage(t *testing.T) {
	currentPage := 1
	w := executeRequestOnRouter(getGetArticlesForPageRequest(currentPage))
	logFailIfNotSuccess(w, t)
	articleQueryResult := getArticleQueryResult(w)
	logFailIfInvalidQueryResult(articleQueryResult, t)
	assert.Equal(t, articleQueryResult.CurrentPage, uint32(currentPage))
}

func TestGetArticlesById(t *testing.T) {
	objectId := "61de26895b18304a46524721"
	w := executeRequestOnRouter(getGetArticlesByIdRequest(objectId))
	logFailIfNotSuccess(w, t)
	if w.Code == http.StatusOK {
		article := getArticleFromResponse(w)
		assert.NotEmpty(t, article, "GetArticleById returned success and expected a real article but found empty")
	}
}

func createTestRouterWithRoutes() *mux.Router {
	config.ReadConfig()
	return CreateRouterWithRoutes()
}

func getArticleQueryResult(w *httptest.ResponseRecorder) *domain.ArticleQueryResult {
	var result *domain.ArticleQueryResult
	json.Unmarshal(w.Body.Bytes(), &result)
	return result
}

func getArticleFromResponse(w *httptest.ResponseRecorder) *domain.Article {
	var result *domain.Article
	json.Unmarshal(w.Body.Bytes(), &result)
	return result
}

func logFailIfNotSuccess(w *httptest.ResponseRecorder, t *testing.T) {
	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}
}

func logFailIfInvalidQueryResult(articleQueryResult *domain.ArticleQueryResult, t *testing.T) {
	if articleQueryResult.TotalCount == 0 {
		t.Error("Article query result total count is expected > 0, got 0")
	}
}

func logFailIfInvalidPublishResponse(w *httptest.ResponseRecorder, t *testing.T) {
	var pubRes *PublishResponse
	json.Unmarshal(w.Body.Bytes(), &pubRes)
	assert.NotEmpty(t, pubRes, "Publish article response returned empty!")
	assert.Equal(t, len(pubRes.Id), 24, "Publish article response invalid. Expected 24 character length object id, got %v!", pubRes.Id)
}

func getGetArticlesRequest() *http.Request {
	return httptest.NewRequest("GET", PATH_ARTICLES, nil)
}

func getPublishArticleRequest() *http.Request {
	article, err := getArticleToPublish()
	if err != nil {
		log.Fatalf("Generate article to publish failed. Error:%v", err)
	}
	return httptest.NewRequest("POST", PATH_ARTICLES, bytes.NewReader(article))
}

func getSearchArticleRequest(searchText string) *http.Request {
	return httptest.NewRequest("GET", PATH_ARTICLES_SEARCH+"?searchText="+searchText+"&page=1", nil)
}

func getGetArticlesForPageRequest(pageId int) *http.Request {
	return httptest.NewRequest("GET", PATH_ARTICLES+"/page/"+strconv.Itoa(pageId), nil)
}

func getGetArticlesByIdRequest(objectId string) *http.Request {
	return httptest.NewRequest("GET", PATH_ARTICLES+"/"+objectId, nil)
}

func executeRequestOnRouter(req *http.Request) *httptest.ResponseRecorder {
	router := createTestRouterWithRoutes()
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

// TODO : FAKER will be used to get random articles
func getArticleToPublish() ([]byte, error) {
	var article domain.Article = domain.Article{
		Title:    "Best Practices in Go Programming at time : " + time.Now().String(),
		Intro:    "10 Best Practices tips for go programming language",
		Body:     "1. Naming Conventions 2. Modularity",
		PostDate: time.Now().Format(domain.TIME_FORMAT_POST_DATE),
	}
	articleStr, err := json.Marshal(article)
	if err == nil {
		return articleStr, nil
	}
	return nil, err
}
