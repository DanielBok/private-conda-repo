package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"private-conda-repo/config"
	"private-conda-repo/store"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatalf("error getting config: %v", err)
	}

	setLogger()
	initStore(conf)

	app := NewApp(conf)
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

func initStore(conf *config.AppConfig) {
	s, err := store.New(conf)
	if err != nil {
		log.Fatalln(err)
	}
	if err := s.Migrate(); err != nil {
		log.Fatalln(err)
	}
}
