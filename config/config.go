package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found. Generating default config...")

			viper.SetDefault("settings.mqtt_broker", "localhost")
			viper.SetDefault("settings.mqtt_port", 1883)
			viper.SetDefault("settings.mqtt_client_id", "lscdoorbellmqtt")
			viper.SetDefault("settings.mqtt_username", "guest")
			viper.SetDefault("settings.mqtt_password", "guest")

			err = viper.SafeWriteConfig()
			if err != nil {
				panic(err)
			}
		} else {
			panic("Unknown error while retrieving config")
		}
	}
}

func GetString(setting string) string {
	foundSetting := viper.GetString(setting)
	if foundSetting == "" {
		panic("Confighandler: Could not find item: " + setting)
	}
	return foundSetting
}

func GetInt64(setting string) int64 {
	foundSetting := viper.GetInt64(setting)
	if foundSetting == 0 {
		panic("Confighandler: Could not find item: " + setting)
	}
	return foundSetting
}
