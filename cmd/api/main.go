package main

import (
	"aurh-service/internal/config"
	"aurh-service/internal/routes"
	"log"
	"net/http"
)

func main(){
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

