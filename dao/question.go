package dao

import (
	"Q-A/model"
	"fmt"
	"strconv"
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

func CheckQuestionExist(QuestionId int) (bool, error) {
	totalRow, err := dB.Query("SELECT COUNT(*) FROM question WHERE id = ?", QuestionId)
	if err != nil {
		fmt.Println("GetKnowledgePointListTotal error", err)
		return false, nil
	}
	defer totalRow.Close()
	total := 0
	for totalRow.Next() {
		err := totalRow.Scan(
			&total,
		)
		if err != nil {
			fmt.Println("GetKnowledgePointListTotal error", err)
			continue
		}
	}
	println("total:", strconv.Itoa(total))
	if total > 0 {
		return true, err
	}
	return false, err
}

func CheckQuestionAuthor(QuestionId int, username string) (bool, error) {
	Question := model.Question{}

	row := dB.QueryRow("SELECT id,username,answer_num,txt,create_time,update_time FROM question WHERE id = ?", QuestionId)

	if row.Err() != nil {
		return false, row.Err()
	}

	err := row.Scan(&Question.Id, &Question.Username, &Question.AnswerNum, &Question.Txt, &Question.CreateTime, &Question.UpdateTime)
	if err != nil {
		return false, err
	}
	if Question.Username == username {
		return true, err
	}
	return false, nil
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
