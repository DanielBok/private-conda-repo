package config

type Channel struct {
	Channel  string `mapstructure:"channel" yaml:"channel"`
	Password string `mapstructure:"password" yaml:"password"`
}
