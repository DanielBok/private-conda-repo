package conda

import (
	"strings"

	"github.com/pkg/errors"
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
	if drv, ok := drivers[name]; !ok {
		return nil, errors.Errorf("Unknown conda repository driver: '%s'", conf.Conda.Type)
	} else {
		return drv, nil
	}
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
