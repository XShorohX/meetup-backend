package models

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	Title             string `gorm:"size:max_length;column:title"`
	Description       string `gorm:"column:description"`
	MeetupID          uint
	Meetup            Meet `gorm:"foreignKey:MeetupID"`
	SpeakerID         uint
	Speaker           Speaker `gorm:"foreignKey:SpeakerID"`
	Order             int    `gorm:"column:order"`
	PresentationLink  string `gorm:"column:presentation_link"`
	VideoLink         string `gorm:"column:video_link"`
}