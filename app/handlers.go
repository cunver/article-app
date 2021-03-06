package app

import (
	"article-app/domain"
	"article-app/service"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
)

var Articles []domain.Article

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
	vars := mux.Vars(r)
	currentPage := getCurrentPageParameter(vars)
	articles, errCode, err := ah.service.GetArticles(currentPage)
	if err != nil {
		message := getErrorMessageFromErrorCode(errCode)
		handleErrorResponse(w, http.StatusInternalServerError, "Could not get articles for page. Error:"+message)
		return
	}
	handleSuccessResponse(w, articles)
}

func (ah *ArticleHandlers) publishArticle(w http.ResponseWriter, r *http.Request) {
	setConentTypeAsJson(w)
	var articleInput domain.Article
	json.NewDecoder(r.Body).Decode(&articleInput)
	err := validateArticleBeforePublish(articleInput)
	if err != nil {
		handleErrorResponse(w, http.StatusBadRequest, "Article input is not valid. Error:"+err.Error())
		return
	}
	id, err := ah.service.PublishArticle(&articleInput)
	if err != nil {
		handleErrorResponse(w, http.StatusInternalServerError, "Could not publish the article. Error:"+err.Error())
		return
	}
	handleSuccessResponse(w, PublishResponse{Id: id})
}

func (ah *ArticleHandlers) searchArticleForPage(w http.ResponseWriter, r *http.Request) {
	setConentTypeAsJson(w)
	parsedURL, err := url.Parse(r.URL.String())
	if err != nil {
		handleErrorResponse(w, http.StatusInternalServerError, "Could not search parse URL. Error:"+err.Error())
		return
	}
	searchTextParam := parsedURL.Query().Get("searchText")
	currentPageParam := r.URL.Query().Get("page")
	page := getCurrentPageFromParamStr(currentPageParam)
	articleQueryResult, errCode, err := ah.service.SearchArticle(searchTextParam, page)
	if err != nil {
		message := getErrorMessageFromErrorCode(errCode)
		handleErrorResponse(w, http.StatusInternalServerError, "Could not search articles. Error:"+message)
		return
	}
	handleSuccessResponse(w, articleQueryResult)
}

func (ah *ArticleHandlers) getArticleById(w http.ResponseWriter, r *http.Request) {
	setConentTypeAsJson(w)
	vars := mux.Vars(r)
	article, errCode, err := ah.service.GetArticleById(vars["id"])
	if err == nil {
		handleSuccessResponse(w, article)
	} else {
		handleErrorResponse(w, getStatusCodeFromErrorCode(errCode), err.Error())
	}
}

func validateArticleBeforePublish(a domain.Article) error {
	if len(a.Id) == 0 {
		return errors.New("id can not be empty")
	} else if len(a.Title) == 0 {
		return errors.New("title can not be empty")
	} else if len(a.Intro) == 0 {
		return errors.New("intro can not be empty")
	} else if len(a.Body) == 0 {
		return errors.New("body can not be empty")
	}
	return nil
}

func getCurrentPageParameter(varsMap map[string]string) uint32 {
	currentPageStr := varsMap["page"]
	if currentPageStr == "" {
		return 1 // if no specific page id given in parameters use 1 to get the first page
	} else {
		currentPageInt, err := strconv.ParseUint(currentPageStr, 10, 64)
		if err == nil {
			return uint32(currentPageInt)
		}
	}
	return 1
}

func getCurrentPageFromParamStr(currentPageStr string) uint32 {
	var page uint32 = 1
	if len(currentPageStr) > 0 {
		currentPageInt, err := strconv.ParseUint(currentPageStr, 10, 64)
		if err == nil {
			page = uint32(currentPageInt)
		} else {
			log.Printf("getCurrentPage parse error for : %v, error:%v", currentPageStr, err)
		}
	}
	return page
}

func handleSuccessResponse(w http.ResponseWriter, v interface{}) {
	json.NewEncoder(w).Encode(v)
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
		message = "Unexpected error."
	}
	return message
}

func getStatusCodeFromErrorCode(errorCode int) int {
	var httpStatusCode int = http.StatusInternalServerError
	if errorCode == domain.ERROR_OBJECT_ID_NOT_VALID {
		httpStatusCode = http.StatusNotAcceptable
	} else if errorCode == domain.ERROR_RECORD_NOT_FOUND {
		httpStatusCode = http.StatusNotFound
	}
	return httpStatusCode
}
