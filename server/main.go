package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	setLogger()

	log.Info("Initalizing application")
	app, err := NewApp()
	if err != nil {
		log.Fatal("Fatal error encountered during application startup", err)
	}

	<-app.RunServers()
}

func setLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
