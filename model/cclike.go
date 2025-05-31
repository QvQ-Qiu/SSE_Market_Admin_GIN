package model

// Cclike [...]
type Cclike struct {
	CclikeID   int      `gorm:"primary_key;column:cclikeID"`
	CctargetID int      `gorm:"index:ccliketarget;column:cctargetID;type:int;not null"`
	Ccomment   Ccomment `gorm:"association_foreignkey:cctargetID;foreignkey:ccommentID"`
	UserID     int      `gorm:"index:cclikeuser;column:userID;type:int;not null"`
	User       User     `gorm:"association_foreignkey:userID;foreignkey:userID"`
}
