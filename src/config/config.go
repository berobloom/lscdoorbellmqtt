package config

import (
	"fmt"
	"lscdoorbellmqtt/logger"
	"lscdoorbellmqtt/utils"

	"github.com/spf13/viper"
)

func Init() {
	dirPath := utils.GetExecutableSourceDir()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dirPath)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found. Generating default config...")

			viper.SetDefault("settings.mqtt_broker", "localhost")
			viper.SetDefault("settings.mqtt_port", 1883)
			viper.SetDefault("settings.mqtt_client_id", "lscdoorbellmqtt")
			viper.SetDefault("settings.mqtt_username", "guest")
			viper.SetDefault("settings.mqtt_password", "guest")
			viper.SetDefault("settings.log_level", "INFO")

			err = viper.SafeWriteConfig()
			if err != nil {
				logger.Fatal(fmt.Sprintf("Failed to write default config: %v", err))
			}
		} else {
			logger.Fatal(fmt.Sprintf("Failed to read config: %v", err))
		}
	}
}

func GetString(setting string) string {
	foundSetting := viper.GetString(setting)
	if foundSetting == "" {
		logger.Fatal("Confighandler: Could not find item: " + setting)
	}
	return foundSetting
}

func GetInt64(setting string) int64 {
	foundSetting := viper.GetInt64(setting)
	if foundSetting == 0 {
		logger.Fatal("Confighandler: Could not find item: " + setting)
	}
	return foundSetting
}
