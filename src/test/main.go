package main

import (
	"Backlog2Teams/src/server"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Panic("env file cannot read")
	}

	r := server.GetRouter()
	r.Run(":8080")
}
