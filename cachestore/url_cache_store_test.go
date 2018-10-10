package cachestore

import (
	"github.com/golang/mock/gomock"
	"github.com/marian-craciunescu/urlenricher/connector"
	"github.com/marian-craciunescu/urlenricher/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

//go:generate mockgen -destination=mock_connector_test.go -mock_names Endpoint=MockConnector -package=cachestore github.com/marian-craciunescu/urlenricher/connector Connector
func TestURLCacheStore_Dump(t *testing.T) {
	a := assert.New(t)

	name, err := ioutil.TempDir("/tmp", "test")
	a.NoError(err)
	defer os.Remove(name)

	mockCtrl := gomock.NewController(t)
	mockCon := NewMockConnector(mockCtrl)
	store, err := NewURLCacheStore(10, 11, mockCon, name)
	a.NoError(err)

	u1 := createUrl("www.google.com", 80, 0)
	u2 := createUrl("www.facebook.com", 70, 1)
	err = store.save(u1.Address, u1)
	a.NoError(err)
	err = store.save(u2.Address, u2)
	a.NoError(err)

	n, err := store.Dump()
	a.NoError(err)
	a.Equal(2, n)
}

func TestStart(t *testing.T) {
	a := assert.New(t)

	name, err := ioutil.TempDir("/tmp", "test")
	a.NoError(err)
	defer os.Remove(name)
	mockCtrl := gomock.NewController(t)
	mockCon := NewMockConnector(mockCtrl)
	store, err := NewURLCacheStore(10, 11, mockCon, name)
	a.NoError(err)

	err = store.Start()
	a.NoError(err)

}

func TestSanityRestart(t *testing.T) {
	a := assert.New(t)

	name, err := ioutil.TempDir("/tmp", "test")
	a.NoError(err)
	defer os.Remove(name)
	mockCtrl := gomock.NewController(t)
	mockCon := NewMockConnector(mockCtrl)
	store, err := NewURLCacheStore(10, 11, mockCon, name)
	a.NoError(err)

	err = store.Start()
	a.NoError(err)

	u1 := createUrl("www.google.com", 80, 0)
	u2 := createUrl("www.facebook.com", 70, 1)
	err = store.save(u1.Address, u1)
	a.NoError(err)
	err = store.save(u2.Address, u2)
	a.NoError(err)

	n, err := store.Dump()
	a.NoError(err)
	a.Equal(2, n)

	newStore, err := NewURLCacheStore(10, 11, mockCon, name)
	a.NoError(err)
	err = newStore.Start()
	a.NoError(err)
	a.Equal(2, newStore.Size())

	url, err := newStore.get("www.google.com")
	a.NoError(err)
	a.Equal(80, url.ReputationPercentage)
}

func TestResolve(t *testing.T) {
	a := assert.New(t)

	name, err := ioutil.TempDir("/tmp", "test")
	a.NoError(err)
	defer os.Remove(name)
	mockCtrl := gomock.NewController(t)
	mockCon := NewMockConnector(mockCtrl)
	store, err := NewURLCacheStore(10, 11, mockCon, name)
	a.NoError(err)

	err = store.Start()
	a.NoError(err)

	u1 := createUrl("www.google.com", 80, 0)
	mockCon.EXPECT().Resolve("www.google.com").Return(u1, nil)

	url, err := store.get("www.google.com")
	a.NoError(err)
	a.Equal(80, url.ReputationPercentage)
}

func TestResolveError(t *testing.T) {
	a := assert.New(t)

	name, err := ioutil.TempDir("/tmp", "test")
	a.NoError(err)
	defer os.Remove(name)
	mockCtrl := gomock.NewController(t)
	mockCon := NewMockConnector(mockCtrl)
	store, err := NewURLCacheStore(10, 11, mockCon, name)
	a.NoError(err)

	err = store.Start()
	a.NoError(err)

	mockCon.EXPECT().Resolve("www.google.com").Return(nil, connector.ErrOauthFailed)

	_, err = store.get("www.google.com")
	a.Error(connector.ErrOauthFailed, err.Error())
	a.Equal(0, store.Size())
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
