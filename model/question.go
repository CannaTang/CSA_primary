package model

import "time"

type Question struct {
	Id         int       `json:"id"`
	AnswerNum  int       `json:"comment_num"`
	LikeNum    int       `json:"like_num"`
	Txt        string    `json:"txt"`
	Username   string    `json:"username"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type QuestionDetail struct {
	Question
	Answers []Answer
}
