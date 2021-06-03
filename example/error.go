package main

import "time"

type APIStatus string

const (
	ACTIVE APIStatus = "active"
	ERROR  APIStatus = "error"
)

func (A APIStatus) GetValues() []interface{} {
	return []interface{}{ACTIVE, ERROR}
}

type BaseResponse struct {
	// human-readable message
	Message string `json:"message"`
	// Corresponding API action
	APIID        string `json:"apiID,omitempty"`
	OrginalError error  `json:"-"`
}

type APIError struct {
	BaseResponse
	Status APIStatus       `json:"status"`
	Causes map[string]bool `json:"causes"`
	// Optional error body
	// ID to identify the request that caused this error
	RequestID string    `json:"requestID,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}
