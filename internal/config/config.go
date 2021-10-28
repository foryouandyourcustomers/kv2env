package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// parsed runtime config
	Cfg Config

	// defaults for configuration values
	defLogLevel      = log.ErrorLevel // debug = 5, warning = 3, error = 2
	defKeyVault      = ""
	defFilterPrefix  = ""
	defPrefixRemove  = false
	defPrefixReplace = ""
	defPrefixTfVar   = true
	defPassParentEnv = true
)

type Config struct {
	LogLevel      log.Level `mapstructure:"LOGLEVEL"`
	KeyVault      string    `mapstructure:"KEYVAULT"`
	FilterPrefix  string    `mapstructure:"FILTERPREFIX"`
	PrefixRemove  bool      `mapstructure:"PREFIXREMOVE"`
	PrefixReplace string    `mapstructure:"PREFIXREPLACE"`
	PrefixTfVar   bool      `mapstructure:"PREFIXTFVAR"`
	PassParentEnv bool      `mapstructure:"PASSPARENTENV"`
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
	viper.SetDefault("LogLevel", defLogLevel)
	viper.SetDefault("KeyVault", defKeyVault)
	viper.SetDefault("FilterPrefix", defFilterPrefix)
	viper.SetDefault("PrefixRemove", defPrefixRemove)
	viper.SetDefault("PrefixReplace", defPrefixReplace)
	viper.SetDefault("PrefixTfVar", defPrefixTfVar)
	viper.SetDefault("PassParentEnv", defPassParentEnv)

	// load environment
	viper.AutomaticEnv()
	err = viper.Unmarshal(&config)
	return
}
