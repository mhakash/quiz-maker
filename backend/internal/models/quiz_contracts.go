package models

type QuizResponse struct {
	ID       uint              `json:"id"`
	Question string            `json:"question"`
	Options  map[string]string `json:"options"`
	Answer   string            `json:"answer"`
	Tags     []TagResponse     `json:"tags"`
}

type TagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
