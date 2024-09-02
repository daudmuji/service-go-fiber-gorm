package main

import (
	"fmt"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"golang-template-service/config"
	"golang-template-service/controller"
	"golang-template-service/repository"
	"golang-template-service/usecase"
	util "golang-template-service/util/db"

	"golang-template-service/helpers"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"time"
)

func main() {
	app := fiber.New()

	os.Setenv("TZ", "Asia/Jakarta")

	apiGroup := app.Group("/go-service-fiber")

	// Read config
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	env := os.Getenv("ENVIRONMENT_APPLICATION")
	fmt.Println("running in env: ", env)
	if env == "" {
		env = "local"
	}

	cfg, err := config.Read(fmt.Sprintf("environment/%s.env", env))
	if err != nil {
		log.Fatalln("read config err:", err)
	}

	app.Use(recover.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		AllowMethods: "OPTIONS, GET, POST, PUT, DELETE",
	}))
	helpers.CheckFolderPath("log")

	logApps := lumberjack.Logger{
		Filename:  "log/apps.log",
		MaxAge:    30, //days
		LocalTime: true,
	}

	logSystems := lumberjack.Logger{
		Filename:  "log/system.log",
		MaxAge:    30, //days
		LocalTime: true,
	}

	go func() {
		for {
			times := time.Now().Format("15:04:05")
			if times == "23:59:59" {
				logApps.Rotate()
				logSystems.Rotate()
				time.Sleep(time.Hour * 24)
			}
		}
	}()

	multiapps := zerolog.MultiLevelWriter(os.Stdout, &logApps)

	loggerApps := zerolog.New(multiapps).With().Str("LogType", "applog").Timestamp().Logger()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &loggerApps,
		Fields: []string{helpers.FieldLatency, helpers.FieldStatus, helpers.FieldMethod, helpers.FieldURL, helpers.FieldError, helpers.FieldQueryParams, helpers.FieldBody, helpers.FieldReqHeaders, helpers.FieldResBody, helpers.FieldReqHeaders},
	}))

	multisys := zerolog.MultiLevelWriter(os.Stdout, &logSystems)

	log.SetFlags(log.Ldate | log.Lshortfile)
	log.SetOutput(zerolog.New(multisys).With().Str("LogType", "syslog").Timestamp().Logger())

	//init db
	connection, err := util.New(cfg.Database)
	//articleRepo := repository.NewArticleRepositoryPostgres(connection.DbConnection)
	uploadRepo := repository.NewSampleUploadPostgresRepository(connection.DbConnection)
	productRepo := repository.NewProductRepositoryPostgres(connection.DbConnection)

	//articleUsecase := usecase.NewArticleUsecase(articleRepo)
	validationUsecase := usecase.NewValidationUsecase()
	uploadUsecase := usecase.NewSampleUploadUsecase(uploadRepo, validationUsecase)
	productUsecase := usecase.NewProductUsecase(productRepo)

	//controller.NewArticleController(apiGroup, articleUsecase)
	controller.NewSampleUploadController(apiGroup, uploadUsecase)
	controller.NewProductController(apiGroup, productUsecase)

	log.Fatal(app.Listen(cfg.Http.Port))
}
