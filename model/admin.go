package model

// Admin [...]
type Admin struct {
	AdminID    int  `gorm:"primary_key;column:adminID"`
	Account  string `gorm:"primary_key;column:account;type:varchar(100);not null"`
	Password string `gorm:"column:password;type:varchar(20)"`
}
