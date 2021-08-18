package main

import (
	"log"
	"os"
	"os/signal"

	"go-restful/internal/data"
	"go-restful/internal/server"

	"github.com/joho/godotenv"
)

func main() {

	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Read the port from environment file
	port := envs["PORT"]

	//Crate a new server
	server, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	// connection to the database.
	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	//Start the server
	go server.Start()

	//Wait for an interruption.
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	<-channel

	//then
	server.Close()
	data.Close()
}
