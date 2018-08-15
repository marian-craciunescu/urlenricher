package main

import (
	"github.com/marian-craciunescu/urlenricher/config"
	"github.com/marian-craciunescu/urlenricher/rest"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	log.WithField("time", time.Now().Format(time.RFC3339)).Info("Starting  urlenricher")

	activeProfile := config.Dev

	conf, err := config.ReadConfig(activeProfile.PropertyFile())
	if err != nil {
		log.WithError(err).Error("Could not read config fail.Exiting")
		os.Exit(-1)
	}

	apiServer := rest.NewAPIServer(conf)
	err = apiServer.Start()
	if err != nil {
		log.WithError(err).Error("Error starting rest server")
		os.Exit(-2)
	}

}
