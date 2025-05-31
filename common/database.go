package common

import (
	"fmt"
	//_ "github.com/alexbrainman/odbc"
	"sse_market_admin/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// 使用viper从配置文件中读取数据库配置
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	// 若没有相应数据库，运行时将根据对应结构体自动创建数据库
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.Plike{})
	db.AutoMigrate(&model.Psave{})
	db.AutoMigrate(&model.Cclike{})
	db.AutoMigrate(&model.Ccdeny{})
	db.AutoMigrate(&model.Pclike{})
	db.AutoMigrate(&model.Pcdeny{})
	db.AutoMigrate(&model.Pcomment{})
	db.AutoMigrate(&model.Ccomment{})
	db.AutoMigrate(&model.Pbrowse{})
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.Feedback{})
	db.AutoMigrate(&model.Notice{})
	db.AutoMigrate(&model.Sue{})
	db.AutoMigrate(&model.CDKey{})
	db.AutoMigrate(&model.Tag{})
	db.AutoMigrate(&model.ChatMsg{})
	db.Model(&model.Pcomment{}).AddForeignKey("ptargetID", "posts(postID)", "CASCADE", "CASCADE")
	db.Model(&model.Ccomment{}).AddForeignKey("ctargetID", "pcomments(pcommentID)", "CASCADE", "CASCADE")
	db.Model(&model.Plike{}).AddForeignKey("ptargetID", "posts(postID)", "CASCADE", "CASCADE")
	db.Model(&model.Cclike{}).AddForeignKey("cctargetID", "ccomments(ccommentID)", "CASCADE", "CASCADE")
	db.Model(&model.Pclike{}).AddForeignKey("pctargetID", "pcomments(pcommentID)", "CASCADE", "CASCADE")
	db.Model(&model.Ccdeny{}).AddForeignKey("cctargetID", "ccomments(ccommentID)", "CASCADE", "CASCADE")
	db.Model(&model.Pcdeny{}).AddForeignKey("pctargetID", "pcomments(pcommentID)", "CASCADE", "CASCADE")

	//db.Model(&model.Post{}).AddForeignKey("userID", "users(userID)", "CASCADE", "CASCADE")
	//db.Model(&model.Pcomment{}).AddForeignKey("userID", "users(userID)", "CASCADE", "CASCADE")
	//db.Model(&model.Ccomment{}).AddForeignKey("userID", "users(userID)", "CASCADE", "CASCADE")
	//db.Model(&model.Plike{}).AddForeignKey("userID", "users(userID)", "CASCADE", "CASCADE")
	//db.Model(model.Cclike{}).AddForeignKey("userID", "users(userID)", "CASCADE", "CASCADE")

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
