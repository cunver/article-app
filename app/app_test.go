package app

import (
	"article-app/config"
	"article-app/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetArticles(t *testing.T) {
	w := executeRequestOnRouter(getGetArticlesRequest())
	fmt.Printf("W body: %v", w.Body.String())
	logFailIfNotSuccess(w, t)
	articleQueryResult := getArticleQueryResult(w)
	logFailIfInvalidQueryResult(articleQueryResult, t)
}

func TestPublishArticle(t *testing.T) {
	w := executeRequestOnRouter(getPublishArticleRequest())
	logFailIfNotSuccess(w, t)
	logFailIfInvalidPublishResponse(w, t)
}

func TestSearchArticleReturnSomeArticles(t *testing.T) {
	w := executeRequestOnRouter(getSearchArticleRequest("URL%20encode%2Fdecode%2BparseQuery"))
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
	objectId := "61db01d7a44382e1f4161b63"
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
	return httptest.NewRequest("POST", PATH_ARTICLES, nil)
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
	//fmt.Printf("W body: %v", w.Body.String())
	return w
}
