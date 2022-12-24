package api

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"Q-A/dao"
	"Q-A/model"
	"Q-A/service"
	"Q-A/tool"
)

func addAnswer(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)

	txt := ctx.PostForm("txt")
	questionIdString := ctx.PostForm("question_id")
	questionId, err := strconv.Atoi(questionIdString)
	if err != nil {
		fmt.Println("question_id string to int err", err)
		tool.RespErrorWithData(ctx, "question_id格式有误")
		return
	}

	Answer := model.Answer{
		QuestionId: questionId,
		Txt:        txt,
		Username:   username,
		CreateTime: time.Now(),
	}
	err = service.AddAnswer(Answer)
	if err != nil {
		fmt.Println("add Answer err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}

func deleteAnswer(ctx *gin.Context) {
	AnswerIdString := ctx.Param("answer_id")
	AnswerId, err := strconv.Atoi(AnswerIdString)
	if err != nil {
		fmt.Println("answer_id to string err:", err)
		tool.RespErrorWithData(ctx, "answer_id格式有误")
		return
	}

	err = service.DeleteAnswer(AnswerId)
	if err != nil {
		if err == dao.ErrAnswerNotExist {
			tool.RespErrorWithData(ctx, "评论不存在")
			return
		}
		fmt.Println("delete Answer by answer_id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, "删除成功")
}

func updateAnswer(ctx *gin.Context) {
	AnswerIdString := ctx.PostForm("answer_id")
	newTxt := ctx.PostForm("new_txt")

	AnswerId, err := strconv.Atoi(AnswerIdString)
	if err != nil {
		fmt.Println("Answer_id  string to int err:", err)
		tool.RespErrorWithData(ctx, "Answer_id格式有误")
		return
	}

	err = service.UpdateAnswer(AnswerId, newTxt)
	if err != nil {
		fmt.Println("update Answer err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, "修改评论成功")

}
