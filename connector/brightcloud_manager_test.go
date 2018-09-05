package connector

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_StartStop(t *testing.T) {
	a := assert.New(t)
	key := "aaaaaaaaaaaa"
	secret := "bbbbbbbbbbbbbbbbbbbb"
	categoriesFilePath = "../" + categoriesFilePath
	c, err := NewBrightCloudConnector2(key, secret)
	a.NoError(err)

	err = c.Start()
	a.NoError(err)
	a.Equal(83, len(c.CategoryMap))
	err = c.Stop()
	a.NoError(err)
	for i, cc := range c.CategoryMap {
		fmt.Println(cc)
		fmt.Println(c.CategoryMap[i])
	}
}

func Test_ResolveInvalidOauth(t *testing.T) {
	a := assert.New(t)
	key := "aaaaaaaaaaaa"
	secret := "bbbbbbbbbbbbbbbbbbbb"
	categoriesFilePath = "../" + categoriesFilePath
	c, err := NewBrightCloudConnector2(key, secret)
	a.NoError(err)

	err = c.Start()
	a.NoError(err)

	u, err := c.Resolve("www.google.com")
	a.Nil(u)
	a.Equal(ErrOauthFailed, err)
}
