package config

import (
	"runtime"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var defaultKeyValueMap = map[string]map[string]interface{}{
	"CONDA.MOUNT_FOLDER": {
		"WIN": "C:/temp/condapkg",
		"LIN": "/var/condapkg",
	},
}

func setDefaults() error {
	var os string
	switch runtime.GOOS {
	case "windows":
		os = "WIN"
	case "linux":
		os = "LIN"
	default:
		return errors.Errorf("Unsupported platform: %s", runtime.GOOS)
	}

	for key, valueMap := range defaultKeyValueMap {
		if val := viper.GetString(key); val == "" {
			viper.Set(key, valueMap[os])
		}
	}

	return nil
}
