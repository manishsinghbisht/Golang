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
	Id          string  `json:"Id"`
	surveyid   string `json:"survey_id"`
	question    string `json:"question"`
	isactive      int `json:"isactive"`
}

type Questions []Question

func (question *Question) Validate() rest_errors.RestErr {
	question.question = strings.TrimSpace(question.question)

	return nil
}
