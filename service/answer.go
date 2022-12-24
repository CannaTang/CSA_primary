package service

import (
	"Q-A/dao"
	"Q-A/model"
	"database/sql"
)

func AddAnswer(Answer model.Answer) error {
	err := dao.InsertAnswer(Answer)
	if err != nil {
		return err
	}
	Question, err := GetQuestionById(Answer.QuestionId)
	if err != nil {
		return err
	}
	err = dao.UpdateQuestionAnswerNum(Question.Id, Question.AnswerNum+1)
	return err
}

func CheckAnswerExist(AnswerId int) (bool, error) {
	flag, err := dao.CheckAnswerExist(AnswerId)
	if flag == true {
		return flag, nil
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return false, nil
}

func CheckAnswerAuthor(AnswerId int, username string) (bool, error) {
	return dao.CheckAnswerAuthor(AnswerId, username)
}

func GetQuestionAnswers(QuestionId int) ([]model.Answer, error) {
	return dao.SelectAnswerByQuestionId(QuestionId)
}

func DeleteAnswer(id int) error {

	Answer, err := dao.GetAnswerByAnswerId(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	err = dao.DeleteAnswerByAnswerId(id)
	if err != nil {
		return err
	}

	Question, err := dao.SelectQuestionById(Answer.QuestionId)
	if err != nil {
		return err
	}
	err = dao.UpdateQuestionAnswerNum(Answer.QuestionId, Question.AnswerNum-1)
	return err
}

func UpdateAnswer(AnswerId int, newTxt string) error {
	return dao.UpdateAnswerByAnswerId(AnswerId, newTxt)
}
