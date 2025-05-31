package model

import "time"

// Feedback [...]
type Feedback struct {
	FeedbackID int    `gorm:"primary_key;column:feedbackID"`
	Ftext      string `gorm:"column:ftext;type:varchar(1000)"`
	Attachment string `gorm:"column:attachment;type:varchar(255)"`
	//UserID     int       `gorm:"index:feedbackuser;column:userID;type:int"`
	Time   time.Time `gorm:"column:time;type:datetime"`
	Status string    `gorm:"column:status;type:enum('ok','wait')"`
}
