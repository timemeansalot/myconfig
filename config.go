package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReadConfig(configFilePath string) error {
	// Initialize a new viper configuration object
	viper.SetConfigFile(configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error reading config file: %s", err)
	}
	return nil
}
