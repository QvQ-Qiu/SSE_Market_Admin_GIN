package model

// Tag [...]
type Tag struct {
	TagID int    `gorm:"primary_key;column:tagID"`
	Name  string `gorm:"column:name;unique;type:varchar(10);not null"`
	Value string `gorm:"column:value;type:varchar(100)"`
	Type  string `gorm:"column:type;type:enum('job','course')"`
	Num   int    `gorm:"column:num;type:integer"`
}
