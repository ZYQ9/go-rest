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

/*
Function that gets all quests from the database
*/

func GetAllQuests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var quests []models.Quest
	models.DB.Find(&quests)

	json.NewEncoder(w).Encode(quests)
}

/*
Get a single quest from the database

To do this we:
1. Create an empty quest object
2. Query the database to find the corresponding record
3. Fill the quest object with the properties of the record
*/
func GetQuest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var quest models.Quest

	if err := models.DB.Where("id = ?", id).First(&quest).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
		return
	}

	json.NewEncoder(w).Encode(quest)
}

// Create a quest
var validate *validator.Validate

type QuestInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Reward      int    `json:"reward" validate:"required"`
}

func CreateQuest(w http.ResponseWriter, r *http.Request) {
	var input QuestInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	quest := &models.Quest{
		Title:       input.Title,
		Description: input.Description,
		Reward:      input.Reward,
	}

	models.DB.Create(quest)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(quest)

}

// Updating a quest
func UpdateQuest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var quest models.Quest

	if err := models.DB.Where("id = ?", id).First(&quest).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
		return
	}

	var input QuestInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	quest.Title = input.Title
	quest.Description = input.Description
	quest.Reward = input.Reward

	models.DB.Save(&quest)

	json.NewEncoder(w).Encode(quest)
}

// Deleting a quest
func DeleteQuest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var quest models.Quest

	if err := models.DB.Where("id = ?", id).First(&quest).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
		return
	}

	models.DB.Delete(&quest)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(quest)
}
