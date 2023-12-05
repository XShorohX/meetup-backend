package serializers

import (
	"go-backend/models"
	"time"
)

type MeetSerializer struct {
	ID                uint       `json:"id"`
	Title             string     `json:"title"`
	Description       string     `json:"description"`
	Date              time.Time  `json:"date"`
	BroadcastLink     string     `json:"broadcast_link"`
	BroadcastLogin    string     `json:"broadcast_login"`
	BroadcastPassword string     `json:"broadcast_password"`
	CountListener     int        `json:"count_listener"`
	RegisrationForm   string     `json:"regisration_form"`
}

func NewMeetSerializer(meet models.Meet) MeetSerializer {
	return MeetSerializer{
		ID:                meet.ID,
		Title:             meet.Title,
		Description:       meet.Description,
		Date:              meet.Date,
		BroadcastLink:     meet.BroadcastLink,
		BroadcastLogin:    meet.BroadcastLogin,
		BroadcastPassword: meet.BroadcastPassword,
		CountListener:     meet.CountListener,
		RegisrationForm:   meet.RegisrationForm,
	}
}