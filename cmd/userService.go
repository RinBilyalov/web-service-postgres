package cmd

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"web-service-postgres/config"
	"web-service-postgres/pkg/models"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	config.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(findUser(w, r))
}

func GetUserCommentById(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(findUser(w, r).Comment)
}

func GetUserComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var comments []models.Comment
	config.DB.Find(&comments)
	json.NewEncoder(w).Encode(comments)
}

var validate *validator.Validate

type UserInput struct {
	ID    uint   `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var input UserInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		config.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	quest := &models.User{
		Name:  input.Name,
		Email: input.Email,
	}

	config.DB.Create(quest)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(quest)

}

func findUser(w http.ResponseWriter, r *http.Request) models.User {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	var user models.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		config.RespondWithError(w, http.StatusNotFound, "User not found")
	}
	return user
}
