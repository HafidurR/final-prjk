package main

import (
	"fmt"
	"os"

	"github.com/HafidurR/final-prjk/src/config"
	"github.com/HafidurR/final-prjk/src/routes"
	"github.com/joho/godotenv"
)


func main()  {
	loadEnv()
	fmt.Println("ya error", os.Getenv("DB_NAME"))	
	config.Connect()
	routes.HandleRequest()
}

func loadEnv(){
	err := godotenv.Load()
  	if err != nil {
    	panic("Error loading .env file")
  	}
}