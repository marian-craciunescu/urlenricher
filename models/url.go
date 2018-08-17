package models

import (
	"errors"
	"time"
)

type URL struct {
	Address    string
	Reputation string
	Ts         time.Time
}

var ErrNotImplemented = errors.New("ERR dummy method to Implement")
