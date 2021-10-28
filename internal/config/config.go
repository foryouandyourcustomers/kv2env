package config

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// parsed runtime config
	Cfg Config

	// defaults for configuration values
	defLogLevel = log.ErrorLevel // debug = 5, warning = 3, error = 2
)

type Config struct {
	LogLevel log.Level `mapstructure:"LOGLEVEL"`
	KeyVault string    `mapstructure:"KEYVAULT"`
}

func init() {
	// load dotenv
	godotenv.Load(".env")
	// load viper config
	var err error
	Cfg, err = LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
}

func LoadConfig(path string) (config Config, err error) {
	// set defaults
	// parse loglevel
	viper.SetDefault("LogLevel", defLogLevel)
	viper.SetDefault("KeyVault", "")

	viper.AutomaticEnv()
	fmt.Println(viper.AllSettings())
	err = viper.Unmarshal(&config)
	return
}
