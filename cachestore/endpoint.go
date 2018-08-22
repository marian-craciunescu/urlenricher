package cachestore

import (
	"github.com/labstack/echo"
	"github.com/marian-craciunescu/urlenricher/models"
	"net/http"
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

	logger.Info("Endpoint Resolve")

	url, err := e.store.resolve("www.whatismyclassification.com")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return err
	}
	ctx.JSON(http.StatusOK, url)
	return nil
}
