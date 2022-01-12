package app

import (
	"article-app/domain"
	"article-app/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//wiring
	//ah := ArticleHandlers{service.NewArticleService(domain.NewArticleRepositoryStub())}
	ah := ArticleHandlers{service.NewArticleService(domain.NewArticleRepositoryDb())}
	handleRequests(&ah)
}

//publishArticle
//getArticles
//searchArticle
func handleRequests(ah *ArticleHandlers) {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/articles/page/{page}", ah.getArticlesForPage).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/articles/id/{id}", ah.getArticleById).Methods(http.MethodGet) //DONE
	router.HandleFunc("/api/v1/articles", ah.getArticles).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/articles", ah.publishArticle).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/articles/search/{searchText}/page/{page}", ah.searchArticleForPage).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/articles/search/{searchText}", ah.searchArticle).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", router))
	//log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
}
