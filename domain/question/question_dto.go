package question

import (
	// "database/sql"
	"strings"
	"github.com/manishsinghbisht/utils-go/rest_errors"
	// "gopkg.in/guregu/null.v4"
)

const (
	StatusActive = "active"
)

type Question struct {
	Id          	string  `json:"Id"`
	surveyid   		string `json:"survey_id"`
	question    	string `json:"question"`
	sequence      	int `json:"sequence"`
	isactive      	int `json:"isactive"`
	answers			[]Answer `json:"answers"`
}

type Answer struct {
	Id          string  `json:"Id"`
	questionid   string `json:"question_id"`
	answer    	string `json:"answer"`
	sequence      int `json:"sequence"`
	isactive      int `json:"isactive"`
}

type Questions []Question

func (question *Question) Validate() rest_errors.RestErr {
	question.question = strings.TrimSpace(question.question)
	if question.question == "" {
		return rest_errors.NewBadRequestError("invalid question")
	}
	return nil
}
