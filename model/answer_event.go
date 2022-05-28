package model

import (
	"time"

	"github.com/volatiletech/null/v8"
)

type AnswerEvent struct {
	EventID        string      `json:"event_id"`
	EventType      string      `json:"event_type"`
	EventTimestamp time.Time   `json:"event_timestamp"`
	AnswerKey      string      `json:"answer_key"`
	AnswerValue    null.String `json:"answer_value,omitempty"`
}
