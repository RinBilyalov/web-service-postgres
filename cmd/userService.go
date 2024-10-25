package cmd

import (
	"encoding/json"
	"net/http"
	"web-service-postgres/config"
	"web-service-postgres/pkg/models"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	//"application/json" говорит о том, что мы работаем с JSON REST API
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	config.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	//"application/json" говорит о том, что мы работаем с JSON REST API
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	config.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUserComments(w http.ResponseWriter, r *http.Request) {
	//"application/json" говорит о том, что мы работаем с JSON REST API
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	config.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUserCommentById(w http.ResponseWriter, r *http.Request) {
	//"application/json" говорит о том, что мы работаем с JSON REST API
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	config.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

//func CreateUser(w http.ResponseWriter, r *http.Request) {
//	var input QuestInput
//
//	body, _ := ioutil.ReadAll(r.Body)
//	_ = json.Unmarshal(body, &input)
//
//	validate = validator.New()
//	err := validate.Struct(input)
//
//	if err != nil {
//		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
//		return
//	}
//
//	quest := &models.Quest{
//		Title:       input.Title,
//		Description: input.Description,
//		Reward:      input.Reward,
//	}
//
//	models.DB.Create(quest)
//
//	w.Header().Set("Content-Type", "application/json")
//
//	json.NewEncoder(w).Encode(quest)
//
//}
