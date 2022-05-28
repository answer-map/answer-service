package service

import "fmt"

var (
	ErrBadRequest       = fmt.Errorf("bad request")
	ErrResourceNotFound = fmt.Errorf("resource not found")
	ErrConflict         = fmt.Errorf("conflict")
)
