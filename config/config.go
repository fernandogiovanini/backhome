package config

import (
	"os"

	"github.com/fernandogiovanini/backhome/logger"
	"github.com/fernandogiovanini/backhome/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Configuration Config
	ConfigFile    string
)

type Config struct {
	BackupItems []string `mapstructure:"targets"`
	Local       string   `mapstructure:"local"`
	Remote      string   `mapstructure:"remote"`
}

func InitConfig() {
	if ConfigFile != "" {
		logger.Debug("config file provided: %s", ConfigFile)
		configPath, err := utils.ResolvePath(ConfigFile)
		if err != nil {
			logger.Fatalf("failed to load config file %s: %v", viper.ConfigFileUsed(), err)
		}
		viper.SetConfigFile(configPath)
	} else {
		logger.Debug("config file default")
		viper.AddConfigPath(DefaultConfigPath())
		viper.SetConfigName(".backhome")
		viper.SetConfigType("yaml")
	}

	logger.Info("loading config file %s", viper.ConfigFileUsed())
	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf("failed to load config file %s: %v", viper.ConfigFileUsed(), err)
	}

	if err := viper.Unmarshal(&Configuration); err != nil {
		logger.Fatalf("invalid config file %s: %v", viper.ConfigFileUsed(), err)
	}
}

func DefaultConfigPath() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return home
}
