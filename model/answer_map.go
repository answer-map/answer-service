package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AnswerMap struct {
	UserID      string `json:"userID"`
	AnswerKey   string `json:"answerKey"`
	AnswerValue string `json:"answerValue"`
}

type AnswerMapKey struct {
	UserID    string `json:"userID"`
	AnswerKey string `json:"answerKey"`
}

func (answerMap *AnswerMap) Validate() error {
	return validation.ValidateStruct(answerMap,
		validation.Field(&answerMap.UserID, validation.Required),
		validation.Field(&answerMap.AnswerKey, validation.Required),
		validation.Field(&answerMap.AnswerValue, validation.Required))
}

func (answerMap *AnswerMapKey) Validate() error {
	return validation.ValidateStruct(answerMap,
		validation.Field(&answerMap.UserID, validation.Required),
		validation.Field(&answerMap.AnswerKey, validation.Required))
}
