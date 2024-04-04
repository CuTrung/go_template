package main

import (
	"fmt"
	"log"
	"os"

	server "github.com/CuTrung/go_template/src"
	"github.com/CuTrung/go_template/src/common/consts"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	r := server.InitApp()
	url := fmt.Sprintf(":%v", os.Getenv(consts.PORT))
	r.Run(url)
}
