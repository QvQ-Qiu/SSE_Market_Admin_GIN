package model

import (
	"time"
)

// Sue [...]
type Sue struct {
	SueID      int       `gorm:"primary_key;column:sueID"`
	Targettype string    `gorm:"column:targettype;type:enum('post','pcomment','ccomment');not null"`
	TargetID   int       `gorm:"column:targetID;type:int;not null"`
	UserID     int       `gorm:"index:psueuser;column:userID;type:int;not null"`
	User       User      `gorm:"association_foreignkey:userID;foreignkey:userID"`
	Reason     string    `gorm:"column:reason;type:varchar(1000)"`
	Time       time.Time `gorm:"column:time;type:datetime"`
	Status     string    `gorm:"column:status;type:enum(' ok','nosin','wait')"`
	Finish     bool      `gorm:"column:finish;type:tinyint(1)"`
}
