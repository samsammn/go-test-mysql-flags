package controller

import (
	"encoding/json"
	"net/http"

	"github.com/samsammn/project7-test/model"
	"github.com/samsammn/project7-test/repository"
	"github.com/samsammn/project7-test/response"
)

type UserController interface {
	FindAll(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
}

type userControllerImpl struct {
	repo repository.UserRepository
}

func NewUserController(repo repository.UserRepository) UserController {
	return &userControllerImpl{repo: repo}
}

func (ctl *userControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	users := ctl.repo.GetAll()
	response.Json(w, users)
}

func (ctl *userControllerImpl) Store(w http.ResponseWriter, r *http.Request) {
	userInput := model.User{}
	json.NewDecoder(r.Body).Decode(&userInput)

	ctl.repo.Save(userInput)
	response.Json(w, response.Map{
		"msg": "data berhasil ditambahkan!",
	})
}
