package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		ServiceName string `json:"service_name"`
		Address     string `json:"address"`
		Tracer      Tracer `json:"tracer"`
		Log         Log    `json:"log"`
	}

	Tracer struct {
		AgentAddress string `json:"agent_address"`
	}

	Log struct {
		Filepath string `json:"filepath"`
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

	viper.SetDefault("LOG_FILEPATH", "./server.log")
	viper.BindEnv("LOG_FILEPATH")
}

func Parse() (config Config, err error) {
	configFilepath, configFilename := getConfigFile()
	viper.SetConfigName(configFilename)
	viper.SetConfigType("env")
	viper.AddConfigPath(configFilepath)

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return
		}
	}

	config = getConfig()
	return config, nil
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
		Log: Log{
			Filepath: viper.Get("LOG_FILEPATH").(string),
		},
	}
	return
}
