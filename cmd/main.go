package main

import (
	"log"
	"vc-sim-go/common"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()
	err := common.ParseArgs()
	if err != nil {
		return
	}
}
