package connector

import (
	"github.com/marian-craciunescu/urlenricher/models"
	"time"
)

type brightCloudConnector struct {
	key    string
	secret string
}

func NewBrightCloudConnector(key, secret string) (Connector, error) {

	return &brightCloudConnector{key, secret}, nil
}

func (c *brightCloudConnector) Start() error {
	return nil
}

func (c *brightCloudConnector) Stop() error {
	return nil
}

func (c *brightCloudConnector) Resolve(u string) (*models.URL, error) {
	logger.WithField("url", u).Info("Return dummy url response")

	return &models.URL{
		Address:    u,
		Reputation: "bad",
		Ts:         time.Now().UTC(),
	}, nil

}
