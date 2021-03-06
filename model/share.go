package model

import "gorm.io/gorm"

type Share struct { //分享
	gorm.Model
	UserID   uint   `json:"user_id" ` //User.ID
	Username string `json:"username"`
	IconSrc  string `json:"icon_src"`
	Address  string `form:"address" json:"address" gorm:"default:'null'"` //存地址
	//Title      string `form:"title" json:"title" binding:"required" gorm:"size:30"`//标题长度不超过30
	Content     string        `form:"content" json:"content"  gorm:"type:longtext"` //附加内容
	ShareStar   uint          `json:"share_star"`
	Like        string        `json:"like" `
	ShareImages []ShareImage  `json:"share_images"`
	UserComment []UserComment `json:"user_comment"`
}
type ShareImage struct { //存share图片地址
	gorm.Model
	ShareID uint   `json:"share_id"`                                               //查图片的时候根据where("ShareID=?")来查
	Src     string `form:"src" json:"src" binding:"required" gorm:"type:longtext"` //图片地址
}
type ShareLike struct {
	gorm.Model
	//Username string	`json:"username" gorm:"size:80"`
	//Iocnsrc string	`json:"icon_src" `
	UserID  uint   `json:"user_id" binding:"required" ` //User.ID
	ShareID uint   `json:"share_id"`
	Like    string `json:"like"`
}
type UserComment struct {
	gorm.Model
	Username    string `json:"username" gorm:"size:80"`
	IconSrc     string `json:"icon_src" `
	UserID      uint   `json:"user_id" binding:"required" ` //User.ID
	ShareID     uint   `json:"share_id"`                    //查图片的时候根据where("ShareID=?")来查
	CommentStar uint   `json:"comment_star"`
	Like        string `json:"like" `
	Comment     string `json:"comment"`
	//CommentSrc []CommentSrc `json:"comment_src"`
}
type ShareCommentLike struct {
	UserID        uint   `json:"user_id" binding:"required" ` //User.ID
	UserCommentID uint   `json:"user_comment_id"`
	Like          string `json:"like"`
}

//type CommentSrc struct {
//	gorm.Model
//	CommentID uint   `json:"comment_id"` //查图片的时候根据where("ShareID=?")来查
//	Src     string `form:"src" json:"src" binding:"required" gorm:"type:longtext"` //图片地址
//}
