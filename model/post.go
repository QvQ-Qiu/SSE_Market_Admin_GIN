package model

import (
	"time"
)

// Post [...]
type Post struct {
	PostID        int       `gorm:"primary_key;column:postID"`
	UserID        int       `gorm:"index:postuser;column:userID;type:int"`
	User          User      `gorm:"association_foreignkey:userID;foreignkey:userID"`
	Partition     string    `gorm:"column:partition;type:varchar(10)"`
	Title         string    `gorm:"column:title;type:varchar(20)"`
	Ptext         string    `gorm:"column:ptext;type:varchar(5000)"`
	CommentNum    int       `gorm:"column:comment_num;type:int"`
	LikeNum       int       `gorm:"column:like_num;type:int"`
	BrowseNum     int       `gorm:"column:browse_num;type:int"`
	PostTime      time.Time `gorm:"column:post_time;type:datetime"`
	Heat          float64   `gorm:"column:heat;type:double"`
	Photos        string    `gorm:"column:photos;type:varchar(1000)"`
	Tag           string    `gorm:"column:tag;type:varchar(100)"`
	IsHighQuality bool      `gorm:"column:is_high_quality;type:bool"`
}
