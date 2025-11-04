package services

import (
	"errors"
)


// list of possible errors
var (
	ErrDeviceNotFound  = errors.New("device not found")
	ErrInvalidHeartbeat = errors.New("invalid heartbeat payload")
	ErrWrongCSVFormat = errors.New("wrong CSV header. Please double check file")
	ErrNoDeviceStats = errors.New("no device stats found")
)


