package models

import (
	"time"

	"gorm.io/gorm"
)

type Meet struct {
	gorm.Model
	Title             string 	`gorm:"size:max_length;column:title"`
	Description       string 	`gorm:"column:description"`
	Date              time.Time `gorm:"not null"`
	BroadcastLink     string 	`gorm:"column:broadcast_link"`
	BroadcastLogin    string 	`gorm:"size:max_length;column:broadcast_login"`
	BroadcastPassword string 	`gorm:"size:max_length;column:broadcast_password"`
	CountListener     int    	`gorm:"column:count_listener"`
	RegisrationForm   string 	`gorm:"column:regisration_form"`
}

type QuizImage struct {
	gorm.Model
	Images string `gorm:"column:images"`
	Meetup uint   `gorm:"column:meetup"`
}
