package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"Q-A/tool"
)

func InitEngine() {
	engine := gin.Default()

	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("lengthOk", tool.LengthOk)
		if err != nil {
			fmt.Println("验证器注册失败")
			return
		}
	}

	engine.POST("/register", register) //注册
	engine.POST("/login", login)       //登陆
	engine.DELETE("/delete", delUser)  //删除

	userGroup := engine.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.POST("/password", changePassword)
	}
	postGroup := engine.Group("/question")
	{
		postGroup.Use(auth)
		postGroup.POST("/", addQuestion)
		postGroup.POST("/:question_id", updateQuestion)
		postGroup.DELETE("/:question_id", deleteQuestion)

		postGroup.GET("/", briefQuestions)
		postGroup.GET("/:question_id", questionDetail)

		postGroup.POST("/like", like)
		postGroup.DELETE("/like", cancelLike)
	}

	commentGroup := engine.Group("/answer")
	{
		commentGroup.Use(auth)
		commentGroup.POST("/", addAnswer)
		commentGroup.DELETE("/:answer_id", deleteAnswer)
		commentGroup.POST("/update", updateAnswer)
	}
	engine.Run()
}
