# 软件中级实训数据库设计

2023.5.21

## 帖子 post

- **帖子ID（主键）**postID int
- 用户ID（用户的外键） userID int
- 所属主题  partition varchar(10)
- 标题 title varchar(20)
- 正文内容 ptext varchar(5000) 
- 评论数量 comment_num int
- 点赞数量 like_num int
- 发帖时间 post_time datetime
- 热度 heat double
- 文件 photos varchar(1000)

## 帖子评论 pcomment

- **帖子评论ID（主键）** pcommentID int 主键
- 用户ID（用户的外键）userID int
- 评论目标ID（帖子的外键）ptargetID int
- 点赞数量 like_num int
- 评论内容 pctext varchar(1000)
- 评论时间 time datetime

## 评论的评论 ccomment

- **评论ID（主键）** ccommentID int
- 用户ID（用户的外键）userID int
- 评论目标ID（评论的外键）ctargetID int
- 点赞数量 like_num bigint
- 评论内容 cctext varchar(100)
- 评论时间 time datetime
- 回复用户ID usertargetID int

## 帖子点赞 plike

- **点赞ID**  plikeID bigint
- 点赞人ID（用户的外键）userID int
- 点赞目标ID（帖子的外键）ptargetID int


## 帖子的评论点赞 pclike

- **点赞ID**  pclikeID int
- 点赞人ID（用户的外键）userID int
- 点赞目标ID（帖子的外键）pctargetID int

## 评论的评论点赞 cclike

- **点赞ID**  cclikeID int
- 点赞人ID（用户的外键）userID int
- 点赞目标ID（评论的外键）cctargetID int

## 用户 user

- **用户ID** userID int
- 手机号 phone char(15)
- 邮箱号 email varchar(255)
- 密码 password varchar(255) 
- 用户名称（昵称) name varchar(50)
- 学号/工号 num int
- 头像 profile  varchar(50)
- 简介 intro varchar(255)
- 是否通过身份验证 idpass boolean /tinyint(1)
- 解封禁时间 banend datetime
- 被惩罚次数 punishnum int 

## 管理员 admin

- **账号** account varchar(100)
- 密码 password varchar(100)

## 举报 sue

- **举报ID** sueID int
- 举报目标类型 targettype enum
  - 帖子
  - 帖子的评论
  - 评论的评论

- 举报目标ID  ptargetID int
- 用户ID（举报人ID，用户的外键）userID int
- 举报原因 reason varchar(1000)
- 举报时间 sue_time datetime
- 举报处理情况 status varchar(20)/enum
  - ok 已处理
  - nosin 经检查无违规
  - wait 受理中
- 是否处理 finish boolean tinyint(1)



## 反馈 feedback

- **反馈ID** feedbackID int
- 用户ID（用户的外键）userID int
- 反馈内容 ftext varchar(1000)
- 反馈时间 time timestamp
- 反馈处理情况 status  enum
  - wait 受理中
  - ok 已处理


## 通知 notice

- **通知ID** noticeID int
- 接受者ID（用户的外键）receiver int
- 发送者ID（用户的外键）sender int 
- 通知类型 type enum
  - 帖子被评论 pcomment
  - 评论被评论 ccomment
  - 被惩罚 punish
  - 反馈 feedback
  - 举报回复 sue
- 通知内容 ntext varchar(1000）
- 通知时间 time datetime
- 通知目标 target int
- 是否已读 read tinyint(1)

## 收藏 psave

- **点赞ID**  psaveID int
- 点赞人ID（用户的外键）userID int
- 点赞目标ID（帖子的外键）ptargetID int

## 邀请码 CDKey

- 邀请码ID cdkeyID int
- 邀请码号 content char(6)
- 是否使用 used tinyint(1)
- 创建时间 createdtime datetime
- 使用时间 usedtime datetime

