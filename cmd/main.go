package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/Wai30Yan/cna-server/pkg/config"
	"github.com/Wai30Yan/cna-server/pkg/driver"
	"github.com/Wai30Yan/cna-server/pkg/handlers"
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
	port := os.Getenv("PORT")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, dbport, user, password, dbname)
	db, err := driver.ConnectDB(psqlInfo)

	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to database")

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandler(repo)

	if port == "" {
		port = "8080"
	}

	s := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", port),
		Handler: routes(),
	}

	fmt.Println("server running on :8080")
	
	log.Fatal(s.ListenAndServe()) 
}