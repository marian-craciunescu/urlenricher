package cachestore

import (
	"errors"
	"github.com/hashicorp/golang-lru"
	"github.com/marian-craciunescu/urlenricher/connector"
	"github.com/marian-craciunescu/urlenricher/models"
)

var ErrLRUCacheValue = errors.New("Unkown value was found ")

type URLCacheStore struct {
	lruCache     *lru.Cache
	size         int
	maxAgeInDays int
	conn         connector.Connector
}

func NewURLCacheStore(size, maxAge int, conn connector.Connector) (*URLCacheStore, error) {
	l, err := lru.New(size)
	if err != nil {
		return nil, err
	}
	return &URLCacheStore{lruCache: l, size: size, maxAgeInDays: maxAge, conn: conn}, nil
}

func (ucs *URLCacheStore) Test() {
	for i := 0; i < ucs.size; i++ {
		ucs.lruCache.Add(i, i)
	}
}

// evicts all LRU cache to Disk
func (ucs *URLCacheStore) Dump() error {

	return models.ErrNotImplemented
}

func (ucs *URLCacheStore) load() error {
	return models.ErrNotImplemented
}

func (ucs *URLCacheStore) save(originalURL string, u *models.URL) error {
	evicted := ucs.lruCache.Add(originalURL, u)
	logger.WithField("url", originalURL).WithField("evicted", evicted).Debug("Saved url")
	return nil
}

func (ucs *URLCacheStore) get(u string) (*models.URL, error) {
	url, ok := ucs.lruCache.Get(u)
	if !ok {
		logger.WithField("url", url).Debug("Requested url was not found.Resolving")
		return ucs.resolve(u)
	}
	original, ok := url.(*models.URL)
	if !ok {
		return nil, ErrLRUCacheValue
	}

	return original, nil

}

func (ucs *URLCacheStore) delete(u models.URL) error {
	return models.ErrNotImplemented
}

func (ucs *URLCacheStore) resolve(u string) (*models.URL, error) {
	logger.Info("cache store resolve")
	url, err := ucs.conn.Resolve(u)
	if err != nil {
		logger.WithError(err).Error("Error resolving url.")
		return nil, err
	}
	err = ucs.save(u, url)
	if err != nil {
		logger.WithError(err).Error("Error saving url.")
		return nil, err
	}
	return url, nil
}
