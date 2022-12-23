package model

type User struct {
	Id       int
	Username string
	Password string
}

type Login struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required,lengthOk"`
	Password string `form:"password" json:"password" xml:"password" binding:"required,lengthOk"`
}
