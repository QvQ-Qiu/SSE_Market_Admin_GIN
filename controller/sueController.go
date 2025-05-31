package controller

import (
	"fmt"
	"math"
	"net/http"
	"sse_market_admin/common"
	"sse_market_admin/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type SueResponse struct {
	SueID        int
	Targettype   string
	Targettitle  string
	Targetdetail string
	Reason       string
	Time         time.Time
}

func GetSues(c *gin.Context) {
	db := common.GetDB()
	var sueres []SueResponse
	var sues []model.Sue
	db.Find(&sues, "finish = ?", 0)
	for _, sue := range sues {
		var Targetdetail string
		var Targettitle string
		Targettitle = ""
		if sue.Targettype == "post" {
			var post model.Post
			db.Where("postID = ?", sue.TargetID).First(&post)
			if post.PostID == 0 {
				continue
			}
			Targetdetail = post.Ptext
			Targettitle = post.Title
		} else if sue.Targettype == "pcomment" {
			var pcomment model.Pcomment
			db.Where("pcommentID = ?", sue.TargetID).First(&pcomment)
			if pcomment.PcommentID == 0 {
				continue
			}
			Targetdetail = pcomment.Pctext
		} else if sue.Targettype == "ccomment" {
			var ccomment model.Ccomment
			db.Where("ccommentID = ?", sue.TargetID).First(&ccomment)
			if ccomment.CcommentID == 0 {
				continue
			}
			Targetdetail = ccomment.Cctext
		}
		suere := SueResponse{
			SueID:        sue.SueID,
			Targettype:   sue.Targettype,
			Targettitle:  Targettitle,
			Targetdetail: Targetdetail,
			Reason:       sue.Reason,
			Time:         sue.Time,
		}
		sueres = append(sueres, suere)
	}
	c.JSON(http.StatusOK, sueres)
}

type IDsue struct {
	SueID uint
}

func NoViolation(c *gin.Context) {
	db := common.GetDB()
	var ID IDsue
	c.Bind(&ID)
	SueID := ID.SueID
	var sue model.Sue
	db.Model(&sue).Where("sueID = ?", SueID).Updates(map[string]interface{}{
		"status": "nosin",
		"finish": true,
	})
	db.Where("sueID = ?", SueID).First(&sue)
	targetType := sue.Targettype
	targetID := sue.TargetID
	var sue1 model.Sue
	db.Model(&sue1).Where("targettype = ? AND targetID = ? AND finish = 0", targetType, targetID).Updates(map[string]interface{}{
		"status": "nosin",
		"finish": true,
	})
}

func Violation(c *gin.Context) {
	db := common.GetDB()
	var ID IDsue
	c.Bind(&ID)
	SueID := ID.SueID
	var sue model.Sue
	db.Model(&sue).Where("sueID = ?", SueID).Updates(map[string]interface{}{
		"status": " ok",
		"finish": true,
	})
	db.Where("sueID = ?", SueID).First(&sue)
	targetType := sue.Targettype
	targetID := sue.TargetID
	var sue1 model.Sue
	db.Model(&sue1).Where("targettype = ? AND targetID = ? AND finish = 0", targetType, targetID).Updates(map[string]interface{}{
		"status": " ok",
		"finish": true,
	})
	var targetuser model.User
	var content string
	var suetype string
	// fmt.Println(sue.Targettype)
	// fmt.Println(sue.TargetID)
	if sue.Targettype == "post" {
		suetype = "帖子"
		var post model.Post
		db.Where("postID = ?", sue.TargetID).First(&post)
		if post.PostID == 0 {
			return
		}
		if len(post.Ptext) <= 30 {
			content = post.Ptext
		} else {
			content = string([]rune(post.Ptext)[:30])
		}
		db.Where("userID = ?", post.UserID).First(&targetuser)
		if targetuser.UserID == 0 {
			return
		}
		db.Delete(&post)
	} else if sue.Targettype == "pcomment" {
		suetype = "评论"
		var pcomment model.Pcomment
		db.Where("pcommentID = ?", sue.TargetID).First(&pcomment)
		if pcomment.PcommentID == 0 {
			return
		}
		if len(pcomment.Pctext) <= 30 {
			content = pcomment.Pctext
		} else {
			content = string([]rune(pcomment.Pctext)[:30])
		}
		var post model.Post
		db.Where("userID = ?", pcomment.UserID).First(&targetuser)
		if targetuser.UserID == 0 {
			return
		}
		// 帖子评论数减相应数字
		db.Where("postID = ?", pcomment.PtargetID).First(&post)
		if post.PostID == 0 {
			return
		}
		var ccomment model.Ccomment
		var count int64
		db.Model(&ccomment).Where("ctargetID = ?", pcomment.PcommentID).Count(&count)
		db.Model(&post).UpdateColumn("comment_num", gorm.Expr("comment_num - ?", count+1))
		// 剪掉相应的热度
		currentTime := time.Now()
		timedif := currentTime.Sub(pcomment.Time)
		hours := timedif.Hours()
		days := int(hours / 24)
		weightComment := float64(6)
		db.Where("postID = ?", pcomment.PtargetID).First(&post)
		if days > 0 {
			weightCommentPower := math.Pow(0.5, float64(days))
			deleteHeat := math.Pow(weightComment, weightCommentPower)
			db.Model(&post).Update("heat", post.Heat-(deleteHeat+float64(count)))
		} else {
			deleteCcommentHeat := float64(count * int64(weightComment))
			db.Model(&post).Update("heat", post.Heat-(weightComment+deleteCcommentHeat))
		}
		//
		db.Delete(&pcomment)
	} else if sue.Targettype == "ccomment" {
		suetype = "评论"
		var ccomment model.Ccomment
		db.Where("ccommentID = ?", sue.TargetID).First(&ccomment)
		if ccomment.CcommentID == 0 {
			return
		}
		if len(ccomment.Cctext) <= 30 {
			content = ccomment.Cctext
		} else {
			content = string([]rune(ccomment.Cctext)[:30])
		}
		// 剪掉相应的热度
		currentTime := time.Now()
		timedif := currentTime.Sub(ccomment.Time)
		hours := timedif.Hours()
		days := int(hours / 24)
		weightComment := float64(6)
		var targetcommentid model.Pcomment
		db.Where("pcommentID= ?", ccomment.CtargetID).First(&targetcommentid)
		var post model.Post
		db.Where("postID = ?", targetcommentid.PtargetID).First(&post)
		if post.PostID == 0 {
			return
		}
		if days > 0 {
			weightCommentPower := math.Pow(0.5, float64(days))
			deleteHeat := math.Pow(weightComment, weightCommentPower)
			db.Model(&post).Update("heat", post.Heat-deleteHeat)
		} else {
			db.Model(&post).Update("heat", post.Heat-weightComment)
		}
		//
		db.Where("userID = ?", ccomment.UserID).First(&targetuser)
		if targetuser.UserID == 0 {
			return
		}
		// 帖子评论数减一
		db.Where("pcommentID= ?", ccomment.CtargetID).First(&targetcommentid)
		db.Where("postID = ?", targetcommentid.PtargetID).First(&post)
		db.Model(&post).UpdateColumn("comment_num", gorm.Expr("comment_num - ?", 1))
		db.Delete(&ccomment)
	}

	targetuser.Punishnum += 1
	targetuser.Banend = time.Now().AddDate(0, 0, targetuser.Punishnum)
	db.Model(&targetuser).Updates(map[string]interface{}{
		"punishnum": targetuser.Punishnum,
		"banend":    targetuser.Banend,
	})
	banEndTime := targetuser.Banend.Format("2006-01-02 15:04:05")
	noticetext := fmt.Sprintf("你的%s[%s]内容涉嫌违规被举报，相关内容已被删除，你的账号被封禁，封禁时间到%s。",
		suetype, content, banEndTime)
	// fmt.Println(noticetext)
	var tonotice model.Notice
	tonotice.Receiver = targetuser.UserID
	tonotice.Type = "punish"
	tonotice.Ntext = noticetext
	// tonotice.Ntext = strings.Join([]string{"你的", suetype, "【", content, "】内容涉嫌违规被举报，相关内容已被删除，你的账号被封禁，封禁时间到", banEndTime, "。"}, " ")
	// tonotice.Ntext = "你的 " + suetype + " 【" + content + "】 内容涉嫌违规被举报，相关内容已被删除，你的账号被封禁，封禁时间到 " + banEndTime + "。"
	tonotice.Time = time.Now()
	tonotice.Read = false
	tonotice.Target = sue.SueID
	db.Create(&tonotice)

	var fromnotice model.Notice
	fromnotice.Receiver = sue.UserID
	fromnotice.Type = "sue"
	fromnotice.Ntext = "我们已经收到你的举报，违规内容已被删除，感谢你的支持"
	fromnotice.Time = time.Now()
	fromnotice.Read = false
	fromnotice.Target = sue.SueID
	db.Create(&fromnotice)
}
