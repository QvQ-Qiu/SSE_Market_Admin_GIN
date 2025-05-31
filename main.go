package main

import (
	"fmt"
	"log"
	"sse_market_admin/common"
	"sse_market_admin/config"
	"sse_market_admin/middleware"
	"sse_market_admin/route"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/spf13/viper"

	"net/http"
	"os/exec"
	"time"
)

func Copy() {
	// 数据库连接信息

	dbPort := viper.GetInt("datasource.port")
	dbUser := viper.GetString("datasource.username")
	dbPassword := viper.GetString("datasource.password")
	dbName := viper.GetString("datasource.database")

	// 备份目录
	backupDir := "/app/database"

	c := cron.New()
	c.AddFunc("@every 12h", func() {
		backupFile := fmt.Sprintf("%s/backup_%s.sql", backupDir, time.Now().Format("2006-01-02 15:04:05"))
		cmd := exec.Command("mysqldump", fmt.Sprintf("-P%d", dbPort), fmt.Sprintf("-u%s", dbUser), fmt.Sprintf("-p%s", dbPassword), dbName, "--result-file="+backupFile)
		err := cmd.Run()
		if err != nil {
			log.Println("备份失败:", err)
			return
		}
		log.Println("备份成功:", backupFile)
	})
	c.Start()
}

var r *gin.Engine

func main() {
	config.InitConfig()
	go Copy()
	db := common.InitDB()
	common.InitJWTkey()
	defer db.Close()
	r = gin.Default()
	r.Use(middleware.LoggerToFile())
	// 使用 http.FileServer 文件服务器处理 "/uploads/" 开头的请求，
	// 文件服务器获取文件的位置在 "./public" 文件夹下。

	r.StaticFS("/uploads", http.Dir("./public/uploads"))
	r.StaticFS("/resized", http.Dir("./public/resized"))

	route.CollectRoute(r)
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Printf("Server started on port 8080")
	select {}
}
