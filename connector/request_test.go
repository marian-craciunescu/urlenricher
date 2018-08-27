package connector

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_createRequest(t *testing.T) {
	a := assert.New(t)

	key := "aaaaaaaaaaaa"
	secret := "bbbbbbbbbbbbbbbbbbbb"
	c, err := NewBrightCloudConnector2(key, secret)
	a.NoError(err)

	request, err := NewRequest("www.google.com", c)
	a.NoError(err)

	oauthHeader := request.oauthAuthorization()

	a.True(strings.Contains(oauthHeader, fmt.Sprintf("%d", request.ts)))
	a.True(strings.Contains(oauthHeader, request.nonce))
}
