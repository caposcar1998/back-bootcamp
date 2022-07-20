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

	router.HandleFunc("/api/v1/users", userController.GetUsersList).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/users/{id}", userController.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/users", userController.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/users/{id}", userController.UpdateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/v1/users/{id}", userController.DeleteUser).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/v1/users/login", userController.Login).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", router))
}
