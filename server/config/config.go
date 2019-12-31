package config

import (
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Admin        adminProfile `mapstructure:"admin"`
	Conda        condaConfig  `mapstructure:"conda"`
	DB           database     `mapstructure:"db"`
	UserConfig   userConfig   `mapstructure:"user"`
	FileServer   server       `mapstructure:"fileserver"`
	AppServer    server       `mapstructure:"application"`
	Decompressor decompressor `mapstructure:"decompressor"`
	TLS          tls          `mapstructure:"tls"`
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

type server struct {
	Port int `mapstructure:"port"`
}

type decompressor struct {
	Type string `mapstructure:"type"`
}

const prefix = "PCR"

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

	if err := config.UserConfig.init(); err != nil {
		return nil, err
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
