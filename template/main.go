package main

import (
	"log"
	"math/rand"
	"time"
)

func init() {
	// Set new rand source to time now
	rand.NewSource(time.Now().UnixNano())
}

func main() {
	log.Println("Starting...")
	Run()
	log.Println("Ending...")
}

func Run() {

}
