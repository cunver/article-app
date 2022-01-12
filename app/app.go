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

func Start() {
	//wiring
	var ah ArticleHandlers
	if config.UseStubDB() {
		ah = ArticleHandlers{service.NewArticleService(domain.NewArticleRepositoryStub())}
	} else {
		ah = ArticleHandlers{service.NewArticleService(domain.NewArticleRepositoryDb())}
	}

	handleRequests(&ah)
}

func handleRequests(ah *ArticleHandlers) {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/articles", ah.getArticles).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/articles", ah.publishArticle).Methods(http.MethodPost) //publishArticle
	router.HandleFunc("/api/v1/articles/search", ah.searchArticleForPage).Methods(http.MethodGet)

	router.HandleFunc("/api/v1/articles/page/{page}", ah.getArticles).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/articles/id/{id}", ah.getArticleById).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe((":" + strconv.Itoa(config.GetServerConfig().Port)), router))
	//log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
}
