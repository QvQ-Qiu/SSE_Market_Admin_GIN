package model

// Ccdeny [...]
type Ccdeny struct {
	CcdenyID   int      `gorm:"primary_key;column:ccdenyID"`
	CctargetID int      `gorm:"index:ccdenytarget;column:cctargetID;type:int;not null"`
	Ccomment   Ccomment `gorm:"association_foreignkey:cctargetID;foreignkey:ccommentID"`
	UserID     int      `gorm:"index:ccdenyuser;column:userID;type:int;not null"`
	User       User     `gorm:"association_foreignkey:userID;foreignkey:userID"`
}
