package rest

import (
	"github.com/golang/mock/gomock"
	"github.com/marian-craciunescu/urlenricher/config"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

//go:generate mockgen -destination=mock_cache_endpoint_test.go -mock_names Endpoint=MockCacheEndpoint -package=rest github.com/marian-craciunescu/urlenricher/cachestore Endpoint
//go:generate mockgen -destination=mock_metrics_endpoint_test.go -mock_names Endpoint=MockMetricsEndpoint -package=rest github.com/marian-craciunescu/urlenricher/metrics Endpoint
func TestAPIServer_StartStop(t *testing.T) {
	a := assert.New(t)
	randomPort := random(8000, 10000)
	c := config.Config{ServerPort: randomPort}

	ctrl := gomock.NewController(t)
	endpoint := NewMockCacheEndpoint(ctrl)
	mEndpoint := NewMockMetricsEndpoint(ctrl)
	srv := NewAPIServer(&c, endpoint, mEndpoint)

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
