package controller

import (
	"fmt"
	"log"
	"net/http"
	"sse_market_admin/common"
	"sse_market_admin/dto"
	"sse_market_admin/model"
	"sse_market_admin/response"
	"sse_market_admin/util"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type PostMsg struct {
	UserTelephone string
	Title         string
	Content       string
	Partition     string
	Photos        string
	TagList       string
}

type User struct {
	UserID    int    `json:"-"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Name      string `json:"name"`
	Num       int    `json:"num"`
	Profile   string `json:"-"`
	Intro     string `json:"-"`
	IDpass    bool   `json:"IDpass"`
	Ban       bool   `json:"ban"`
	Punishnum int    `json:"punishnum"`
}

type Username struct {
	Name string
}

type modifyUser struct {
	Account   string
	Password1 string
	Password2 string
}

type Check struct {
	Name   string
	Phone  string
	IdPass int
}

type Key struct {
	Key         string
	Used        bool
	CreatedTime time.Time
}

type BrowseMeg struct {
	UserTelephone string
	Partition     string
	Searchinfo    string
	Tag           string
	Searchsort    string //用于分表查询 分为home,history,save三种
	Limit         int
	Offset        int
}

type AdPostResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	User   int    `json:"user_id"`
	Status string `json:"status"`
}

type AdKeyResponse struct {
	Id        int    `json:"id"`
	Code      string `json:"code"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
}
type FeedbackResponse struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
	Time string `json:"time"`
}
type PostDetailResponse struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	User_id    int    `json:"user_id"`
	LikeNum    int    `json:"like_num"`
	CommentNum int    `json:"comment_num"`
	BrowseNum  int    `json:"browse_num"`
}
type PostID struct {
	Id int `json:"post_id"`
}

type IDmsg struct {
	PostID uint
}

// 登录
func AdminLogin(ctx *gin.Context) {
	db := common.GetDB()
	var requesrAdmin model.Admin
	ctx.Bind(&requesrAdmin)

	account := requesrAdmin.Account
	password := requesrAdmin.Password
	fmt.Print(account, password)
	password = util.Decrypt(password) //加密

	var admin model.Admin
	db.Where("account = ?", account).First(&admin)
	if admin.AdminID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "管理员账号不存在")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	// if password != admin.Password {
	// 	response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "密码错误")
	// 	return
	// }
	//发放token
	token, err := common.ReleaseToken_admin(admin)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 400, nil, "系统异常")

		log.Printf("token generate error: %v", err)
		return
	}
	// //返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func AdminInfo(c *gin.Context) {
	admin, _ := c.Get("admin")
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"admin": dto.ToAdminDto(admin.(model.Admin))}})
}

// 输出所有用户
func ShowFilterUsers(ctx *gin.Context) {
	db := common.DB
	var userList []User
	var requestInfo = Check{}
	ctx.Bind(&requestInfo)

	name := requestInfo.Name
	phone := requestInfo.Phone
	idPass := requestInfo.IdPass

	if phone != "" {
		db = db.Model(&model.User{}).Where("phone = ?", phone)
	}
	if name != "" {
		db = db.Model(&model.User{}).Where("name like ?", name+"%")
	}
	if idPass != -1 {
		db = db.Model(&model.User{}).Where("idPass = ?", idPass)
	}

	db.Where("name <> ?", "用户已注销").Find(&userList)
	response.Success(ctx, gin.H{"data": userList}, "Successfully show all users")
}

// 更改是否审查
//func PassUsers(ctx *gin.Context) {
//	fmt.Println("start to pass")
//	db := common.DB
//	var username = Username{}
//	ctx.Bind(&username)
//	name := username.Name
//	if name == "" {
//		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "您未选择用户，无法审核")
//		return
//	}
//	var user model.User
//	db.Where("name = ?", name).Find(&user)
//	if user.IDpass == true {
//		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "该用户已审核通过")
//		return
//	}
//	db.Model(&model.User{}).Where("name = ?", name).Update("IDpass", true)
//	db.Where("name = ?", name).Find(&user)
//	response.Success(ctx, gin.H{"data": user}, "Successfully pass user")
//}

// 添加管理员
func AddAdmin(ctx *gin.Context) {
	fmt.Println("Start to add")
	db := common.GetDB()
	var newAdmin modifyUser
	ctx.Bind(&newAdmin)
	account := newAdmin.Account
	pass1 := newAdmin.Password1
	pass2 := newAdmin.Password2
	pass1 = util.Decrypt(pass1)
	pass2 = util.Decrypt(pass2)
	var admin model.Admin
	if account == "" {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "账号不能为空")
		return
	}

	if pass1 == "" || pass2 == "" {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "密码不能为空")
		return
	}

	if pass1 != pass2 {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "两次密码不同，请重新输入")
		return
	}

	if len(pass1) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "密码必须大于等于6位")
		return
	}

	db.Where("account = ?", account).First(&admin)
	fmt.Println(admin.Account)
	if admin.Account != "" {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "管理员账号已存在")
		return
	}

	cnt := 0
	db.Model(&model.Admin{}).Count(&cnt)
	fmt.Println(cnt)

	addAdmin := model.Admin{
		Account:  account,
		Password: pass1,
		AdminID:  cnt + 1,
	}
	db.Create(&addAdmin)

	response.Success(ctx, gin.H{"data": addAdmin}, "添加管理员成功")
}

// 修改密码
func ChangeAdminPassword(ctx *gin.Context) {
	db := common.GetDB()
	var admin modifyUser
	var newAdmin model.Admin

	ctx.Bind(&admin)
	account := admin.Account
	pass1 := admin.Password1
	pass2 := admin.Password2
	pass1 = util.Decrypt(pass1)
	pass2 = util.Decrypt(pass2)

	db.Where("Account = ?", account).First(&newAdmin)
	if newAdmin.Account == "" {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "管理员不存在")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(newAdmin.Password), []byte(pass1)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass2), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "密码加密错误")
		return
	}

	if err := db.Model(&newAdmin).Update("Password", string(hashedPassword)).Error; err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "更新密码失败")
		return
	}
	response.Success(ctx, gin.H{"data": newAdmin}, "成功修改管理员密码")
}

// 注销用户账号
func DeleteUser(ctx *gin.Context) {
	db := common.DB
	var user = Username{}
	ctx.Bind(&user)

	fmt.Println(user)
	name := user.Name
	fmt.Println(name)
	var checkUser model.User
	db.Where("name = ?", name).First(&checkUser)
	if checkUser.UserID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "未找到该用户")
		return
	}

	db.Delete(&checkUser)
	response.Response(ctx, http.StatusOK, 200, nil, "成功删除该用户")
}

// 注销管理员
func DeleteAdmin(ctx *gin.Context) {
	db := common.DB
	var user model.Admin
	ctx.Bind(&user)

	account := user.Account
	var checkUser model.Admin
	db.Where("account = ?", account).First(&checkUser)
	if checkUser.Account == "" {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "未找到该管理员")
		return
	}

	db.Where("account = ?", account).Delete(&checkUser)
	response.Success(ctx, nil, "成功删除该管理员")
}

func AdminPost(c *gin.Context) {
	db := common.GetDB()
	var requestPostMsg PostMsg
	c.Bind(&requestPostMsg)
	// 获取参数
	userTelephone := requestPostMsg.UserTelephone
	title := requestPostMsg.Title
	content := requestPostMsg.Content
	partition := requestPostMsg.Partition
	photos := requestPostMsg.Photos
	tagList := requestPostMsg.TagList
	tags := strings.Split(tagList, "|")
	tagString := strings.Join(tags, ",")
	// 验证数据
	if len(userTelephone) == 0 {
		response.Response(c, http.StatusBadRequest, 400, nil, "返回的手机号为空")
		return
	}
	if len(title) == 0 {
		response.Response(c, http.StatusBadRequest, 400, nil, "标题不能为空")
		return
	}

	if utf8.RuneCountInString(title) > 15 {
		response.Response(c, http.StatusBadRequest, 400, nil, "标题最多为15个字")
		return
	}

	if len(content) == 0 {
		response.Response(c, http.StatusBadRequest, 400, nil, "内容不能为空")
		return
	}

	if utf8.RuneCountInString(title) > 5000 {
		response.Response(c, http.StatusBadRequest, 400, nil, "内容最多为5000个字")
		return
	}

	if len(partition) == 0 {
		response.Response(c, http.StatusBadRequest, 400, nil, "分区不能为空")
		return
	}

	// if api.GetSuggestion(title) == "Block" || api.GetSuggestion(content) == "Block" {
	// 	response.Response(c, http.StatusBadRequest, 400, nil, "标题或内容含有不良信息,请重新编辑")
	// 	return
	// }

	newPost := model.Post{
		UserID:     0,
		Partition:  partition,
		Title:      title,
		Ptext:      content,
		LikeNum:    0,
		CommentNum: 0,
		BrowseNum:  0,
		Heat:       0,
		PostTime:   time.Now(),
		Photos:     photos,
		Tag:        tagString,
	}
	db.Create(&newPost)
	response.Response(c, http.StatusOK, 200, nil, "发帖成功")
}

func AdminGetPost(c *gin.Context) {
	db := common.GetDB()

	var requestBrowseMsg BrowseMeg
	c.Bind(&requestBrowseMsg)
	// limit := requestBrowseMsg.Limit
	// offset := requestBrowseMsg.Offset
	var posts []model.Post

	// db.Order("postID DESC").Offset(offset).Limit(limit).Find(&posts)
	db.Order("postID DESC").Find(&posts)

	// 转换数据格式
	var responsePosts []AdPostResponse
	for _, post := range posts {
		status := "已发布"
		if post.IsHighQuality == true {
			status = "优质贴"
		}

		responsePosts = append(responsePosts, AdPostResponse{
			ID:     post.PostID,
			Title:  post.Title,
			User:   post.UserID, // 假设 ptext 是发帖人内容
			Status: status,
		})
	}

	c.JSON(http.StatusOK, responsePosts)
}

func AdminDeletePost(c *gin.Context) {
	db := common.GetDB()
	var ID IDmsg
	c.Bind(&ID)
	PostID := ID.PostID
	var post model.Post
	db.Where("postID = ?", PostID).First(&post)
	if post.PostID == 0 {
		response.Response(c, http.StatusBadRequest, 400, nil, "帖子不存在")
		return
	}
	db.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "帖子删除成功"})
}

func AdminTopPost(c *gin.Context) {

}

func MarkHQPost(c *gin.Context) {
	db := common.GetDB()
	var ID IDmsg
	c.Bind(&ID)
	PostID := ID.PostID
	var post model.Post
	db.Where("postID = ?", PostID).First(&post)
	if post.PostID == 0 {
		response.Response(c, http.StatusBadRequest, 400, nil, "帖子不存在")
		return
	}
	db.Model(&post).Update("is_high_quality", true)
}

func RemoveHQPost(c *gin.Context) {
	db := common.GetDB()
	var ID IDmsg
	c.Bind(&ID)
	postID := ID.PostID
	var post model.Post
	db.Where("postID = ?", postID).First(&post)
	if post.PostID == 0 {
		response.Response(c, http.StatusBadRequest, 400, nil, "帖子不存在")
		return
	}
	db.Model(&post).Update("is_high_quality", false)
}

func MuteUser(c *gin.Context) {
	db := common.GetDB()
	var username = Username{}
	c.Bind(&username)
	name := username.Name
	var checkUser model.User
	db.Where("name = ?", name).First(&checkUser)
	if checkUser.UserID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 400, nil, "未找到该用户")
		return
	}
	muteUntil := time.Now().Add(12 * time.Hour)
	db.Model(&model.User{}).Where("name = ?", name).Update("Banded", muteUntil)
	response.Response(c, http.StatusOK, 200, nil, "成功禁言该用户")
}

func Release(c *gin.Context) {
	db := common.GetDB()
	var username = Username{}
	c.Bind(&username)
	name := username.Name
	var checkUser model.User
	db.Where("name = ?", name).First(&checkUser)
	if checkUser.UserID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 400, nil, "未找到该用户")
		return
	}
	db.Model(&model.User{}).Where("name = ?", name).Update("Banded", time.Now())
	response.Response(c, http.StatusOK, 200, nil, "成功解除该用户禁言")
}
func GetKey(c *gin.Context) {
	db := common.GetDB()
	//var requestBrowseMsg BrowseMeg
	//c.Bind(&requestBrowseMsg)
	// limit := requestBrowseMsg.Limit
	// offset := requestBrowseMsg.Offset
	var keys []model.CDKey
	// db.Order("cd_Keys DESC").Where("used = ?", 0).Offset(offset).Limit(limit).Find(&keys)
	db.Order("cdkeyID  DESC").Find(&keys)

	var responseKeys []AdKeyResponse
	for _, key := range keys {
		status := "未使用"
		if key.Used == true {
			status = "已使用"
		}

		responseKeys = append(responseKeys, AdKeyResponse{
			Id:        key.CDKeyID,
			Status:    status,
			CreatedAt: key.CreatedTime.String(),
			Code:      key.Content,
		})
	}
	c.JSON(http.StatusOK, responseKeys)
}
func Getfeedback(c *gin.Context) {
	db := common.GetDB()
	var feedbacks []model.Feedback
	if err := db.Order("feedbackID DESC").Find(&feedbacks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve feedbacks",
		})
		return
	}
	var response []FeedbackResponse
	for _, feedback := range feedbacks {
		response = append(response, FeedbackResponse{
			Id:   feedback.FeedbackID,
			Text: feedback.Ftext,
			Time: feedback.Time.Format("2006-01-02 15:04:05"),
		})
	}
	c.JSON(http.StatusOK, response)
}
func AddKey(c *gin.Context) {
	db := common.GetDB()
	var key Key
	c.Bind(&key)
	println(key.Key)
	newkey := model.CDKey{
		Content:     key.Key,
		Used:        false,
		CreatedTime: time.Now(),
		UsedTime:    time.Now().Add(time.Second * 100),
	}
	db.Create(&newkey)
	response.Response(c, http.StatusOK, 200, nil, "邀请码添加成功")
}
