package config

import (
	"os/user"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Admin      adminProfile `mapstructure:"admin"`
	Conda      condaConfig  `mapstructure:"conda"`
	DB         database     `mapstructure:"db"`
	Salt       string       `mapstructure:"salt"`
	RepoServer server       `mapstructure:"repository"`
	AppServer  server       `mapstructure:"application"`
}

type adminProfile struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type condaConfig struct {
	ImageName   string `mapstructure:"image_name"`
	MountFolder string `mapstructure:"mount_folder"`
	Type        string `mapstructure:"type"`
}

type database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbname"`
	Type     string `mapstructure:"type"`
}

type server struct {
	Port int `mapstructure:"port"`
}

var (
	once        sync.Once
	config      AppConfig
	configError error
)

const prefix = "PCR"

func New() (*AppConfig, error) {
	once.Do(func() {
		if err := setConfigDirectory(); err != nil {
			configError = err
			return
		}

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.SetEnvPrefix(prefix)
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			configError = errors.Wrap(err, "could not read in configuration")
			return
		}

		if err := setDefaults(); err != nil {
			configError = err
			return
		}

		if err := viper.Unmarshal(&config); err != nil {
			configError = errors.Wrap(err, "could not unmarshal config")
			return
		}
	})

	return &config, configError
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
