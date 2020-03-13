package conda

import (
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var drivers = make(map[string]Conda)

func New(condaDriver string) (Conda, error) {
	if drv, ok := drivers[strings.ToLower(strings.TrimSpace(condaDriver))]; !ok {
		return nil, errors.Errorf("Unknown conda repository driver: '%s'", condaDriver)
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
