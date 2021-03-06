package service

import (
	"SchoolCat/model"
	"SchoolCat/util/responser"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func NewCard(c *gin.Context) {
	var card model.CatCard
	err := c.ShouldBind(&card)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(card)
	err = DB.Create(&card).Error
	if err != nil {
		log.Println(err)
	}
	responser.CardSucceed(c, card.ID)
}

func DeleteCard(c *gin.Context) {

	var card model.CatCard
	cardId := c.Query("card_id")

	var user model.User
	uid, _ := strconv.Atoi(c.Query("user_id")) //执行删除操作者的id
	res := DB.Where("id = ?", uid).Take(&user)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}

	if AdminExist(user.Email) {
		responser.UserIdWrong(c)
	} else {
		res = DB.Where("id = ?", cardId).Take(&card)
		if res.Error != nil {
			fmt.Println(res.Error)
			return
		}
		DB.Delete(&card)
		responser.DeleteSucceed(c)
	}
}

func ViewCard(c *gin.Context) {

	var cards []model.CatCard
	res := DB.Find(&cards).Limit(10).Offset(10)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	for i := 0; i < len(cards); i++ {
		var cardImg []model.CatCardSrc
		var comment []model.CatCardComment
		res = DB.Where("card_id=?", cards[i].ID).Find(&cardImg)
		if res.Error != nil {
			fmt.Println(res.Error)
			return
		}
		res = DB.Where("card_id=?", cards[i].ID).Find(&comment)
		if res.Error != nil {
			fmt.Println(res.Error)
			return
		}
		cards[i].CatCardSrc = cardImg
		cards[i].CatCardComment = comment
	}
	//fmt.Println(cards)
	responser.DisplayCards(c, cards)
}

func NewCardComment(c *gin.Context) {
	var comment model.CatCardComment
	err := c.ShouldBind(&comment)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = DB.Create(&comment).Error
	if err != nil {
		log.Println(err)
	}

	responser.CommentSucceed(c, comment.ID)
}
func DeleteCardComment(c *gin.Context) {
	var comment model.CatCardComment
	uid, _ := strconv.Atoi(c.Query("user_id"))
	commentId := c.Query("comment_id")

	res := DB.Where("id = ?", commentId).Take(&comment)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	if comment.UserID != uint(uid) {
		responser.UserIdWrong(c)
	} else {
		DB.Delete(&comment)
		responser.DeleteSucceed(c)
	}
}
