package model

import "gorm.io/gorm"

type CatCard struct {
	gorm.Model
	CreateId       uint             `json:"create_id"`
	DeleteID       uint             `json:"delete_id"`
	CatName        string           `json:"cat_name"`
	FurColor       string           `json:"fur_color"`
	Health         string           `json:"health"`
	Sterilization  string           `json:"sterilization"` //绝育情况
	Area           string           `json:"area"`          //出没范围·
	Appearance     string           `json:"appearance"`
	Time           string           `json:"time"`         //第一次目击时间
	Relationship   string           `json:"relationship"` //猫际关系
	CatCardComment []CatCardComment `json:"cat_card_comment"`
	CatCardSrc     []CatCardSrc     `json:"cat_card_src"`
}
type CatCardSrc struct {
	gorm.Model
	CatCardID uint   `json:"cat_card_id"`
	Src       string `form:"src" json:"src" binding:"required" gorm:"type:longtext"` //图片地址
}
type CatCardComment struct {
	gorm.Model
	Username    string `json:"username" gorm:"size:80"`
	IconSrc     string `json:"icon_src" `
	UserID      uint   `json:"user_id" binding:"required" ` //User.ID
	CatCardID   uint   `json:"cat_card_id"`
	CommentStar uint   `json:"comment_star"`
	Comment     string `json:"comment"`
}
