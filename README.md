# urlenricher
[![Build Status](https://api.travis-ci.org/marian-craciunescu/urlenricher.svg?branch=master)](https://travis-ci.org/marian-craciunescu/urlenricher)
[![Go Report Card](https://goreportcard.com/badge/github.com/marian-craciunescu/urlenricher)](https://goreportcard.com/report/github.com/marian-craciunescu/urlenricher)
[![Coverage Status](https://coveralls.io/repos/github/marian-craciunescu/urlenricher/badge.svg?branch=master)](https://coveralls.io/github/marian-craciunescu/urlenricher?branch=master)

# Purpose:
expose a REST api for URL repudiation.

One connecter is already written(it uses the Brigthcloud api [BRIGTHCLOUD API](https://www.brightcloud.com/web-service/api-documentation)

# USAGE
``` 
go get marian-craciunescu/urlenricher
cd $GOPATH/src/github.com/marian-craciunescu/urlenricher 
go build 

./urlenricher --api_key=BrightcloudApiKey --api_secret=BrightcloudApiSecret
```

# ROADMAP
1) Add more tests and Benchmark
2) Add persistence on disk for the LRU cache (at every restart the LRU cache is deleted)
3) Add more connectors.


