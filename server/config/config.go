package config

import (
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Admin      *AdminProfile  `mapstructure:"admin"`
	Indexer    *IndexerConfig `mapstructure:"indexer"`
	DB         *DbConfig      `mapstructure:"db"`
	FileServer *ServerConfig  `mapstructure:"fileserver"`
	AppServer  *ServerConfig  `mapstructure:"api"`
	TLS        *TLSConfig     `mapstructure:"tls"`
}

type IConfig interface {
	Init() error
}

type AdminProfile struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

const prefix = "pcr"

func New() (*AppConfig, error) {
	if err := setConfigDirectory(); err != nil {
		return nil, err
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "could not read in configuration")
	}

	if err := setDefaults(); err != nil {
		return nil, err
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal config")
	}

	// Check all configurations are valid
	for _, c := range []IConfig{
		config.Indexer,
	} {
		err := c.Init()
		if err != nil {
			return nil, err
		}
	}

	return &config, nil
}

func setConfigDirectory() error {
	usr, err := user.Current()
	if err != nil {
		return errors.Wrap(err, "could not get user directory")
	}

	switch runtime.GOOS {
	case "windows":
		viper.AddConfigPath("C:/Projects/private-conda-repo")
		viper.AddConfigPath("C:/Projects/private-conda-repo/server")
	case "linux":
		viper.AddConfigPath("/var/private-conda-repo")

	default:
		return errors.Errorf("Unsupported platform: %s", runtime.GOOS)
	}
	viper.AddConfigPath(filepath.Join(usr.HomeDir, "private-conda-repo"))

	return nil
}
