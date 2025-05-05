package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"layersapi/controllers"
	"layersapi/repositories/memory"
	"layersapi/services"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	userRepository := memory.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(*userService)

	r.HandleFunc("/users", userController.GetAllUsersHandler).Methods(http.MethodGet)
	r.HandleFunc("/users", userController.GetAllUsersHandler).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", userController.GetUserByIdHandler).Methods(http.MethodGet)
	r.HandleFunc("/users", userController.CreateUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", userController.UpdateUserHandler).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", userController.DeleteUserHandler).Methods(http.MethodDelete)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)

}
