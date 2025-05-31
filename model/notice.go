package model

import "time"

// Notice [...]
type Notice struct {
	NoticeID int       `gorm:"primary_key;column:noticeID"`
	Receiver int       `gorm:"index:noticereceiver;column:receiver;type:int;not null"`
	User     User      `gorm:"association_foreignkey:receiver;foreignkey:userID"`
	Sender   int       `gorm:"index:noticesender;column:sender;type:int"`
	Type     string    `gorm:"column:type;type:enum('pcomment','ccomment','punish','feedback','sue','post')"`
	Ntext    string    `gorm:"column:ntext;type:varchar(1000)"`
	Time     time.Time `gorm:"column:time;type:datetime"`
	Read     bool      `gorm:"column:read;type:tinyint(1)"`
	Target   int       `gorm:"column:target;type:int"`
}
