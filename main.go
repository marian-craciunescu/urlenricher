package main

import (
	"github.com/marian-craciunescu/urlenricher/config"
	"github.com/marian-craciunescu/urlenricher/rest"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.WithField("time", time.Now().Format(time.RFC3339)).Info("Starting  urlenricher")

	activeProfile := config.Dev

	conf, err := config.ReadConfig(activeProfile.PropertyFile())
	if err != nil {
		panic("Could not read config fail.Exiting")
	}

	initLogging(conf.LogLevel, conf.ElkLogging)

	apiServer := rest.NewAPIServer(conf, nil)
	err = apiServer.Start()
	if err != nil {
		log.WithError(err).Error("Error starting rest server")
		os.Exit(-2)
	}

	waitForTerminationAndExit(func() {
		err := apiServer.Stop()
		if err != nil {
			logger.WithError(err).Error("Failed to correctly stop api server")
		}
	})

}

func waitForTerminationAndExit(callback func()) {
	signalC := make(chan os.Signal)
	signal.Notify(signalC, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signalC
	log.Infof("Got signal '%v' .. exiting gracefully now", sig)
	callback()

	log.Info("Exit gracefully now")
	os.Exit(0)
}
