package model

type Pcdeny struct{
	PcdenyID   int      `gorm:"primary_key;column:pcdenyID"`
	PctargetID int      `gorm:"index:pcdenytarget;column:pctargetID;type:int;not null"`
	Pcomment   Pcomment `gorm:"association_foreignkey:pctargetID;foreignkey:pcommentID"`
	UserID     int      `gorm:"index:pcdenyuser;column:userID;type:int;not null"`
	User       User     `gorm:"association_foreignkey:userID;foreignkey:userID"`
}
