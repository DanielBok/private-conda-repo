package conda

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"private-conda-repo/config"
)

var drivers = make(map[string]Conda)

func New() (Conda, error) {
	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	name := strings.ToLower(strings.TrimSpace(conf.Conda.Type))
	return drivers[name], nil
}

func Register(name string, c Conda) {
	name = strings.ToLower(strings.TrimSpace(name))

	if name == "" {
		log.Fatalln("cannot register empty name in conda drivers")
	} else if _, duplicated := drivers[name]; duplicated {
		log.Fatalf("%s already registered in conda drivers\n", name)
	}

	drivers[name] = c
}
