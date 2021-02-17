package controller

import (
	"encoding/json"
	"net/http"

	"github.com/samsammn/project7-test/model"
	"github.com/samsammn/project7-test/repository"
	"github.com/samsammn/project7-test/response"
	"github.com/samsammn/project7-test/router"
)

type UserController interface {
	FindById(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type userControllerImpl struct {
	repo repository.UserRepository
}

func NewUserController(repo repository.UserRepository) UserController {
	return &userControllerImpl{repo: repo}
}

func (ctl *userControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	users := ctl.repo.GetAll()

	response.Json(w, response.Map{
		"data": users,
	})
}

func (ctl *userControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	id := router.Arg(r)
	user := ctl.repo.GetById(id)

	response.Json(w, response.Map{
		"data": user,
	})
}

func (ctl *userControllerImpl) Store(w http.ResponseWriter, r *http.Request) {
	userInput := model.User{}

	json.NewDecoder(r.Body).Decode(&userInput)
	ctl.repo.Save(userInput)

	response.Json(w, response.Map{
		"msg": "data inserted successfully",
	})
}

func (ctl *userControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	id := router.Arg(r)
	userInput := model.User{}

	json.NewDecoder(r.Body).Decode(&userInput)
	ctl.repo.Update(id, userInput)

	response.Json(w, response.Map{
		"msg": "data updated successfully",
	})
}

func (ctl *userControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	id := router.Arg(r)
	ctl.repo.Delete(id)

	response.Json(w, response.Map{
		"msg": "data deleted successfully",
	})
}

func (ctl *userControllerImpl) Login(w http.ResponseWriter, r *http.Request)  {}
func (ctl *userControllerImpl) Logout(w http.ResponseWriter, r *http.Request) {}
