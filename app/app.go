package app

import (
	"article-app/config"
	"article-app/domain"
	"article-app/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const PATH_ARTICLES string = "/api/v1/articles"
const PATH_ARTICLES_SEARCH string = "/api/v1/articles/search"
const PATH_ARTICLES_PAGE string = "/api/v1/articles/page/{page}"
const PATH_ARTICLES_ID string = "/api/v1/articles/{id}"

func Start() {
	config.ReadConfig()
	router := CreateRouterWithRoutes()
	serveRouter(router)
}

func CreateRouterWithRoutes() *mux.Router {
	router := mux.NewRouter()
	setHandlers(router, getArticleHandler())
	return router
}

func setHandlers(router *mux.Router, ah *ArticleHandlers) {
	router.HandleFunc(PATH_ARTICLES, ah.getArticles).Methods(http.MethodGet)                 //getArticles
	router.HandleFunc(PATH_ARTICLES, ah.publishArticle).Methods(http.MethodPost)             //publishArticle
	router.HandleFunc(PATH_ARTICLES_SEARCH, ah.searchArticleForPage).Methods(http.MethodGet) //searchArticles
	router.HandleFunc(PATH_ARTICLES_PAGE, ah.getArticles).Methods(http.MethodGet)            //getArticlesForPage
	router.HandleFunc(PATH_ARTICLES_ID, ah.getArticleById).Methods(http.MethodGet)           //getArticleById
}

func getArticleHandler() *ArticleHandlers {
	ah := ArticleHandlers{service.NewArticleService(getRepository())}
	return &ah
}

func getRepository() domain.ArticleRepository {
	var repository domain.ArticleRepository
	if config.UseStubDB() {
		repository = domain.NewArticleRepositoryStub()
	} else {
		repository = domain.NewArticleRepositoryDb()
	}
	return repository
}
func serveRouter(router *mux.Router) {
	log.Fatal(http.ListenAndServe((":" + strconv.Itoa(config.GetServerConfig().Port)), router))
	//log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
}
