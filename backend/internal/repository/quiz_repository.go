package repository

import (
	"quiz/internal/models"

	"gorm.io/gorm"
)

type QuizRepository interface {
	Create(quiz *models.Quiz) error
	Get() ([]models.Quiz, error)
}

type quizRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) QuizRepository {
	return &quizRepository{db}
}

func (r *quizRepository) Create(quiz *models.Quiz) error {
	return r.db.Create(quiz).Error
}

func (r *quizRepository) Get() ([]models.Quiz, error) {
	var quizzes []models.Quiz
	if err := r.db.Preload("Tags").Find(&quizzes).Error; err != nil {
		return nil, err
	}
	return quizzes, nil
}
