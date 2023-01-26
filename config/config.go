package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	*AppConfig
}

type AppConfig struct {
	ServerConfig  ServerConfig  `mapstructure:"server"`
	MongoDBConfig MongoDBConfig `mapstructure:"mongodb"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type MongoDBConfig struct {
	URI string `mapstructure:"uri"`
}

func NewConfiguration() *Config {
	appConfig := &AppConfig{}
	appConfig.LoadConfig()
	return &Config{AppConfig: appConfig}
}

func (c *AppConfig) LoadConfig() {
	env, found := os.LookupEnv("ACTIVE_ENV")
	if !found {
		env = "local"
	}
	log.Println("Loading config for environment: ", env)

	viper.SetTypeByDefaultValue(true)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	readConfigErr := viper.ReadInConfig()

	if readConfigErr != nil {
		panic("Could not read config file")
	}

	sub := viper.Sub(env)

	unMarshallErr := sub.Unmarshal(c)

	if unMarshallErr != nil {
		panic("Could not unmarshall config file")

	}
}
