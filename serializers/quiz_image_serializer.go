package serializers

import "go-backend/models"

type QuizImageSerializer struct {
	Images  string `json:"images"`
	Meetup  uint   `json:"meetup"`
}

func NewQuizImageSerializer(quizImage models.QuizImage) QuizImageSerializer {
	return QuizImageSerializer{
		Images:  quizImage.Images,
		Meetup:  quizImage.Meetup,
	}
}