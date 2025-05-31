package model

import (
	"time"
)

// Ccomment [...]
type Ccomment struct {
	CcommentID     int       `gorm:"primary_key;column:ccommentID"`
	UserID         int       `gorm:"index:ccommentuser;column:userID;type:int;not null"`
	User           User      `gorm:"association_foreignkey:userID;foreignkey:userID"`
	CtargetID      int       `gorm:"index:ccommenttarget;column:ctargetID;type:int;not null"`
	Pcomment       Pcomment  `gorm:"association_foreignkey:ctargetID;foreignkey:pcommentID"`
	LikeNum        int       `gorm:"column:like_num;type:int"`
	DenyNum		   int  	 `gorm:"column:deny_num;type:int"`
	Cctext         string    `gorm:"column:cctext;type:varchar(100)"`
	Time           time.Time `gorm:"column:time;type:datetime"`
	UserTargetName string    `gorm:"column:usertargetName;type:varchar(50)"`
}
