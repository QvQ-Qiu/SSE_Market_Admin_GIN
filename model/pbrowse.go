package model

import (
	"time"
)

// Pbrowse [...]
type Pbrowse struct {
	PbrowseID   int       `gorm:"primary_key;column:pbrowseID"`
	PtargetID int       `gorm:"index:pbrowsetarget;column:ptargetID"`
	Post      Post      `gorm:"association_foreignkey:ptargetID;foreignkey:postID"`
	UserID    int       `gorm:"index:pbrowseuser;column:userID"`
	User      User      `gorm:"association_foreignkey:userID;foreignkey:userID"`
	Time      time.Time `gorm:"column:time;type:datetime"`
}
