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

	router.HandleFunc("/players", GetAllPlayers).Methods("GET")
	router.HandleFunc("/players/{id}", GetPlayer).Methods("GET")
	router.HandleFunc("/players", CreatePlayer).Methods("POST")
	router.HandleFunc("/players/{id}", UpdatePlayer).Methods("PUT")
	router.HandleFunc("/players/{id}", DeletePlayer).Methods("DELETE")

	return router
}
