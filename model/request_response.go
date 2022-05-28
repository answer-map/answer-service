package model

import "time"

// CreateRequest represents the params required to perform the create operation.
type CreateRequest = AnswerMap

// UpdateRequest represents the params required to perform the update operation.
type UpdateRequest = AnswerMap

// DeleteRequest represents the params required to perform the delete operation.
type DeleteRequest = AnswerMapKey

// GetRequest represents the params required to perform the get operation.
type GetRequest = AnswerMapKey

// GetResponse represents the response from the get operation.
type GetResponse AnswerMap

// GetHistoryRequest represents the params required to perform the get history operation.
type GetHistoryRequest struct {
	MinimumEventTimestamp time.Time `json:"minimumEventTimestamp"`
	PageSize              uint32    `json:"pageSize"`
}

// GetHistoryResponse represents the response from the get history operation.
type GetHistoryResponse []AnswerEvent
