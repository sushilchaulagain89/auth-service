package main

import (
	"auth-service/internal/config"
	"auth-service/internal/db"
	"auth-service/internal/handler"
	"auth-service/internal/repository"
	"auth-service/internal/routes"
	"auth-service/internal/service"
	"log"
	"net/http"
)

func main(){
	// database connection
	db.Connect()

	// dependency build
	userRepo := &repository.UserRepository{}
	userService := &service.UserService{
		Repo: userRepo,
	}

	userHandler := &handler.UserHandler{
		Service: userService,
	}
	//initialize router
	router := config.SetUpRouter()

	//register router
	routes.RegisterRoutes(router,userHandler)

	//start server
	port :=":8080"
	log.Println("Server is running on port",port)

	err := http.ListenAndServe(port,router)
	if err != nil {
		log.Fatal("Server failed:",err)
	}
}

