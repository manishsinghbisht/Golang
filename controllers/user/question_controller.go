package question

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/manishsinghbisht/bookstore_oauth-go/oauth"
	"github.com/manishsinghbisht/greenleaf-api/domain/question"
	"github.com/manishsinghbisht/greenleaf-api/services"
	"github.com/manishsinghbisht/utils-go/rest_errors"
)

func getUserId(userIdParam string) (int64, rest_errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func GetQuestion(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	user, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	if oauth.GetCallerId(c.Request) == user.Id {
		questionId := c.Query("questionId")
		user, getErr := services.QuestionService.GetQuestion(questionId)
		if getErr != nil {
			c.JSON(getErr.Status(), getErr)
			return
		}
		c.JSON(http.StatusOK, question.Question.Marshall(c.GetHeader("X-Public") == "true"))
	}
}

func GetQuestionBySurvey(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	user, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	if oauth.GetCallerId(c.Request) == user.Id {
		surveyId := c.Query("surveyId")
		user, getErr := services.QuestionService.GetQuestionBySurvey(surveyId)
		if getErr != nil {
			c.JSON(getErr.Status(), getErr)
			return
		}
		c.JSON(http.StatusOK, question.Question.Marshall(c.GetHeader("X-Public") == "true"))
	}
}
