package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	setLogger()

	app, err := NewApp()
	if err != nil {
		log.Fatal(err)
	}

	<-app.runServers()
}

func setLogger() {
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
