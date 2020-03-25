package main

import (
	"context"
	"flag"
	"log"

	"github.com/RiskyFeryansyahP/UKM/internal/postgres"
	"github.com/RiskyFeryansyahP/UKM/internal/service/user/handler"
	"github.com/RiskyFeryansyahP/UKM/internal/service/user/repository"
	"github.com/RiskyFeryansyahP/UKM/internal/service/user/usecase"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
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
	flag.Parse()
	log.Print(*migrate)

	// Run migration
	if *migrate {
		if err := client.Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}

	userRepo := repository.NewUserRepository(client)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	router := fasthttprouter.New()
	router.POST("/api/v0/user/register", userHandler.RegisterUser)
	router.POST("/api/v0/user/login", userHandler.LoginUser)

	log.Printf("Server Running at http://localhost%v \n", port)
	log.Fatal(fasthttp.ListenAndServe(port, router.Handler))
}
