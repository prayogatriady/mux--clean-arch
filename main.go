package main

import (
	"fmt"
	"go-rest-api/app"
	"go-rest-api/controller"
	"go-rest-api/helper"
	"go-rest-api/middleware"
	"go-rest-api/repository"
	"go-rest-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, *validate)
	userController := controller.NewUserController(userService)

	router := mux.NewRouter().StrictSlash(true)

	router.Use(middleware.AuthMiddleware)

	router.HandleFunc("/api/users", userController.FindAllUser).Methods("GET")
	router.HandleFunc("/api/profile", userController.Profile).Methods("GET")
	router.HandleFunc("/api/signup", userController.Signup).Methods("POST")
	router.HandleFunc("/api/edit", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/delete", userController.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/login", userController.Login).Methods("POST")
	router.HandleFunc("/api/logout", userController.Logout).Methods("POST")

	fmt.Println("Server is running on port 3000")
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe() // untuk menjalan servernya
	helper.PanicIfErr(err)
}
