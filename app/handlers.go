package app

import (
	"article-app/domain"
	"article-app/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Articles []domain.Article

/*
/publishArticle
/getArticles
/searchArticle
*/

type ErrorResponse struct {
	Message string `json:"error"`
}

type PublishResponse struct {
	Id string `json:"id"`
}
type ArticleHandlers struct {
	service service.ArticleService
}

func (ah *ArticleHandlers) getArticles(w http.ResponseWriter, r *http.Request) {
	setConentTypeAsJson(w)
	articles, errCode, err := ah.service.GetArticles(1)
	if err != nil {
		message := getErrorMessageFromErrorCode(errCode)
		handleErrorResponse(w, http.StatusInternalServerError, "Could not get articles. Error:"+message)
		return
	}
	json.NewEncoder(w).Encode(articles)
}

func getErrorMessageFromErrorCode(errorCode int) string {
	var message string
	switch errorCode {
	case domain.ERROR_COUNT_QUERY_FAILED:
		message = "Count query error."
	case domain.ERROR_QUERY_FAILED:
		message = "Query error."
	case domain.ERROR_QUERY_RESULT_MAPPING_FAILED:
		message = "Query result mapping error."
	default:
		message = "Unexpected error"
	}
	return message
}

func (ah *ArticleHandlers) getArticlesForPage(w http.ResponseWriter, r *http.Request) {
	setConentTypeAsJson(w)
	vars := mux.Vars(r)
	currentPage := getCurrentPage(vars)
	articles, errCode, err := ah.service.GetArticles(currentPage)
	if err != nil {
		message := getErrorMessageFromErrorCode(errCode)
		handleErrorResponse(w, http.StatusInternalServerError, "Could not get articles for page. Error:"+message)
		return
	}
	json.NewEncoder(w).Encode(articles)
}

func (ah *ArticleHandlers) publishArticle(w http.ResponseWriter, r *http.Request) {
	setConentTypeAsJson(w)
	var articleInput domain.Article
	json.NewDecoder(r.Body).Decode(&articleInput)

	id, err := ah.service.PublishArticle(&articleInput)
	if err != nil {
		handleErrorResponse(w, http.StatusInternalServerError, "Article could not published. Error"+err.Error())
		return
	}
	publishResponse := PublishResponse{
		Id: id,
	}
	json.NewEncoder(w).Encode(publishResponse)
	log.Println(err)
}

func (ah *ArticleHandlers) searchArticle(w http.ResponseWriter, r *http.Request) {
	ah.searchArticleForTextAndCurrentPage(w, r, 1)
}

func (ah *ArticleHandlers) searchArticleForPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currentPage := getCurrentPage(vars)
	ah.searchArticleForTextAndCurrentPage(w, r, currentPage)
}

func getCurrentPage(varsMap map[string]string) uint32 {
	currentPageStr := varsMap["page"]
	if currentPageStr == "" {
		return 1
	} else {
		currentPageInt, err := strconv.ParseUint(currentPageStr, 10, 64)
		if err == nil {
			return uint32(currentPageInt)
		}
	}
	return 1
}

func (ah *ArticleHandlers) searchArticleForTextAndCurrentPage(w http.ResponseWriter, r *http.Request, currentPage uint32) {
	setConentTypeAsJson(w)
	vars := mux.Vars(r)
	searchText := vars["searchText"]
	articleQueryResult, _ := ah.service.SearchArticle(searchText, currentPage)
	json.NewEncoder(w).Encode(articleQueryResult)
}

func (ah *ArticleHandlers) getArticleById(w http.ResponseWriter, r *http.Request) {
	setConentTypeAsJson(w)
	vars := mux.Vars(r)
	article, errCode, err := ah.service.GetArticleById(vars["id"])
	if err == nil {
		json.NewEncoder(w).Encode(article)
	} else {
		var httpStatusCode int = http.StatusInternalServerError
		if errCode == domain.ERROR_OBJECT_ID_NOT_VALID {
			httpStatusCode = http.StatusNotAcceptable
		} else if errCode == domain.ERROR_RECORD_NOT_FOUND {
			httpStatusCode = http.StatusNotFound
		}
		handleErrorResponse(w, httpStatusCode, err.Error())

	}
}

func handleErrorResponse(w http.ResponseWriter, statusCode int, err string) {
	w.WriteHeader(statusCode)
	errResponse := ErrorResponse{
		Message: err,
	}
	json.NewEncoder(w).Encode(errResponse)
}

func setConentTypeAsJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

/*
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	switch status {
	case http.StatusBadRequest:
		fmt.Println("Custom bad request error")
	case http.StatusNotFound:
		fmt.Println("Custom Not found")
	default:
		fmt.Println("Custom Default error")
	}
}
*/
