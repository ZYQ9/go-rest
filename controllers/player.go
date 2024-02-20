package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ZYQ9/go-rest/models"
	"github.com/ZYQ9/go-rest/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

// Function that gets all players from the database

func GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var players []models.Player
	models.DB.Find(&players)

	json.NewEncoder(w).Encode(players)
}

// Get a single player from the database
func GetPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the content type to JSON

	id := mux.Vars(r)["id"] // Filter on the id parameter from the request
	var player models.Player

	// Error handling for the database query
	if err := models.DB.Where("id = ?", id).First(&player).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Player not found")
		return
	}

	json.NewEncoder(w).Encode(player)
}

/* -------------------------------
Create a player
----------------------------------*/

// Sets the required fields for the player
type PlayerInput struct {
	Class string `json:"class" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Level int    `json:"level" validate:"required"`
}

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var input PlayerInput

	// Read the request body and parse
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)
	validate = validator.New()
	err := validate.Struct(input)

	// Error handling for the request body
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Create a new player
	player := models.Player{
		Class: input.Class,
		Name:  input.Name,
		Level: input.Level,
	}

	models.DB.Create(player)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(player)
}

/* -------------------------------
Update a player
----------------------------------*/

func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var player models.Player

	if err := models.DB.Where("id = ?", id).First(&player).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Player not found")
		return
	}

	var input PlayerInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	player.Class = input.Class
	player.Name = input.Name
	player.Level = input.Level

	models.DB.Save(player)

	json.NewEncoder(w).Encode(player)
}

/* -------------------------------
Delete a player
----------------------------------*/

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var player models.Player

	if err := models.DB.Where("id = ?", id).First(&player).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Player not found")
		return
	}

	models.DB.Delete(&player)

	json.NewEncoder(w).Encode(player)
}
