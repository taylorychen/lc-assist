package main

import (
	"fmt"
	utils "lc-assist/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env.local")
	if err != nil {
		fmt.Printf("%v\n", err)
		log.Fatal("Error loading .env file")
	}

	questionsFile, err := os.Create("questionsList.csv")
	if err != nil {
		log.Fatal("Failed to create file")
	}
	defer questionsFile.Close()

	utils.GetData(1000, questionsFile, false)

	// utils.ParseQuestions()
	// host := os.Getenv("POSTGRES_HOST")
	// user := os.Getenv("POSTGRES_USER")
	// dbname := os.Getenv("POSTGRES_DATABASE")
	// password := os.Getenv("POSTGRES_PASSWORD")

	// psqlInfo := fmt.Sprintf("host=%s user=%s "+
	// 	"password=%s dbname=%s  sslmode=require",
	// 	host, user, password, dbname)

	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	log.Fatal("Unable to connect to db")
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// 	log.Fatal("Unable to ping db")
	// }

	// res, err := db.Exec()
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// 	log.Fatal("Unable to ping db")
	// }
	// fmt.Printf("%v\n", res)

	// fmt.Println("Created Table")

}
