package connector

import "github.com/marian-craciunescu/urlenricher/models"

type Connector interface {
	Start() error
	Stop() error
	Resolve(u string) (*models.URL, error)
}
