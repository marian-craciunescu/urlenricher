package urlcache

import "github.com/marian-craciunescu/urlenricher/connector"

type Endpoint interface {
	Save(url *connector.URL) error
	Get(baseURL string) *connector.URL
	Delete(baseURL string)
}
