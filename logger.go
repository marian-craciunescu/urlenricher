package main

import (
	"net"
	"os"

	"fmt"
	"github.com/bshuster-repo/logrus-logstash-hook"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

func initLogging(lvl log.Level, useElk bool) *log.Entry {
	name, err := os.Hostname()
	if err != nil {
		fmt.Println("Could not get hostname.Using default")
		name = "localhost"
	}
	ip, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Could not get host ip.Using default")
		name = "127.0.0.1"
	}

	logger := log.WithField("app_name", "urlenricher").
		WithField("host_name", name).
		WithField("host_ip", ip)

	if useElk {
		logger.Infof("using ELK %v", useElk)
		conn, hostErr := net.Dial("tcp", "elk01.esolutions.ro:4560")
		if hostErr != nil {
			logger.Fatal(hostErr)
		}
		hook := logrustash.New(conn, logrustash.DefaultFormatter(log.Fields{}))
		log.AddHook(hook)
	}
	logger.Infof("Setting log level to %s", lvl)
	log.SetLevel(lvl)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, QuoteEmptyFields: true})
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//log.SetLevel(log.DebugLevel)

	return logger
}
