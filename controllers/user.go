package controllers

import (
	"backbootcamp/models"
	"backbootcamp/services"
	"encoding/json"
	"fmt"
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

// func (h *UserController) Login(w http.ResponseWriter, r *http.Request) {
// 	var err error
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	json.Unmarshal(reqBody, &userToCreate)
// 	email := "luis"
// 	password := "luis"
// 	user, err := h.userService.Login(email, password)
// 	if err != nil {
// 		fmt.Printf("ERROR - %s", err)
// 	}
// 	json.NewEncoder(w).Encode(user)
// }

func (h *UserController) GetUsersList(w http.ResponseWriter, r *http.Request) {
	var userList []models.User
	var err error
	userList, err = h.userService.GetAllUsers()
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	json.NewEncoder(w).Encode(userList)
}

func (h *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	var err error
	query := r.URL.Query()
	idStr := query.Get("id")
	idInt, _ := strconv.Atoi(idStr)
	user, err = h.userService.GetUser(idInt)
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userToCreate models.User
	var err error
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &userToCreate)
	user, err := h.userService.InsertUser(&userToCreate)
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userToUpdate models.User
	var err error
	query := r.URL.Query()
	idStr := query.Get("id")
	idInt, _ := strconv.Atoi(idStr)
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &userToUpdate)
	user, err := h.userService.UpdateUser(idInt, &userToUpdate)
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var err error
	query := r.URL.Query()
	idStr := query.Get("id")
	idInt, _ := strconv.Atoi(idStr)
	err = h.userService.DeleteUser(idInt)
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
}
