package rest

import (
	"github.com/marian-craciunescu/urlenricher/config"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestAPIServer_StartStop(t *testing.T) {
	a := assert.New(t)
	randomPort := random(8000, 10000)

	c := config.Config{ServerPort: randomPort}

	srv := NewAPIServer(&c)

	err := srv.Start()
	a.NoError(err)

	time.Sleep(100 * time.Millisecond)

	err = srv.Stop()
	a.NoError(err)

}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
