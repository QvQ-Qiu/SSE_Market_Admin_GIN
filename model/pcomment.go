package model

import (
	"time"
)

// Pcomment [...]
type Pcomment struct {
	PcommentID int       `gorm:"primary_key;column:pcommentID"`
	UserID     int       `gorm:"index:pcommentuser;column:userID;type:int;not null"`
	User       User      `gorm:"association_foreignkey:userID;foreignkey:userID"`
	PtargetID  int       `gorm:"index:pcommenttarget;column:ptargetID;type:int;not null"`
	Post       Post      `gorm:"association_foreignkey:ptargetID;foreignkey:postID"`
	LikeNum    int       `gorm:"column:like_num;type:int"`
	DenyNum    int 		 `gorm:"column:deny_num;type:int"`
	Pctext     string    `gorm:"column:pctext;type:varchar(1000)"`
	Time       time.Time `gorm:"column:time;type:datetime"`
}
