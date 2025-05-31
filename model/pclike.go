package model

// Pclike [...]
type Pclike struct {
	PclikeID   int      `gorm:"primary_key;column:pclikeID"`
	PctargetID int      `gorm:"index:pcliketarget;column:pctargetID;type:int;not null"`
	Pcomment   Pcomment `gorm:"association_foreignkey:pctargetID;foreignkey:pcommentID"`
	UserID     int      `gorm:"index:pclikeuser;column:userID;type:int;not null"`
	User       User     `gorm:"association_foreignkey:userID;foreignkey:userID"`
}
