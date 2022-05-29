package model

import (
	"time"

	"github.com/answer-map/answer-repository/entity"
	"github.com/volatiletech/null/v8"
)

type AnswerEvent struct {
	EventID        string      `json:"eventID"`
	EventType      string      `json:"eventType"`
	EventTimestamp time.Time   `json:"eventTimestamp"`
	AnswerKey      string      `json:"answerKey"`
	AnswerValue    null.String `json:"answerValue,omitempty"`
}

func FromEntity(e *entity.AnswerEvent) AnswerEvent {
	return AnswerEvent{
		EventID:        e.EventID,
		EventType:      e.EventType,
		EventTimestamp: e.EventTimestamp,
		AnswerKey:      e.AnswerKey,
		AnswerValue:    e.AnswerValue,
	}
}
