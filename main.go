package main

import (
	"go-backend/db"
	"go-backend/handlers"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Загрузка временной зоны Москвы
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal("Не удалось загрузить временную зону:", err)
	}

	// Установка временной зоны по умолчанию
	time.Local = loc

	db.InitDB() // Инициализация базы данных
	
	// Инициализация обработчиков
	meetHandler := &handlers.MeetHandler{DB: db.GetDB()}
	quizImageHandler := &handlers.QuizImageHandler{DB: db.GetDB()}
	reportHandler := &handlers.ReportHandler{DB: db.GetDB()}
	speakerHandler := &handlers.SpeakerHandler{DB: db.GetDB()}
	
	// Инициализация маршрутизатора Gin
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	// Эндпоинты для Meet
	meetGroup := router.Group("api/meets")
	{
		meetGroup.GET("/", meetHandler.GetMeets)
		meetGroup.GET("/:id", meetHandler.GetMeet)
		meetGroup.GET("/lastStartedMeetIfExist", meetHandler.LastStartedMeetIfExist)
		meetGroup.POST("/", meetHandler.CreateMeet)
		meetGroup.PUT("/:id", meetHandler.UpdateMeet)
		meetGroup.DELETE("/:id", meetHandler.DeleteMeet)
	}

	// Эндпоинты для QuizImage
	quizImageGroup := router.Group("/quizimages")
	{
		quizImageGroup.GET("/", quizImageHandler.GetQuizImages)
		quizImageGroup.POST("/", quizImageHandler.CreateQuizImage)
		quizImageGroup.PUT("/:id", quizImageHandler.UpdateQuizImage)
		quizImageGroup.DELETE("/:id", quizImageHandler.DeleteQuizImage)
	}

	// Эндпоинты для Report
	reportGroup := router.Group("api/reports")
	{
		reportGroup.GET("/", reportHandler.GetReports)
		reportGroup.GET("/:id", reportHandler.GetReport)
		reportGroup.POST("/", reportHandler.CreateReport)
		reportGroup.PUT("/:id", reportHandler.UpdateReport)
		reportGroup.DELETE("/:id", reportHandler.DeleteReport)
	}

	// Эндпоинты для Speaker
	speakerGroup := router.Group("api/speakers")
	{
		speakerGroup.GET("/", speakerHandler.GetSpeakers)
		speakerGroup.GET("/:id", speakerHandler.GetSpeaker)
		speakerGroup.POST("/", speakerHandler.CreateSpeaker)
		speakerGroup.PUT("/:id", speakerHandler.UpdateSpeaker)
		speakerGroup.DELETE("/:id", speakerHandler.DeleteSpeaker)
	}

	// Запуск сервера
	router.Run(":8000") // Порт, на котором запускается сервер
}