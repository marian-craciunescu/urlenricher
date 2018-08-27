package connector

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	nonce2 "github.com/LarryBattle/nonce-golang"
	"github.com/PuerkitoBio/purell"
	"net/url"
	"time"
)

type Request struct {
	initialURL string
	nonce      string
	ts         int64
	url        *url.URL

	baseString         string
	normalizedURL      string
	normalizedMethod   string
	signableParameters string
	signature          string
	connector          *brightCloudConnector
}

func NewRequest(intialURL string, c *brightCloudConnector) (*Request, error) {

	request := Request{
		initialURL: intialURL,
		connector:  c,
	}
	err := request.fromConsumerAndToken(intialURL, c)
	if err != nil {
		return nil, err
	}
	request.build()
	return &request, nil
}

func (r *Request) build() {
	r.getNormalizedMethod()
	r.getNormalizedURL()
	r.getSignableParameters()
	r.sign()
}

func (r *Request) getNormalizedMethod() {
	r.normalizedMethod = "GET"
}

func (r *Request) getNormalizedURL() {
	initialQuery := r.url.RawQuery
	r.url.RawQuery = ""
	r.normalizedURL = r.url.String()
	r.url.RawQuery = initialQuery
}

func (r *Request) getSignableParameters() {
	q := r.url.Query()
	q.Set("oauth_consumer_key", r.connector.key)
	q.Set("oauth_nonce", r.nonce)
	q.Set("oauth_signature_method", "HMAC-SHA1")
	q.Set("oauth_timestamp", fmt.Sprintf("%d", r.ts))
	q.Set("oauth_version", "1.0")

	encodedParams := q.Encode()
	r.signableParameters = encodedParams
}

func (r *Request) fromConsumerAndToken(u string, c *brightCloudConnector) error {
	rawURL, err := purell.NormalizeURLString(restEndpoint+urisPaths+"/"+u,
		purell.FlagRemoveDefaultPort|purell.FlagLowercaseScheme|purell.FlagLowercaseHost|purell.FlagUppercaseEscapes)
	if err != nil {
		logger.WithError(err).Info("Could not normalize url")
		return err
	}

	r.ts = time.Now().Unix()
	r.nonce = nonce2.NewToken()
	r.url, err = url.Parse(rawURL)
	if err != nil {
		logger.WithError(err).Info("Could not construct url")
		return err
	}

	q := r.url.Query()
	q.Set("oauth_consumer_key", c.key)
	q.Set("oauth_nonce", r.nonce)
	q.Set("oauth_timestamp", fmt.Sprintf("%d", r.ts))
	q.Set("oauth_version", "1.0")

	encodedParams := q.Encode()

	r.url.RawQuery = encodedParams

	return nil
}

func (r *Request) sign() {
	baseString := r.normalizedMethod + "&" + url.QueryEscape(r.normalizedURL) + "&" + url.QueryEscape(r.signableParameters)
	keyToSign := url.QueryEscape(r.connector.secret) + "&"
	r.signature = HMACSHA1(keyToSign, baseString)
}

func (r *Request) oauthAuthorization() string {
	return fmt.Sprintf(`OAuth realm="",oauth_version="1.0",oauth_nonce="%s",oauth_timestamp="%d",oauth_consumer_key="%s",oauth_signature_method="HMAC-SHA1",oauth_signature="%s"`,
		r.nonce, r.ts, r.connector.key, url.QueryEscape(r.signature))
}

func HMACSHA1(key, input string) string {
	keyToSign := []byte(key)
	h := hmac.New(sha1.New, keyToSign)
	h.Write([]byte(input))

	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
