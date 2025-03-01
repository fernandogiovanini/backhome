package config

import (
	"github.com/fernandogiovanini/backhome/logger"
	"github.com/fernandogiovanini/backhome/utils"
	"github.com/spf13/viper"
)

var (
	Configuration Config
	ConfigFile    string
)

type Config struct {
	Targets []string `mapstructure:"targets"`
	Local   string   `mapstructure:"local"`
	Remote  string   `mapstructure:"remote"`
}

func InitConfig() {
	if ConfigFile != "" {
		logger.Debug("config file provided", logger.Args("configFile", ConfigFile))
		configPath, err := utils.ResolvePath(ConfigFile)
		if err != nil {
			logger.Fatalf("failed to load config file %s: %s", viper.ConfigFileUsed(), err)
		}
		viper.SetConfigFile(configPath)
	} else {
		logger.Debug("config file default")
		viper.AddConfigPath(defaultConfigPath())
		viper.SetConfigName(".config")
		viper.SetConfigType("yaml")
		logger.Debug("config file used", logger.Args("configFIle", viper.ConfigFileUsed()))
	}

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf("failed to load config file %s: %s", viper.ConfigFileUsed(), err)
	}

	if err := viper.Unmarshal(&Configuration); err != nil {
		logger.Fatalf("invalid config file %s: %s", viper.ConfigFileUsed(), err)
	}
}

func defaultConfigPath() string {
	// if environment == "DEV" {
	return "./"

	// if not DEV
	// home, err := os.UserHomeDir()
	// cobra.CheckErr(err)

	// return home
}
