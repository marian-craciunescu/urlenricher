package models

import (
	"errors"
	"time"
)

type UrlCategories struct {
	ID         int
	Name       string
	Group      string
	Confidence int
}

type URL struct {
	Address              string
	ReputationPercentage int
	Categories           []UrlCategories
	SubdomainNumber      int
	Ts                   time.Time
}

var ErrNotImplemented = errors.New("ERR dummy method to Implement")
