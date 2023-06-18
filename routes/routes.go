package routes

import (
	"go-api-rest/controllers"
	"go-api-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	request := mux.NewRouter()
	request.Use(middleware.ContentType)
	request.HandleFunc("/", controllers.Home)
	request.HandleFunc("/", controllers.Home)
	request.HandleFunc("/personalities", controllers.AllPersonalities).Methods("GET")
	request.HandleFunc("/personalities", controllers.CreatePerson).Methods("POST")
	request.HandleFunc("/personalities/{id}", controllers.ReturnPersonality).Methods("GET")
	request.HandleFunc("/personalities/{id}", controllers.EditPerson).Methods("PUT")
	request.HandleFunc("/personalities/{id}", controllers.DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3333", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(request)))
}
