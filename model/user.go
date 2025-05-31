package model

import "time"

// User [...]
type User struct {
	UserID   int    `gorm:"primary_key;column:userID"`
	Phone    string `gorm:"column:phone;type:char(20)" `
	Email    string `gorm:"column:email;type:varchar(255)"`
	Password string `gorm:"column:password;type:varchar(255)" json:"-"`
	Name     string `gorm:"column:name;type:varchar(50)"`
	//Num       int       `gorm:"column:num;type:int"`
	Intro string `gorm:"column:intro;type:varchar(255)"`
	//IDpass    bool      `gorm:"column:idpass;type:tinyint(1)"`
	Banend    time.Time `gorm:"column:banend;type:datetime"`
	Punishnum int       `gorm:"column:punishnum;type:int"`
	AvatarURL string    `gorm:"column:avatar_url;type:varchar(255)"`
	Identity  string    `gorm:"column:identity;type:enum('student','teacher');default:'student'"`
	Score     int       `gorm:"column:score;type:int;default:0"`
}
