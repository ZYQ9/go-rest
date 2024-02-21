package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ZYQ9/go-rest/models"
	"github.com/ZYQ9/go-rest/utils"
	"github.com/gorilla/mux"
)

func GetAllAssignments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var assignments []models.Assigned
	models.DB.Find(&assignments)

	json.NewEncoder(w).Encode(assignments)
}

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var assignment models.Assigned

	if err := models.DB.Where("id = ?", id).First(&assignment).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Assigned Quest not found")
		return
	}

	json.NewEncoder(w).Encode(assignment)
}

// Assign a quest
type AssignedInput struct {
	PlayerID []models.Player `json:"player_id" validate:"required"`
	QuestID  []models.Quest  `json:"quest_id" validate:"required"`
}

func AssignQuest(w http.ResponseWriter, r *http.Request) {
	var input AssignedInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	assignment := models.Assigned{
		PlayerID: input.PlayerID,
		QuestID:  input.QuestID,
	}

	models.DB.Create(&assignment)

	utils.RespondWithJSON(w, http.StatusCreated, assignment)
}
