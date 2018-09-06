package connector

import (
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

	a.Equal("Internet Communications", c.CategoryMap[66].CatName)
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
