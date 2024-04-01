package main

import (
	"fmt"
	"log"
	"os"

	server "github.com/CuTrung/go_template/src"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	r := server.InitApp()
	url := fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("PORT"))
	r.Run(url)
}
