package dao

import (
	"Q-A/model"
)

func InsertQuestion(Question model.Question) error {
	_, err := dB.Exec("INSERT INTO question(username,txt,create_time,update_time) values(?,?,?,?);", Question.Username, Question.Txt, Question.CreateTime, Question.UpdateTime)
	return err
}

func SelectQuestionById(QuestionId int) (model.Question, error) {
	Question := model.Question{}

	row := dB.QueryRow("SELECT id,username,answer_num,txt,create_time,update_time FROM question WHERE id = ?", QuestionId)

	if row.Err() != nil {
		return model.Question{}, row.Err()
	}

	err := row.Scan(&Question.Id, &Question.Username, &Question.AnswerNum, &Question.Txt, &Question.CreateTime, &Question.UpdateTime)
	if err != nil {
		return model.Question{}, err
	}
	return Question, nil
}

func SelectQuestions() ([]model.Question, error) {
	var Questions []model.Question
	rows, err := dB.Query("SELECT id,username,txt,create_time,update_time,answer_num FROM question")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Question model.Question
		err = rows.Scan(&Question.Id, &Question.Username, &Question.Txt, &Question.UpdateTime, &Question.UpdateTime, &Question.AnswerNum)
		if err != nil {
			return nil, err
		}
		Questions = append(Questions, Question)
	}
	return Questions, nil

}

func DeleteQuestion(QuestionId int) error {
	_, err := dB.Exec("DELETE FROM question WHERE id = ?", QuestionId)
	return err
}

func UpdateQuestionTxt(QuestionId int, newTxt string) error {
	_, err := dB.Exec("UPDATE question SET txt = ? WHERE id = ?", newTxt, QuestionId)
	return err
}

func UpdateQuestionAnswerNum(QuestionId, newAnswerNum int) error {
	_, err := dB.Exec("UPDATE question SET answer_num = ? WHERE id = ?", newAnswerNum, QuestionId)
	return err
}

func AddLike(QuestionId int) error {
	_, err := dB.Exec("UPDATE question SET like_num = like_num + 1 WHERE id = ?", QuestionId)
	return err
}

func DeleteLike(QuestionId int) error {
	_, err := dB.Exec("UPDATE question SET like_num = like_num - 1 WHERE id = ?", QuestionId)
	return err
}
