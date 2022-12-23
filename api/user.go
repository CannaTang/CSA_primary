package api

import (
	"Q-A/dao"
	"fmt"

	"github.com/gin-gonic/gin"

	"Q-A/model"
	"Q-A/service"
	"Q-A/tool"
)

/*
传入字段:username, old_password, new_password
判断:用户名存在, 旧密码正确
*/
func changePassword(ctx *gin.Context) {
	oldPassword := ctx.PostForm("old_password")
	newPassword := ctx.PostForm("new_password")

	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)

	//检查用户名是否存在
	flag, err := service.IsRepeatUsername(username)
	if err != nil {
		fmt.Println("judge repeat username err:", err)
		tool.RespInternalError(ctx)
		return
	}

	if !flag {
		tool.RespErrorWithData(ctx, "用户名不存在")
		return
	}

	//检验旧密码是否正确

	if flag, err := service.IsPasswordCorrect(username, oldPassword); err != nil {
		fmt.Println("judge password correct err:", err)
		tool.RespInternalError(ctx)
		return
	} else if !flag {
		tool.RespErrorWithData(ctx, "旧密码错误")
		return
	}

	//修改新密码
	err = service.ChangePassword(username, newPassword)
	if err != nil {
		tool.RespInternalError(ctx)
		fmt.Println("change password err:", err)
		return
	}
	tool.RespSuccessful(ctx)
}

/*
传入字段:username, password
判断:用户名存在， 密码正确
*/
func login(ctx *gin.Context) {
	var u model.Login
	if err := ctx.ShouldBind(&u); err != nil {
		tool.RespErrorWithData(ctx, err.Error())
		return
	}

	flag, err := service.IsRepeatUsername(u.Username)
	if err != nil {
		fmt.Println("judge repeat username err:", err)
		tool.RespInternalError(ctx)
		return
	}

	if !flag {
		tool.RespErrorWithData(ctx, "用户名不存在")
		return
	}

	flag, err = service.IsPasswordCorrect(u.Username, u.Password)
	if err != nil {
		fmt.Println("judge password correct err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "密码错误")
		return
	}
	ctx.SetCookie("username", u.Username, 700, "/", "", false, false)
	tool.RespSuccessful(ctx)
}

/*
传入字段:username, password
判断:用户名不存在，
*/
func register(ctx *gin.Context) {
	var u model.Login
	if err := ctx.ShouldBind(&u); err != nil {
		tool.RespErrorWithData(ctx, err.Error())
		return
	}
	user := model.User{
		Username: u.Username,
		Password: u.Password,
	}
	flag, err := service.IsRepeatUsername(u.Username)
	if err != nil {
		fmt.Println("judge repeat username err:", err)
		tool.RespInternalError(ctx)
		return
	}

	if flag {
		tool.RespErrorWithData(ctx, "用户名已存在")
		return
	}
	err = service.Register(user)
	if err != nil {
		fmt.Println("register err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)

}

/*
传入字段:username, password
判断:用户存在， 密码正确
*/
func delUser(ctx *gin.Context) {
	var u model.Login
	if err := ctx.ShouldBind(&u); err != nil {
		tool.RespErrorWithData(ctx, err.Error())
		return
	}

	flag, err := service.IsRepeatUsername(u.Username)
	if err != nil {
		fmt.Println("judge repeat username err:", err)
		tool.RespInternalError(ctx)
		return
	}

	if !flag {
		tool.RespErrorWithData(ctx, "用户名不存在")
		return
	}

	flag, err = service.IsPasswordCorrect(u.Username, u.Password)
	if err != nil {
		fmt.Println("judge password correct err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "密码错误")
		return
	}

	err = dao.DelUser(u.Username)
	if err != nil {
		tool.RespErrorWithData(ctx, "删除失败")
		return
	}

	ctx.SetCookie("username", u.Username, 700, "/", "", false, false)
	tool.RespSuccessfulWithData(ctx, "删除成功")
}
