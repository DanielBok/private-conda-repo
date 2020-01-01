package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Registry          string  `mapstructure:"registry" yaml:"registry"`
	Channel           Channel `mapstructure:"channel" yaml:"channel"`
	PackageRepository string  `mapstructure:"package_repository" yaml:"package_repository"`
}

var (
	once           sync.Once
	conf           Config
	configFilePath string
)

func New() *Config {
	once.Do(func() {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalln(err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".pcrrc")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalln(errors.Wrap(err, "could not read in config"))
		}

		if err := viper.Unmarshal(&conf); err != nil {
			log.Fatalln(errors.Wrap(err, "could not unmarshal config"))
		}
		conf.Registry = strings.TrimSpace(strings.TrimRight(conf.Registry, "/"))
	})
	return &conf
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalln(err)
	}

	configFilePath = filepath.Join(home, ".pcrrc.yaml")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		file, err := os.Create(configFilePath)
		if err != nil {
			log.Fatalln(errors.Wrap(err, "could not write config"))
		}
		_ = file.Close()
	}
}

func (c *Config) HasRegistry() bool {
	if conf.Registry == "" {
		log.Println("Registry location not set. Please use 'pcr registry set' to specify your private conda repo registry")
		return false
	}
	return true
}

// Saves configuration to the config file
func (c *Config) Save() {
	out, err := yaml.Marshal(conf)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(configFilePath, out, 0666); err != nil {
		log.Fatal(errors.Wrap(err, "error encountered when saving configurations"))
	}
}
