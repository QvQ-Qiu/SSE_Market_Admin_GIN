package model

type Psave struct {
	PsaveID   int  `gorm:"primary_key;column:psaveID"`
	PtargetID int  `gorm:"index:psavetarget;column:ptargetID;type:int;not null"`
	Post      Post `gorm:"association_foreignkey:ptargetID;foreignkey:postID"`
	UserID    int  `gorm:"index:plikeuser;column:userID;type:int;not null"`
	User      User `gorm:"association_foreignkey:userID;foreignkey:userID"`
}
