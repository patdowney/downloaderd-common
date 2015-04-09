package common

import (
	"time"
)

type TimestampedError struct {
	Time          time.Time
	OriginalError string
}

func (e TimestampedError) Error() string {
	return e.OriginalError
}

func NewTimestampedError(err error, time time.Time) *TimestampedError {
	return &TimestampedError{OriginalError: err.Error(), Time: time}
}
