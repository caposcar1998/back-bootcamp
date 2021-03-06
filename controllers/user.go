package controllers

import (
	"backbootcamp/models"
	"backbootcamp/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, DELETE, PUT, POST, OPTIONS")
}

func (h *UserController) Login(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var userToLogin models.User
	var err error
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &userToLogin)
	user, err := h.userService.Login(userToLogin.Email, userToLogin.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(user)
	}

}

func (h *UserController) GetUsersList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var userList []models.User
	var err error
	userList, err = h.userService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(userList)
	}
}

func (h *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var user *models.User
	var err error
	query := r.URL.Query()
	idStr := query.Get("id")
	idInt, _ := strconv.Atoi(idStr)
	user, err = h.userService.GetUser(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func (h *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var userToCreate models.User
	var err error
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &userToCreate)
	user, err := h.userService.InsertUser(&userToCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

func (h *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var userToUpdate models.User
	var err error
	query := r.URL.Query()
	idStr := query.Get("id")
	idInt, _ := strconv.Atoi(idStr)
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &userToUpdate)
	user, err := h.userService.UpdateUser(idInt, &userToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func (h *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var err error
	query := r.URL.Query()
	idStr := query.Get("id")
	idInt, _ := strconv.Atoi(idStr)
	err = h.userService.DeleteUser(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
