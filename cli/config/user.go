package config

type User struct {
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
}
