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
	db := database.NewMySQLConn()
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	router := mux.NewRouter()
	fmt.Println("Server is running on Port 8002")

	router.HandleFunc("/api/v1/users", userController.GetUsersList).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", userController.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/users/{id}", userController.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8002", router))
}
