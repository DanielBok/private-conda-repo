package config

type CondaConfig struct {
	Use         string `mapstructure:"use"`
	ImageName   string `mapstructure:"image_name"`
	MountFolder string `mapstructure:"mount_folder"`
	Type        string `mapstructure:"type"`
}
