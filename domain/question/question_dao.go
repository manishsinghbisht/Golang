package question

import (
	"errors"
	"fmt"
	"github.com/manishsinghbisht/greenleaf-api/datasources"
	"github.com/manishsinghbisht/utils-go/logger"
	"github.com/manishsinghbisht/utils-go/rest_errors"
)

const (
	queryFindById = "select q.id, q.surveyid, q.question, a.id as answerid, a.answer from  questionmaster q, answermaster a" +
		"where q.Id = a.questionid" +
		"and q.isactive = 1" +
		"and a.isactive = 1" +
		"and q.id = 'fb813845-8c84-4989-85c4-ea55952fbc1b'" +
		"order by q.surveyid, q.sequence , a.sequence"

	queryGetQuestionsBySurveyId = "select q.id, q.surveyid, q.question, a.id as answerid, a.answer from  questionmaster q, answermaster a" +
		"where q.Id = a.questionid" +
		"and q.isactive = 1" +
		"and a.isactive = 1" +
		"and q.surveyid = '428e8918-cc73-4593-becf-8252a7c263e9'" +
		"order by q.surveyid, q.sequence , a.sequence"
)

func (question *Question) Get() rest_errors.RestErr {
	stmt, err := mysql_db.Client.Prepare(queryFindById)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return rest_errors.NewInternalServerError("error when trying to get questions", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(question.Id)

	if getErr := result.Scan(&question.Id, &question.surveyid, &question.sequence); getErr != nil {
		logger.Error("error when trying to get question by Id", getErr)
		return rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}

	return nil
}

func (question *Question) FindBySurvey(surveyid string) ([]Question, rest_errors.RestErr) {
	stmt, err := mysql_db.Client.Prepare(queryGetQuestionsBySurveyId)
	if err != nil {
		logger.Error("error when trying to prepare find questions by survey statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get questions by surveyid", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(surveyid)
	if err != nil {
		logger.Error("error when trying to find questions by surveyid", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get questions by surveyid", errors.New("database error"))
	}
	defer rows.Close()

	results := make([]Question, 0)
	for rows.Next() {
		var question Question
		if err := rows.Scan(&question.Id, &question.surveyid, &question.question, &question.sequence, &question.isactive); err != nil {
			logger.Error("error when scan question row into question struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to get question", errors.New("database error"))
		}
		results = append(results, question)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no questions matching surveyId %s", surveyid))
	}
	return results, nil
}
