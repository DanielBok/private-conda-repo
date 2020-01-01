package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"private-conda-repo/config"
	"private-conda-repo/store"
)

func main() {
	setLogger()
	listConfig()
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

func listConfig() {
	conf, err := config.New()
	if err != nil {
		log.Fatalf("error getting config: ", err)
	}
	log.Printf("%+v\n", *conf)
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
