package cachestore

import (
	"encoding/json"
	"errors"
	"github.com/hashicorp/golang-lru"
	"github.com/marian-craciunescu/urlenricher/connector"
	"github.com/marian-craciunescu/urlenricher/models"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var ErrLRUCacheValue = errors.New("Unkown value was found ")
var ErrLRUCacheError = errors.New("Error getting value from cache ")

type URLCacheStore struct {
	lruCache     *lru.Cache
	size         int
	maxAgeInDays int
	conn         connector.Connector
	dumpPath     string
}

func NewURLCacheStore(size, maxAge int, conn connector.Connector, path string) (*URLCacheStore, error) {
	l, err := lru.New(size)
	if err != nil {
		return nil, err
	}
	return &URLCacheStore{lruCache: l, size: size, maxAgeInDays: maxAge, conn: conn, dumpPath: filepath.Join(".", path)}, nil
}

func (ucs *URLCacheStore) Start() error {
	err := ucs.load()
	if err != nil {
		logger.Info("Could not correctly load previous on-disk cache")
		return err
	}
	logger.WithField("cache_size", len(ucs.lruCache.Keys())).Info("Loaded from disk")
	return nil
}

// evicts all LRU cache to Disk
func (ucs *URLCacheStore) Dump() (int, error) {

	allKeys := ucs.lruCache.Keys()

	for i, key := range allKeys {
		u, ok := ucs.lruCache.Get(key)
		if !ok {
			logger.WithField("url", u).Debug("Could not dump")
			return -1, ErrLRUCacheError
		}
		original, ok := u.(*models.URL)
		if !ok {
			return -1, ErrLRUCacheValue
		}
		file, err := os.OpenFile(filepath.Join(ucs.dumpPath, "/", original.Address), os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			logger.WithError(err).WithField("file", filepath.Join(ucs.dumpPath, original.Address)).Error("Could not open file for writing")
			return -1, err
		}

		logger.WithField("file", file.Name()).Info("Writing file")

		defer file.Close()

		bytes, err := json.Marshal(original)
		if err != nil {
			logger.WithError(err).Error("Could not marshal data")
			continue
		}
		n, err := file.Write(bytes)
		if err != nil {
			logger.WithError(err).WithField("total_bytes_written", n).Error("Could not write data in file")
			continue
		}
		logger.WithField("i", i).WithField("key", key).Debug("Wrote on disk number")
	}
	return len(allKeys), nil
}

func (ucs *URLCacheStore) load() error {
	logger.WithField("path", ucs.dumpPath).Info("Using path")
	err := os.MkdirAll(ucs.dumpPath, 0777)
	if err != nil {
		logger.WithError(err).Error("Error creating dump directory")
		return err
	}

	files, err := ioutil.ReadDir(ucs.dumpPath)
	logger.WithField("f", files).Info("Files in folder")
	if err != nil {
		logger.WithError(err).WithField("path", ucs.dumpPath).Error("Could not read data dir for files")
		return err
	}

	for _, file := range files {
		jsonFile, err := os.Open(filepath.Join(ucs.dumpPath, file.Name()))
		defer jsonFile.Close()
		if err != nil {
			logger.WithError(err).WithField("filename", file.Name()).Error("Could not open file for reading")
		} else {

			// read our opened xmlFile as a byte array.
			byteValue, err := ioutil.ReadAll(jsonFile)
			if err != nil {
				logger.WithError(err).WithField("filename", file.Name()).Error("Could not read file")
				continue
			}
			var u *models.URL
			if err := json.Unmarshal(byteValue, &u); err != nil {
				logger.WithError(err).WithField("filename", file.Name()).Error("Could not decode  file as json")
				continue
			}
			if duration := time.Now().Sub(u.Ts).Hours() / 24; int(duration) <= ucs.maxAgeInDays {
				logger.WithField("url", u.Address).WithField("daysPassed", duration).Info("ReSaving")
				err := ucs.save(u.Address, u)
				if err != nil {
					logger.WithError(err).Error("Error resaving url ")
				}
			}

		}
	}

	return nil
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
