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

/*
所有函数均需Cookie内username参数
*/

/*
获取路由参数question_id
判断:question_d格式正确, 对应question_id存在, (留言存在)
*/
func questionDetail(ctx *gin.Context) {
	questionIdString := ctx.Param("question_id")
	questionId, err := strconv.Atoi(questionIdString)
	if err != nil {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, "question_id格式有误")
		return
	}
	flag, err := service.CheckQuestionExist(questionId)
	if err != nil {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, err)
		return
	}
	if flag == false {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, "对应问题不存在")
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

/*
传出所有问题
*/
func briefQuestions(ctx *gin.Context) {
	questions, err := service.GetQuestions()
	if err != nil {
		fmt.Println("get questions err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, questions)
}

/*
Cookie获取:username
传入字段:txt
*/
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

/*
获取Cookie:username
传入字段:question_id
判断:question_id格式正确, question_id存在, 是该问题的创建者
*/
func deleteQuestion(ctx *gin.Context) {
	questionIdString := ctx.Param("question_id")
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	questionId, err := strconv.Atoi(questionIdString)
	if err != nil {
		fmt.Println("question_id string to int err", err)
		tool.RespErrorWithData(ctx, "question_id格式有误")
		return
	}
	flag, err := service.CheckQuestionExist(questionId)
	if err != nil {
		fmt.Println("question_id string to int err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if flag == false {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, "对应问题不存在")
		return
	}
	flag, err = service.CheckQuestionAuthor(questionId, username)
	if err != nil {
		fmt.Println("question_id string to int err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if flag == false {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, "您不是该问题的创建者")
		return
	}
	err = service.DeleteQuestion(questionId)
	if err != nil {
		fmt.Println("Delete question err by questionId err:", err)
		tool.RespInternalError(ctx)
	}
	tool.RespSuccessful(ctx)
}

/*
获取Cookies:username
传入字段:question_id, new_txt
判断:question_id格式正确, question_id存在, 是该问题创建者, nex_txt不为空
*/
func updateQuestion(ctx *gin.Context) {
	questionIdString := ctx.Param("question_id")
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	newTxt := ctx.PostForm("new_txt")
	questionId, err := strconv.Atoi(questionIdString)
	if err != nil {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, "question_id格式有误")
		return
	}
	flag, err := service.CheckQuestionExist(questionId)
	if err != nil {
		fmt.Println("judge QuestionExist err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "问题不存在")
		return
	}
	flag, err = service.CheckQuestionAuthor(questionId, username)
	if err != nil {
		fmt.Println("question_id string to int err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if flag == false {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, "您不是该问题的创建者")
		return
	}
	if newTxt == "" {
		tool.RespErrorWithData(ctx, "新问题不能为空")
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

/*
获取路由参数:question_id
判断:question_id格式正确, 对应question_id存在
*/
func like(ctx *gin.Context) {
	questionIdString := ctx.Query("question_id")
	questionId, err := strconv.Atoi(questionIdString)
	if err != nil {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, "question_id格式有误")
		return
	}
	flag, err := service.CheckQuestionExist(questionId)
	if err != nil {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, err)
		return
	}
	if flag == false {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, "对应问题不存在")
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

/*
获取路由参数:question_id
判断:question_id正确, 判断对应question_id存在
*/
func cancelLike(ctx *gin.Context) {
	questionIdString := ctx.Query("question_id")
	questionId, err := strconv.Atoi(questionIdString)
	if err != nil {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, "question_id格式有误")
		return
	}
	flag, err := service.CheckQuestionExist(questionId)
	if err != nil {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, err)
		return
	}
	if flag == false {
		fmt.Println("question_id string to int err:", err)
		tool.RespErrorWithData(ctx, "对应问题不存在")
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
