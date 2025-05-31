package model

import "time"

// CDKey [...]
type CDKey struct {
	CDKeyID     int       `gorm:"primary_key;column:cdkeyID"`
	Content     string    `gorm:"column:content;type:char(9)"`
	Used        bool      `gorm:"column:used;type:tinyint(1)"`
	CreatedTime time.Time `gorm:"column:createdtime;type:datetime"`
	UsedTime    time.Time `gorm:"column:usedtime;type:datetime"`
}
