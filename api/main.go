package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/confus1on/UKM/internal/postgres"
	profileHandler "github.com/confus1on/UKM/internal/service/profile/handler"
	profileRepo "github.com/confus1on/UKM/internal/service/profile/repository"
	profileUC "github.com/confus1on/UKM/internal/service/profile/usecase"
	usrHandler "github.com/confus1on/UKM/internal/service/user/handler"
	usrRepo "github.com/confus1on/UKM/internal/service/user/repository"
	usrUC "github.com/confus1on/UKM/internal/service/user/usecase"

	"github.com/buaazp/fasthttprouter"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
)

func main() {
	// load environment variable from .env
	_ = godotenv.Load()

	// get port in environment
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	port = ":" + port

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
		err := client.Schema.Create(ctx)
		if err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}

		err = postgres.InitValueRoleDB(client)
		if err != nil {
			log.Fatalf("failed initialize value role: %v", err)
		}
	}

	// initialize user service
	userRepo := usrRepo.NewUserRepository(client)
	userUsecase := usrUC.NewUserUsecase(userRepo)
	user := usrHandler.NewUserHandler(userUsecase)

	// initialize profile service
	profileRepo := profileRepo.NewProfileRepository(client)
	profileUsecase := profileUC.NewProfileUsecase(profileRepo)
	profile := profileHandler.NewProfileHandler(profileUsecase)

	router := fasthttprouter.New()
	router.POST("/api/v0/user/register", user.RegisterUser)
	router.POST("/api/v0/user/login", user.LoginUser)

	router.GET("/api/v0/profile/:email", profile.GetDetailProfile)
	router.PUT("/api/v0/profile/:email", profile.UpdateProfile)

	log.Printf("Server Running at http://localhost%v \n", port)
	log.Fatal(fasthttp.ListenAndServe(port, router.Handler))
}
