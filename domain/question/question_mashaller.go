package question

import (
	"encoding/json"
)

type PublicQuestion struct {
	Id       string   `json:"Id"`
	surveyid string   `json:"survey_id"`
	question string   `json:"question"`
	sequence int      `json:"sequence"`
	isactive int      `json:"isactive"`
	answers  []Answer `json:"answers"`
}

type PrivateQuestion struct {
	Id       string `json:"Id"`
	surveyid string `json:"survey_id"`
	question string `json:"question"`
	sequence int    `json:"sequence"`
	isactive int    `json:"isactive"`
}

func (questions Questions) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(questions))
	for index, question := range questions {
		result[index] = question.Marshall(isPublic)
	}
	return result
}

func (question *Question) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicQuestion{
			Id:       question.Id,
			surveyid: question.surveyid,
			question: question.question,
			sequence: question.sequence,
			isactive: question.isactive,
			answers:  question.answers,
		}
	}
	questionJson, _ := json.Marshal(question)
	var privateQuestion PrivateQuestion
	json.Unmarshal(questionJson, &privateQuestion)
	return privateQuestion

}
