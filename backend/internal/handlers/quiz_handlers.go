package handlers

import (
	"net/http"
	"quiz/internal/models"
	"quiz/internal/repository"

	"github.com/labstack/echo/v4"
)

type QuizHandler struct {
	QuizRepo repository.QuizRepository
}

func (h *QuizHandler) CreateQuiz(c echo.Context) error {
	quiz := new(models.Quiz)

	if err := c.Bind(quiz); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.QuizRepo.Create(quiz); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, quiz)
}

func (h *QuizHandler) GetQuizzes(c echo.Context) error {
	quizzes, err := h.QuizRepo.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var quizzesResponses []models.QuizResponse
	for _, quiz := range quizzes {
		quizzesResponses = append(quizzesResponses, quiz.ToResponse())
	}

	return c.JSON(http.StatusOK, map[string]any{"quizzes": quizzesResponses})
}
