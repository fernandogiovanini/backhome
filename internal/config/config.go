package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fernandogiovanini/backhome/internal/logger"
	"github.com/fernandogiovanini/backhome/internal/printer"
	"github.com/fernandogiovanini/backhome/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// LocalPath is public so it can be set by cobra command flags
	// but it should be read from config.GetLocalPath()
	LocalPath string

	localPath     string
	configFile    = "backhome.yaml"
	configuration Config
)

type Config struct {
	Filenames []string `mapstructure:"files"`
	Remote    string   `mapstructure:"remote"`
}

func InitConfig() {
	if err := InitLocalPath(); err != nil {
		printer.Error("Failed to get local path:\n%v", err)
		logger.Fatalf("failed to set local path: %v", err)
	}

	if err := LoadConfig(); err != nil {
		printer.Error("Failed to load config file:\n%v", err)
		logger.Fatalf("failed to load config file: %v", err)
	}
}

func LoadConfig() error {
	configFile := GetConfigFilePath()

	viper.AddConfigPath(filepath.Dir(configFile))

	viper.SetConfigName(strings.TrimSuffix(filepath.Base(configFile), filepath.Ext(configFile)))
	viper.SetConfigType("yaml")

	viper.SetDefault("files", []string{})
	viper.SetDefault("remote", nil)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file %s: %w", viper.ConfigFileUsed(), err)
	}

	if err := viper.UnmarshalExact(&configuration); err != nil {
		return fmt.Errorf("invalid config file %s: %w", viper.ConfigFileUsed(), err)
	}

	logger.Info("files found in config file: %d", len(configuration.Filenames))

	return nil
}

func ConfigExists() bool {
	file, err := os.Open(GetConfigFilePath())
	if err != nil {
		return false
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	if stat.IsDir() {
		return false
	}

	return true
}

func DefaultLocal() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return strings.Join([]string{home, ".backhome", ""}, string(os.PathSeparator))
}

func GetConfigFilePath() string {
	localPath, err := GetLocalPath()
	if err != nil {
		printer.Error("Failed to get local path:\n%v", err)
		logger.Fatalf("failed to get local path: %v", err)
	}

	path, err := filepath.Abs(
		strings.Join([]string{localPath, configFile}, string(os.PathSeparator)),
	)
	if err != nil {
		printer.Error("Failed to resolve config file path: %v", err)
		logger.Fatalf("failed to resolve config file path %s: %v", path, err)
	}
	return path
}

func GetFilenames() []string {
	if configuration.Filenames == nil {
		return []string{}
	}
	return configuration.Filenames
}

func GetRemote() string {
	if configuration.Remote == "" {
		return ""
	}
	return configuration.Remote
}

func InitLocalPath() error {
	if LocalPath == "" {
		return errors.New("local path cannot be empty")
	}
	local, err := utils.ResolvePath(LocalPath)
	if err != nil {
		return fmt.Errorf("failed to resolve local path %s: %w", local, err)
	}
	localPath = local
	return nil
}

func GetLocalPath() (string, error) {
	if localPath == "" {
		return "", errors.New("local path is not set, call InitLocalPath() first")
	}
	return localPath, nil
}

func AddFile(filename string) error {
	if filename == "" {
		return errors.New("Filename cannot be empty")
	}

	path, err := utils.ResolvePath(filename)
	if err != nil {
		return fmt.Errorf("failed to resolve %s: %w", filename, err)
	}

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file %s does not exist", path)
		}
		if os.IsPermission(err) {
			return fmt.Errorf("permission denied for %s", path)
		}
		return fmt.Errorf("failed to open %s: %w", path, err)
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to read %s stats: %w", path, err)
	}
	if fileinfo.IsDir() {
		return fmt.Errorf("file %s is a directory", path)
	}

	for _, f := range configuration.Filenames {
		if f == filename {
			return fmt.Errorf("file %s already exists in config", filename)
		}
	}

	configuration.Filenames = append(configuration.Filenames, filename)
	viper.Set("files", configuration.Filenames)

	return nil
}

func Save() error {
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write config file %s: %w", viper.ConfigFileUsed(), err)
	}
	return nil
}
