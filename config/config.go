package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		Address string `json:"address"`
	}
)

var Instance Config

func init() {
	viper.SetDefault("CONFIG_FILEPATH", ".")
	viper.BindEnv("CONFIG_FILEPATH")
	viper.SetDefault("CONFIG_FILENAME", ".env")
	viper.BindEnv("CONFIG_FILENAME")

	viper.SetDefault("MOVIE_SERVER_ADDRESS", "http://localhost")
	viper.BindEnv("MOVIE_SERVER_ADDRESS")
	viper.SetDefault("MOVIE_API_KEY", "api-key")
	viper.BindEnv("MOVIE_API_KEY")
}

func Parse() (config Config, err error) {
	configFilepath, configFilename := getConfigFile()
	viper.SetConfigName(configFilename)
	viper.SetConfigType("env")
	viper.AddConfigPath(configFilepath)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	config = getConfig()
	return
}

func getConfigFile() (configFilepath string, configFilename string) {
	configFilepath = viper.Get("CONFIG_FILEPATH").(string)
	configFilename = viper.Get("CONFIG_FILENAME").(string)
	return
}

func getConfig() (config Config) {
	config = Config{
		Address: viper.Get("ADDRESS").(string),
	}
	return
}
