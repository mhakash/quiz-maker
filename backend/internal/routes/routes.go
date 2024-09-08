package routes

import (
	"quiz/internal/handlers"
	"quiz/internal/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	quizRepo := repository.New(db)

	quizHandler := &handlers.QuizHandler{QuizRepo: quizRepo}

	e.POST("/quiz", quizHandler.CreateQuiz)
	e.GET("/quiz", quizHandler.GetQuizzes)
}
