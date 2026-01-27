package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevindurb/web-app-tmpl/internal/app"
)

func main() {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	a := app.New(pool)
	log.Println("Running!")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), a.Mux))
}
