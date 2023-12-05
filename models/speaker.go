package models

import "gorm.io/gorm"

type Speaker struct {
	gorm.Model
	LastName    string `gorm:"size:max_length;column:last_name"`
	FirstName   string `gorm:"size:max_length;column:first_name"`
	Surname     string `gorm:"size:max_length;column:surname"`
	Photo       string `gorm:"column:photo"`
	SpeakerInfo string `gorm:"size:max_length;column:speaker_info"`
}

func (s *Speaker) FullName() string {
	return s.LastName + " " + s.FirstName + " " + s.Surname
}
