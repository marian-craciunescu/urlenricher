package connector

import (
	"fmt"
	"github.com/jpillora/backoff"
	"github.com/marian-craciunescu/urlenricher/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/dghubble/oauth1"
	goauth "github.com/gillesdemey/go-oauth"
)

var (
	restEndpoint       = "http://thor.brightcloud.com:80/rest"
	urisPaths          = "/uris"
	MaxIdleConnections = 100
	RequestTimeout     = 500 * time.Millisecond
)

type brightCloudConnector struct {
	key        string
	secret     string
	httpClient *http.Client
}

func composeEndpoint(u string) (*url.URL, error) {
	rawURL := restEndpoint + urisPaths + "/" + u
	return url.Parse(rawURL)
}

func NewBrightCloudConnector(key, secret string) (Connector, error) {

	return &brightCloudConnector{key: key, secret: secret}, nil
}

type retryable struct {
	backoff.Backoff
	maxTries int
}

func (c *brightCloudConnector) Start() error {
	c.httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: RequestTimeout,
	}
	return nil
}

func (c *brightCloudConnector) Stop() error {
	return nil
}

func (c *brightCloudConnector) Resolve(u string) (*models.URL, error) {
	logger.WithField("url", u).Info("Return dummy url response")

	ur, err := composeEndpoint(u)
	if err != nil {
		logger.WithError(err).Error("Error composing the url")
	}

	ts := time.Now().Unix()
	n := "63839d83b0eb762461e3c164bef291ce"
	params := map[string]string{

		// any additional headers for your request

		"oauth_nonce":     "b636113e3918e367e5245bde77f85a47",
		"oauth_timestamp": "1534857839",

		"oauth_consumer_key": c.key,
		"oauth_version":      "1.0",
	}

	signature := goauth.Sign("HMAC-SHA1", "GET", ur.String(), params, c.key, c.secret)

	q := ur.Query()
	q.Set("oauth_signature", signature)
	q.Set("oauth_signature_method", "HMAC-SHA1")
	q.Set("oauth_consumer_key", c.key)
	q.Set("oauth_token", "")
	q.Set("oauth_nonce", n)
	q.Set("oauth_timestamp", fmt.Sprintf("%d", ts))
	q.Set("oauth_version", "1.0")

	encoded := q.Encode()
	ur.RawQuery = encoded

	logger.WithField("ur", ur.String()).WithField("esc", ur.RequestURI()).Info("sth")
	request, err := http.NewRequest(http.MethodGet, ur.String(), nil)
	if err != nil {
		logger.WithError(err).Error("Error compsing  request")
		return nil, err
	}
	request.Header.Add("Content-type", "text/xml")

	logger.WithField("signature", signature).Info("signature")

	config := oauth1.NewConfig(c.key, c.secret)
	token := oauth1.NewToken("", "")

	httpClient := config.Client(oauth1.NoContext, token)

	resp, err := httpClient.Do(request)
	if err != nil {
		logger.WithError(err).Error("Error reading  response")
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Raw Response Body:\n%v\n", string(body))

	return &models.URL{
		Address:    u,
		Reputation: "bad",
		Ts:         time.Now().UTC(),
	}, nil

}
