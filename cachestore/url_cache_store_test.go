package cachestore

import (
	"github.com/golang/mock/gomock"
	"github.com/marian-craciunescu/urlenricher/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//go:generate mockgen -destination=mock_connector_test.go -mock_names Endpoint=MockConnector -package=cachestore github.com/marian-craciunescu/urlenricher/connector Connector
func TestURLCacheStore_Dump(t *testing.T) {
	a := assert.New(t)

	mockCtrl := gomock.NewController(t)
	mockCon := NewMockConnector(mockCtrl)
	store, err := NewURLCacheStore(10, 11, mockCon)
	a.NoError(err)

	u1 := createUrl("www.google.com", 80, 0)
	u2 := createUrl("www.facebook.com", 70, 1)
	err = store.save(u1.Address, u1)
	a.NoError(err)
	err = store.save(u2.Address, u2)
	a.NoError(err)

	n, err := store.Dump()
	a.NoError(err)
}

func createUrl(url string, rep, alt1 int) *models.URL {

	urlCategory := models.UrlCategories{
		ID:         1,
		Confidence: rep,
		Name:       "Food and dining",
		Group:      "Security",
	}

	u := &models.URL{
		Address:              url,
		ReputationPercentage: rep,
		SubdomainNumber:      1,
		Ts:                   time.Now().UTC(),
		Categories:           []models.UrlCategories{urlCategory},
	}
	return u
}
