package main

import (
	"./database"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.CreateTables()
	database.PopulateCurrencyValues()
	database.PopulateExpenseCategoryValues()
	database.PopulateIncomeCategoryValues()
}
