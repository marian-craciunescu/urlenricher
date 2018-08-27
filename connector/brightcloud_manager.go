package connector

import (
	"errors"
	"fmt"
	"github.com/jpillora/backoff"
	"github.com/marian-craciunescu/urlenricher/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	restEndpoint       = "http://thor.brightcloud.com:80/rest"
	urisPaths          = "/uris"
	MaxIdleConnections = 100
	RequestTimeout     = 500 * time.Millisecond
	CategoriesFilePath = "resources/categories.xml"
	ErrOauthFailed     = errors.New("oauth signature verification failed.Wrong Credentials where used")
)

type brightCloudConnector struct {
	key         string
	secret      string
	httpClient  *http.Client
	CategoryMap map[int]*models.Category
}

func composeEndpoint(u string) (*url.URL, error) {
	rawURL := restEndpoint + urisPaths + "/" + u
	return url.Parse(rawURL)
}

func NewBrightCloudConnector(key, secret string) (Connector, error) {
	m := make(map[int]*models.Category, 0)
	return &brightCloudConnector{key: key, secret: secret, CategoryMap: m}, nil
}

func NewBrightCloudConnector2(key, secret string) (*brightCloudConnector, error) {
	m := make(map[int]*models.Category, 0)

	return &brightCloudConnector{key: key, secret: secret, CategoryMap: m}, nil
}

type retryable struct {
	backoff.Backoff
	maxTries int
}

func (c *brightCloudConnector) readCategoryFile() ([]byte, error) {
	// Open our xmlFile
	xmlFile, err := os.Open(CategoriesFilePath)
	defer xmlFile.Close()
	if err != nil {
		logger.WithError(err).Info("Error opening the categories.xml")
		return nil, err
	}

	// read our opened xmlFile as a byte array.
	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		logger.WithError(err).Info("Error reading the categories.xml")
		return nil, err
	}
	return byteValue, nil
}

func (c *brightCloudConnector) Start() error {
	c.httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: RequestTimeout,
	}
	byteValue, err := c.readCategoryFile()
	if err != nil {
		return nil
	}

	categories, err := models.DecodeCATResponse(byteValue)
	if err != nil {
		logger.WithError(err).Info("Error decoing the categories.xml")
		return err
	}

	for _, cat := range categories.Categories {
		c.CategoryMap[cat.CatID] = &cat
	}

	logger.Infof("Loaded %d categories from file", len(c.CategoryMap))

	return nil
}

func (c *brightCloudConnector) Stop() error {
	return nil
}

func (c *brightCloudConnector) Resolve(u string) (*models.URL, error) {
	logger.WithField("url", u).Info("Resolving")

	bcRequest, err := NewRequest(u, c)
	if err != nil {
		logger.WithError(err).Error("Error composing brightcloud request.")
		return nil, err
	}

	request, err := http.NewRequest(http.MethodGet, bcRequest.normalizedURL, nil)
	if err != nil {
		logger.WithError(err).Error("Error composing http request.")
		return nil, err
	}
	request.Header.Add("Authorization", bcRequest.oauthAuthorization())
	request.Header.Add("Host", "thor.brightcloud.com:80")

	fmt.Println(request)

	resp, err := c.httpClient.Do(request)
	if err != nil {
		logger.WithError(err).Error("Error reading  response")
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.WithError(err).Error("Error decoding bcws body response")
		return nil, err
	}
	fmt.Printf("Raw UriResponse Body:\n%v\n", string(body))

	r, err := models.DecodeURIResponse(body)
	if err != nil {
		logger.WithError(err).Error("Error decoding xml  response")
		return nil, err
	}

	if r.StatusCode != 200 {
		return nil, ErrOauthFailed
	}

	return &models.URL{
		Address:    u,
		Reputation: "bad",
		Ts:         time.Now().UTC(),
	}, nil

}
