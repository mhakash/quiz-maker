package models

import (
	"encoding/json"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	Question string         `json:"question"`
	Options  datatypes.JSON `json:"options"`
	Answer   string         `json:"answer"`
	Tags     []Tag          `json:"tags" gorm:"many2many:quiz_tags;"`
}

type Tag struct {
	gorm.Model
	Name    string `json:"name"`
	Quizzes []Quiz `json:"-" gorm:"many2many:quiz_tags;"`
}

func (q *Quiz) ToResponse() QuizResponse {
	var options map[string]string
	if err := json.Unmarshal(q.Options, &options); err != nil {
		return QuizResponse{}
	}

	var tags []TagResponse
	for _, tag := range q.Tags {
		tags = append(tags, TagResponse{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}

	return QuizResponse{
		ID:       q.ID,
		Question: q.Question,
		Options:  options,
		Answer:   q.Answer,
		Tags:     tags,
	}
}
