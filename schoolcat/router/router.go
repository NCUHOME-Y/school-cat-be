package router

import (
	"SchoolCat/config"
	"SchoolCat/handler"
	"SchoolCat/midware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Router()   {
	gin.SetMode(config.AppMode)

	engine := gin.Default()
	engine.Use(midware.CORS())

	engine.POST("/login",handler.Login)//登录
	engine.POST("/register",handler.Register)//注册

	admin := engine.Group("/admin",midware.Admin(),midware.JWT())
	user :=	engine.Group("/user",midware.JWT())

	//token := engine.Group()

	//百科相关
		admin.POST("/tip",handler.Tip)//add tip
		admin.DELETE("/tip",handler.DeleteTip)//delete tip

	user.POST("/newTipComment",handler.NewTipComment)//用户添加评论
	user.DELETE("/deleteTipComment",handler.DeleteTipComment)
	user.GET("/newTip",handler.ViewTip)
	user.PUT("/tipCommentLike",handler.TipCommentLike)//用户点赞


	//资料卡片相关
	admin.POST("/newCard",handler.NewCard)//add card
	admin.DELETE("/deleteCard",handler.DeleteCard)//delete card

	user.POST("/newCArdComment",handler.NewCardComment)//用户添加评论
	user.DELETE("/deleteCardComment",handler.DeleteCardComment)
	user.GET("newCard",handler.ViewCard)


	//用户相关
		 user.POST("/info",handler.Info)//添加更改信息：邮箱，性别，学校，昵称，简介


	//share相关
		 user.GET("/selfShare",handler.SelfShare)//用户界面刷新自己的分享
		 user.GET("/viewShare",handler.ViewShare)//用户请求share

		 user.POST("newShare",handler.NewShare)//用户添加share
		 user.DELETE("/deleteShare",handler.DeleteShare)
		 user.POST("/newShareComment",handler.NewShareComment)//用户添加评论
		 user.DELETE("/deleteShareComment",handler.DeleteShareComment)

		 user.GET("/search",handler.Search)//用户搜索

		 user.PUT("/shareCommentLike",handler.ShareCommentLike)//用户点赞
		 user.PUT("/shareLike",handler.ShareLike)//用户点赞

	err := engine.Run(config.HttpPort)
	if err != nil {
		fmt.Println("路由运行端口出错",err)
	}
}
