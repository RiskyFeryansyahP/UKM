package main

import (
	"context"
	"flag"
	"github.com/buaazp/fasthttprouter"
	"github.com/confus1on/UKM/internal/postgres"
	"github.com/valyala/fasthttp"
	"log"
)

const port = ":8080"

func main() {
	// Connect to database
	client, err := postgres.NewPostgreSQL()
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	migrate := flag.Bool("migrate", false, "a migrate database")
	log.Print(*migrate)

	// Run migration
	if *migrate {
		if err := client.Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}

	router := fasthttprouter.New()

	log.Printf("Server Running at localhost%v \n", port)
	log.Fatal(fasthttp.ListenAndServe(port, router.Handler))
}
