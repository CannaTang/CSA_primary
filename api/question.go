package api

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"Q-A/model"
	"Q-A/service"
	"Q-A/tool"
)

func questionDetail(ctx *gin.Context) {
	questionIdString := ctx.Param("question_id")
	questionId, err := strconv.Atoi(questionIdString)
	if err != nil {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, "question_id格式有误")
		return
	}

	question, err := service.GetQuestionById(questionId)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithData(ctx, "没有对应的留言")
			return
		}
		fmt.Println("get question by id err:", err)
		tool.RespInternalError(ctx)
		return
	}

	answers, err := service.GetQuestionAnswers(questionId)
	if err != nil {
		fmt.Println("get answers by question_Id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	questionDetail := model.QuestionDetail{
		Question: question,
		Answers:  answers,
	}
	tool.RespSuccessfulWithData(ctx, questionDetail)

}

func briefQuestions(ctx *gin.Context) {
	questions, err := service.GetQuestions()
	if err != nil {
		fmt.Println("get questions err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, questions)
}

func addQuestion(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)

	txt := ctx.PostForm("txt")

	question := model.Question{
		Txt:        txt,
		Username:   username,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := service.AddQuestion(question)
	if err != nil {
		fmt.Println("add question err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)

}

func deleteQuestion(ctx *gin.Context) {
	questionIdString := ctx.Param("question_id")
	questionId, err := strconv.Atoi(questionIdString)
	if err != nil {
		fmt.Println("question_id string to int err", err)
		tool.RespErrorWithData(ctx, "question_id格式有误")
		return
	}
	err = service.DeleteQuestion(questionId)
	if err != nil {
		fmt.Println("Delete question err by questionId err:", err)
		tool.RespInternalError(ctx)
	}
	tool.RespSuccessful(ctx)
}

func updateQuestion(ctx *gin.Context) {
	questionIdString := ctx.Param("question_id")
	newTxt := ctx.PostForm("new_txt")
	if newTxt == "" {
		tool.RespErrorWithData(ctx, "新问题不能为空")
		return
	}

	questionId, err := strconv.Atoi(questionIdString)
	if err != nil {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, "question_id格式有误")
		return
	}
	err = service.UpdateQuestion(questionId, newTxt)
	if err != nil {
		fmt.Println("Update question err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, "修改成功")

}

func like(ctx *gin.Context) {
	questionIdString := ctx.Query("question_id")
	questionId, err := strconv.Atoi(questionIdString)
	if err != nil {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, "question_id格式有误")
		return
	}
	err = service.Like(questionId)
	if err != nil {
		fmt.Println("AddLike to question err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, "点赞成功")

}

func cancelLike(ctx *gin.Context) {
	questionIdString := ctx.Query("question_id")
	questionId, err := strconv.Atoi(questionIdString)
	if err != nil {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, "question_id格式有误")
		return
	}
	err = service.CancelLike(questionId)
	if err != nil {
		fmt.Println("cancelLike to question err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, "已取消点赞")
}
