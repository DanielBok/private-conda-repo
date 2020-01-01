package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"private-conda-repo/store"
)

func main() {
	setLogger()
	initStore()

	app := NewApp()
	app.updateIndexer()
	<-app.runServers()
}

func setLogger() {
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func initStore() {
	s, err := store.New()
	if err != nil {
		log.Fatalln(err)
	}
	if err := s.Migrate(); err != nil {
		log.Fatalln(err)
	}
}
