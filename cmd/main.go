package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/Wai30Yan/cna-server/config"
	"github.com/Wai30Yan/cna-server/driver"
	"github.com/Wai30Yan/cna-server/handlers"
)
const (
	dbport = "5432"
)

var app config.AppConfig

func main() {
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("USER")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, dbport, user, password, dbname)
	db, err := driver.ConnectDB(psqlInfo)

	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to database")

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandler(repo)

	s := &http.Server{
		Addr: ":8080",
		Handler: routes(),
	}

	fmt.Println("server running on :8080")
	
	log.Fatal(s.ListenAndServe()) 
}