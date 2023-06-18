package controllers

import (
	"encoding/json"
	"fmt"
	"go-api-rest/database"
	"go-api-rest/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprint(writer, "Home Page")
}

func AllPersonalities(writer http.ResponseWriter, request *http.Request) {
	var person []models.Personality
	database.DB.Find(&person)
	json.NewEncoder(writer).Encode(person)
}

func ReturnPersonality(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	var person models.Personality

	database.DB.First(&person, id)
	json.NewEncoder(writer).Encode(person)
}

func CreatePerson(writer http.ResponseWriter, request *http.Request) {
	var newPerson models.Personality
	json.NewDecoder(request.Body).Decode(&newPerson)
	database.DB.Create(&newPerson)
	json.NewEncoder(writer).Encode(newPerson)
}

func DeletePerson(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	var person models.Personality

	database.DB.Delete(&person, id)
	json.NewEncoder(writer).Encode(person)
}

func EditPerson(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	var person models.Personality

	database.DB.First(&person, id)
	json.NewDecoder(request.Body).Decode(&person)
	database.DB.Save(&person)
	json.NewEncoder(writer).Encode(person)
}