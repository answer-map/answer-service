package model

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// CreateUserRequest represents the params required to perform the create user operation.
type CreateUserRequest = AnswerUser

// CreateRequest represents the params required to perform the create operation.
type CreateRequest = AnswerMap

// UpdateRequest represents the params required to perform the update operation.
type UpdateRequest = AnswerMap

// DeleteRequest represents the params required to perform the delete operation.
type DeleteRequest = AnswerMapKey

// GetRequest represents the params required to perform the get operation.
type GetRequest = AnswerMapKey

// GetResponse represents the response from the get operation.
type GetResponse = AnswerMap

// GetHistoryRequest represents the params required to perform the get history operation.
type GetHistoryRequest struct {
	MinimumEventTimestamp time.Time `json:"minimumEventTimestamp"`
	PageSize              uint32    `json:"pageSize"`
}

func NewGetHistoryRequest(query url.Values) (*GetHistoryRequest, error) {
	minimumEventTimestampQP := query.Get("minimumEventTimestamp")
	pageSizeQP := query.Get("pageSize")

	if len(minimumEventTimestampQP) == 0 || len(pageSizeQP) == 0 {
		return nil, fmt.Errorf("missing required query params")
	}

	minimumEventTimestamp, err := time.Parse(time.RFC3339, minimumEventTimestampQP)
	if err != nil {
		return nil, fmt.Errorf("failed to parse minimumEventTimestamp")
	}

	pageSize64, err := strconv.ParseUint(pageSizeQP, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pageSize")
	}

	return &GetHistoryRequest{MinimumEventTimestamp: minimumEventTimestamp, PageSize: uint32(pageSize64)}, nil
}

// GetHistoryResponse represents the response from the get history operation.
type GetHistoryResponse = []AnswerEvent
