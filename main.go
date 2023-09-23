package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/ChaosIsFramecode/horinezumi/data"
	"github.com/ChaosIsFramecode/horinezumi/subroutes/wikiroute"
)

func main() {
	// Load .env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables: %s", err)
	}

	// Connect to our database
	conn, err := data.ConnectToDataBase()
	if err != nil {
		log.Fatalf("Error connecting to data base: %s", err)
	} else {
		log.Printf("Successfully connected to data base")
	}
	defer data.CloseDataBase(conn)

	if err = data.CreateTables(conn); err != nil {
		log.Fatalf("Failed to create table: %s", err)
	}

	rt := chi.NewRouter()

	// Use logger
	rt.Use(middleware.Logger)

	// Redirect root path to main page
	rt.Get("/", http.RedirectHandler("/wiki/Main_Page", http.StatusSeeOther).ServeHTTP)

	// Wiki sub router
	wikiroute.SetupWikiroute(rt, conn)

	log.Println("Running on " + os.Getenv("ADDR"))
	http.ListenAndServe(os.Getenv("ADDR"), rt)
}
