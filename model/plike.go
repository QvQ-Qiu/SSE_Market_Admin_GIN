package model

import (
	"time"
)

// Plike [...]
type Plike struct {
	PlikeID   int       `gorm:"primary_key;column:plikeID"`
	PtargetID int       `gorm:"index:pliketarget;column:ptargetID"`
	Post      Post      `gorm:"association_foreignkey:ptargetID;foreignkey:postID"`
	UserID    int       `gorm:"index:plikeuser;column:userID"`
	User      User      `gorm:"association_foreignkey:userID;foreignkey:userID"`
	Time      time.Time `gorm:"column:time;type:datetime"`
}
