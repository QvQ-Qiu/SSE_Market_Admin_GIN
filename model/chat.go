package model

import "time"

type ChatMsg struct {
	ChatMsgID    int       `gorm:"primary_key;column:chatMsgID" json:"chatMsgID"`
	TargetUserID int       `gorm:"column:targetUserID;type:int;not null;index:idx_target_sender" json:"targetUserID"`
	TargetUser   User      `gorm:"foreignkey:TargetUserID;references:UserID" json:"targetUser"`
	SenderUserID int       `gorm:"column:senderUserID;type:int;not null;index:idx_target_sender" json:"senderUserID"`
	SenderUser   User      `gorm:"foreignkey:SenderUserID;references:UserID" json:"senderUser"`
	Content      string    `gorm:"column:content;type:varchar(5000)" json:"content"`
	CreatedAt    time.Time `gorm:"column:createdAt;type:datetime" json:"createdAt"`
}
