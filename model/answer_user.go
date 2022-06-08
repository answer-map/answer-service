package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AnswerUser struct {
	UserID string `json:"userID"`
}

func (answerUser *AnswerUser) Validate() error {
	return validation.ValidateStruct(answerUser,
		validation.Field(&answerUser.UserID, validation.Required))
}
