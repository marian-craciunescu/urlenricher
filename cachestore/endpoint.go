package cachestore

import (
	"github.com/labstack/echo"
	"github.com/marian-craciunescu/urlenricher/models"
)

type Endpoint interface {
	Save(url *models.URL) error
	Get(baseURL string) (*models.URL, error)
	Delete(baseURL string) error
	Resolve(ctx echo.Context) error
}

type urlEndpoint struct {
	store *URLCacheStore
}

func NewURLEndpoint(store *URLCacheStore) Endpoint {
	return &urlEndpoint{store}
}

func (e *urlEndpoint) Save(url *models.URL) error {
	return e.store.save(url.Address, url)
}
func (e *urlEndpoint) Get(baseURL string) (*models.URL, error) {
	return e.store.get(baseURL)
}
func (e *urlEndpoint) Delete(baseURL string) error {
	return models.ErrNotImplemented
}

func (e *urlEndpoint) Resolve(ctx echo.Context) error {
	//originalUrl =ctx.Request().URL.
	return nil
}
