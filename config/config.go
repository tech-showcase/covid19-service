package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		ServiceName string `json:"service_name"`
		Address     string `json:"address"`
		Tracer      Tracer `json:"tracer"`
	}

	Tracer struct {
		AgentAddress string `json:"agent_address"`
	}
)

var Instance Config

func init() {
	viper.SetDefault("CONFIG_FILEPATH", ".")
	viper.BindEnv("CONFIG_FILEPATH")
	viper.SetDefault("CONFIG_FILENAME", ".env")
	viper.BindEnv("CONFIG_FILENAME")

	viper.SetDefault("SERVICE_NAME", "covid19-service")
	viper.BindEnv("SERVICE_NAME")

	viper.SetDefault("ADDRESS", "http://localhost")
	viper.BindEnv("ADDRESS")

	viper.SetDefault("TRACER_AGENT_ADDRESS", "localhost:5775")
	viper.BindEnv("TRACER_AGENT_ADDRESS")
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
		ServiceName: viper.Get("SERVICE_NAME").(string),
		Address:     viper.Get("ADDRESS").(string),
		Tracer: Tracer{
			AgentAddress: viper.Get("TRACER_AGENT_ADDRESS").(string),
		},
	}
	return
}
