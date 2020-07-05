package services

import (
	"github.com/manishsinghbisht/greenleaf-api/domain/question"
	"github.com/manishsinghbisht/utils-go/date_utils"
	"github.com/manishsinghbisht/utils-go/crypto_utils"
	"github.com/manishsinghbisht/utils-go/rest_errors"
)

var (
	QuestionService questionServiceInterface = &questionService{}
)

type questionService struct{}

type questionServiceInterface interface {
	GetQuestion(string) (*question.Question, rest_errors.RestErr)
	SearchQuestionBySurvey(string) (question.Question, rest_errors.RestErr)
}

func (s *questionService) GetQuestion(Id int64) (*question.Question, rest_errors.RestErr) {
	dao := &question.Question{Id: Id}
	if err := dao.Get(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (s *questionService) SearchQuestionBySurvey(surveyId string) (question.Questions, rest_errors.RestErr) {
	dao := &question.Question{}
	return dao.FindBySurvey(surveyId)
}

