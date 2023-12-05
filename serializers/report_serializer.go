package serializers

import (
	"go-backend/models"
)

type ReportSerializer struct {
	ID               uint               `json:"id"`
	Title            string             `json:"title"`
	Description      string             `json:"description"`
	Meetup           MeetSerializer     `json:"meetup"`
	Speaker          SpeakerSerializer  `json:"speaker"`
	Order            int                `json:"order"`
	PresentationLink string             `json:"presentation_link"`
	VideoLink        string             `json:"video_link"`
}

func NewReportSerializer(report models.Report) ReportSerializer {
	return ReportSerializer{
		ID:               report.ID,
		Title:            report.Title,
		Description:      report.Description,
		Meetup:           NewMeetSerializer(report.Meetup),
		Speaker:          NewSpeakerSerializer(report.Speaker),
		Order:            report.Order,
		PresentationLink: report.PresentationLink,
		VideoLink:        report.VideoLink,
	}
}
