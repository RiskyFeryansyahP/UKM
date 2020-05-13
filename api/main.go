package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/confus1on/UKM/internal/postgres"
	announcementHandler "github.com/confus1on/UKM/internal/service/announcement/handler"
	announcementRepo "github.com/confus1on/UKM/internal/service/announcement/repository"
	announcementUC "github.com/confus1on/UKM/internal/service/announcement/usecase"
	profileHandler "github.com/confus1on/UKM/internal/service/profile/handler"
	profileRepo "github.com/confus1on/UKM/internal/service/profile/repository"
	profileUC "github.com/confus1on/UKM/internal/service/profile/usecase"
	ukmHandler "github.com/confus1on/UKM/internal/service/ukm/handler"
	ukmRepo "github.com/confus1on/UKM/internal/service/ukm/repository"
	ukmUC "github.com/confus1on/UKM/internal/service/ukm/usecase"
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

		err = postgres.InitDefaultValue(client)
		if err != nil {
			log.Fatalf("failed initialize default value: %v", err)
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

	// initialize ukm service
	ukmRepo := ukmRepo.NewRepositoryUKM(client)
	ukmUsecase := ukmUC.NewUKMUsecase(ukmRepo)
	ukm := ukmHandler.NewUKMHandler(ukmUsecase)

	// initialize announcement service
	announcementRepo := announcementRepo.NewAnnouncementRepository(client)
	announcementUsecase := announcementUC.NewAnnouncementUsecase(announcementRepo)
	announcement := announcementHandler.NewAnnouncementHandler(announcementUsecase)

	router := fasthttprouter.New()
	router.POST("/api/v0/user/register", user.RegisterUser)
	router.POST("/api/v0/user/login", user.LoginUser)

	router.GET("/api/v0/profile/:email", profile.GetDetailProfile)
	router.PUT("/api/v0/profile/:email", profile.UpdateProfile)

	router.GET("/api/v0/ukm", ukm.GetAllUKM)
	router.POST("/api/v0/ukm/register/profile/:profileID", ukm.RegisterUKM)
	router.PUT("/api/v0/ukm/:ukmID", ukm.UpdateUKM)

	router.GET("/api/v0/announcement/ukm/:ukmID", announcement.GetAnnouncement)
	router.POST("/api/v0/announcement/ukm/:ukmID", announcement.PostAnnouncement)
	router.PUT("/api/v0/announcement/:announcementID", announcement.UpdateAnnouncement)
	router.DELETE("/api/v0/announcement/:announcementID", announcement.DeleteAnnouncement)

	log.Printf("Server Running at http://localhost%v \n", port)
	log.Fatal(fasthttp.ListenAndServe(port, router.Handler))
}
