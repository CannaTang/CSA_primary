package model

import "time"

type Answer struct {
	Id         int
	QuestionId int
	Txt        string
	Username   string
	CreateTime time.Time
}
