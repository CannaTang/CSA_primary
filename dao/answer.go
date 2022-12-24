package dao

import (
	"Q-A/model"
	"errors"
)

var ErrAnswerNotExist = errors.New("回答不存在")

func InsertAnswer(Answer model.Answer) error {
	_, err := dB.Exec("INSERT INTO answer(question_id,username,create_time,txt) VALUES(?,?,?,?)", Answer.QuestionId, Answer.Username, Answer.CreateTime, Answer.Txt)
	if err != nil {
		return err
	}
	return nil
}

func SelectAnswerByQuestionId(QuestionId int) ([]model.Answer, error) {
	var Answers []model.Answer

	rows, err := dB.Query("SELECT id,Question_id,username,txt,create_time FROM answer WHERE question_id = ?", QuestionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Answer model.Answer
		err = rows.Scan(&Answer.Id, &Answer.QuestionId, &Answer.Username, &Answer.Txt, &Answer.CreateTime)
		if err != nil {
			return nil, err
		}
		Answers = append(Answers, Answer)
	}
	return Answers, nil
}

func DeleteAnswerByAnswerId(AnswerId int) error {
	ret, err := dB.Exec("DELETE FROM answer WHERE id = ?", AnswerId)
	if err != nil {
		return err
	}
	n, err := ret.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrAnswerNotExist
	}
	return nil
}

func DeleteAnswersByQuestionId(QuestionId int) error {
	_, err := dB.Exec("DELETE FROM answer WHERE Question_id = ?", QuestionId)
	return err
}

func UpdateAnswerByAnswerId(AnswerId int, newTxt string) error {
	_, err := dB.Exec("UPDATE answer SET txt = ? WHERE id = ?", newTxt, AnswerId)
	return err
}

func GetAnswerByAnswerId(AnswerId int) (model.Answer, error) {
	var Answer model.Answer
	row := dB.QueryRow("SELECT id,question_id,username,txt,create_time FROM answer WHERE id = ?", AnswerId)

	if row.Err() != nil {
		return Answer, row.Err()
	}
	err := row.Scan(&Answer.Id, &Answer.QuestionId, &Answer.Username, &Answer.Txt, &Answer.CreateTime)
	if err != nil {
		return Answer, err
	}
	return Answer, nil

}
