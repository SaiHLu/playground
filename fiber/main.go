package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/playground/fiber/database"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Could not load Env file.", err.Error())
	}

	database.SetupDB()

	fmt.Println("Hello World")
}
