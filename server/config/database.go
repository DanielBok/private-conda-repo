package config

import "fmt"

type DbConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DbName      string `mapstructure:"dbname"`
	AutoMigrate bool   `mapstructure:"auto_migrate"`
}

func (d *DbConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.User, d.Password, d.DbName)
}

func (d *DbConfig) MaskedConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=XXXXXXXXX dbname=%s sslmode=disable",
		d.Host, d.Port, d.User, d.DbName)
}
