package main

import (
	"auth-service/internal/config"
	"auth-service/internal/db"
	"auth-service/internal/routes"
	"log"
	"net/http"
)

func main(){
	db.Connect()
	//initialize router
	router := config.SetUpRouter()

	//register router
	routes.RegisterRoutes(router)

	//start server
	port :=":8080"
	log.Println("Server is running on port",port)

	err := http.ListenAndServe(port,router)
	if err != nil {
		log.Fatal("Server failed:",err)
	}
}

