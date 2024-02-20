package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {

	router := mux.NewRouter()

	router.HandleFunc("/quests", GetAllQuests).Methods("GET")
	router.HandleFunc("/quests/{id}", GetQuest).Methods("GET")
	router.HandleFunc("/quests", CreateQuest).Methods("POST")
	router.HandleFunc("/quests/{id}", UpdateQuest).Methods("PUT")
	router.HandleFunc("/quests/{id}", DeleteQuest).Methods("DELETE")

	return router
}
