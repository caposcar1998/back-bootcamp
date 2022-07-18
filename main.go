package main

import (
	"fmt"
	"log"
	"net/http"

	"backbootcamp/controllers"
	"backbootcamp/database"
	"backbootcamp/repositories"
	"backbootcamp/services"

	"github.com/gorilla/mux"
)

func main() {
	db, _ := database.NewMySQLConn()
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	router := mux.NewRouter()
	fmt.Println("Server is running on Port 8080")

	router.HandleFunc("/api/v1/users", userController.GetUsersList).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", userController.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/users/{id}", userController.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/v1/login", userController.Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
