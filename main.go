package main

import (

	"github.com/HafidurR/final-prjk/src/config"
	"github.com/HafidurR/final-prjk/src/routes"
	"github.com/joho/godotenv"
)


func main()  {
	loadEnv()
	config.Connect()
	routes.HandleRequest()
}

func loadEnv(){
	err := godotenv.Load()
  	if err != nil {
    	panic("Error loading .env file")
  	}
}