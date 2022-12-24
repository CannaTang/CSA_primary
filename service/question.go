package service

import (
	"Q-A/dao"
	"Q-A/model"
	"database/sql"
)

func AddQuestion(Question model.Question) error {
	err := dao.InsertQuestion(Question)
	return err
}

func GetQuestions() ([]model.Question, error) {
	return dao.SelectQuestions()
}

func GetQuestionById(QuestionId int) (model.Question, error) {
	return dao.SelectQuestionById(QuestionId)
}

func DeleteQuestion(QuestionId int) error {
	err := dao.DeleteQuestion(QuestionId)
	if err != nil {
		return err
	}
	err = dao.DeleteAnswersByQuestionId(QuestionId)
	if err != nil {
		return err
	}
	return nil
}

func UpdateQuestion(QuestionId int, newTxt string) error {
	return dao.UpdateQuestionTxt(QuestionId, newTxt)
}

func Like(QuestionId int) error {
	return dao.AddLike(QuestionId)
}

func CancelLike(QuestionId int) error {
	return dao.DeleteLike(QuestionId)
}

func CheckQuestionExist(QuestionId int) (bool, error) {
	flag, err := dao.CheckQuestionExist(QuestionId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return flag, nil
}

func CheckQuestionAuthor(QuestionId int, username string) (bool, error) {
	return dao.CheckQuestionAuthor(QuestionId, username)
}
