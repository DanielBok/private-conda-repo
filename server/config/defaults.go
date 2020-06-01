package config

import (
	"runtime"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var defaultKeyValue = map[string]string{}

var defaultKeyValueMap = map[string]map[string]interface{}{
	"indexer.mount_folder": {
		"win": "C:/temp/condapkg",
		"lin": "/var/condapkg",
	},
}

func setDefaults() error {
	var os string
	switch runtime.GOOS {
	case "windows":
		os = "win"
	case "linux":
		os = "lin"
	default:
		return errors.Errorf("Unsupported platform: %s", runtime.GOOS)
	}

	for key, valueMap := range defaultKeyValueMap {
		if val := viper.GetString(key); val == "" {
			viper.Set(key, valueMap[os])
		}
	}

	for key, value := range defaultKeyValue {
		viper.Set(key, value)
	}

	return nil
}
