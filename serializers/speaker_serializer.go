package serializers

import (
	"go-backend/models"
)

type SpeakerSerializer struct {
	ID          uint   `json:"id"`
	LastName    string `json:"last_name"`
	FirstName   string `json:"first_name"`
	Surname     string `json:"surname"`
	Photo       string `json:"photo"`
	SpeakerInfo string `json:"speaker_info"`
}

func NewSpeakerSerializer(speaker models.Speaker) SpeakerSerializer {
	return SpeakerSerializer{
		ID:          speaker.ID,
		LastName:    speaker.LastName,
		FirstName:   speaker.FirstName,
		Surname:     speaker.Surname,
		Photo:       speaker.Photo,
		SpeakerInfo: speaker.SpeakerInfo,
	}
}