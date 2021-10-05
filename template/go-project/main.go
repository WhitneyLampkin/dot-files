package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hunter32292/weeklyProject/template/pkg/database"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Println("Starting...")
	Execute()
	log.Println("Exiting ...")
}

// Execute - Run the project
func Execute() {
	err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	err = database.CreateUsersTable()
	if err != nil {
		log.Fatal(err)
	}
	err = database.AddNewUser("John", "CXP")
	if err != nil {
		log.Fatal(err)
	}
	user, err := database.GetUser(1)
	if err != nil {
		log.Fatal(err)
	}
	err = database.DropUsersTable()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", user)
}
