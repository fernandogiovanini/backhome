package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fernandogiovanini/backhome/internal/backhome"
	"github.com/fernandogiovanini/backhome/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultConfigFilename = "backhome.yaml"

var (
	// LocalPath is public so it can be set by cobra command flags
	// but it should be read from config.GetLocalPath()
	LocalPath string
)

type Config struct {
	Filenames []string `mapstructure:"files"`
	Remote    string   `mapstructure:"remote"`

	localPath  string
	configFile string
}

func InitConfig() (*Config, error) {
	config := &Config{}

	if err := config.initLocalPath(LocalPath); err != nil {
		return nil, fmt.Errorf("failed to set local path: %w", err)
	}

	if err := config.initConfigPath(defaultConfigFilename); err != nil {
		return nil, fmt.Errorf("failed to set local path: %w", err)
	}

	viper.AddConfigPath(filepath.Dir(config.configFile))

	viper.SetConfigName(strings.TrimSuffix(filepath.Base(config.configFile), filepath.Ext(config.configFile)))
	viper.SetConfigType("yaml")

	viper.SetDefault("files", []string{})
	viper.SetDefault("remote", nil)

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := viper.UnmarshalExact(&config); err != nil {
		return nil, fmt.Errorf("invalid config file %s: %w", viper.ConfigFileUsed(), err)
	}

	return config, nil
}

func (c *Config) initLocalPath(localPath string) error {
	if localPath == "" {
		return errors.New("local path cannot be empty")
	}

	localPath, err := utils.ResolvePath(localPath)
	if err != nil {
		return fmt.Errorf("failed to resolve local path %s: %w", localPath, err)
	}
	c.localPath = localPath

	return nil
}

func (c *Config) initConfigPath(configFilename string) error {
	configPath, err := filepath.Abs(
		strings.Join([]string{c.localPath, configFilename}, string(os.PathSeparator)),
	)
	if err != nil {
		return fmt.Errorf("failed to resolve config file path: %w", err)
	}
	c.configFile = configPath

	return nil
}

func (c Config) GetConfigFilePath() string {
	return c.configFile
}

func (c Config) GetFilenames() []string {
	if c.Filenames == nil {
		return []string{}
	}
	return c.Filenames
}

func (c Config) GetRemote() string {
	if c.Remote == "" {
		return ""
	}
	return c.Remote
}

func (c Config) GetLocalPath() (string, error) {
	if c.localPath == "" {
		return "", errors.New("local path is not set, call InitLocalPath() first")
	}
	return c.localPath, nil
}

func (c *Config) AddFile(filename string) error {
	if filename == "" {
		return errors.New("ilename cannot be empty")
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

	for _, f := range c.Filenames {
		if f == filename {
			return fmt.Errorf("file %s already exists in config", filename)
		}
	}

	c.Filenames = append(c.Filenames, filename)
	viper.Set("files", c.Filenames)

	return nil
}

func (c Config) Save() error {
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write config file %s: %w", viper.ConfigFileUsed(), err)
	}
	return nil
}

func (c Config) MakeLocalRepository() error {
	if _, err := backhome.MakeLocal(c.localPath); err != nil {
		return fmt.Errorf("failed to create local repository %s: %w", c.localPath, err)
	}

	return nil
}

func (c Config) CreateConfigFile() error {
	file, err := os.OpenFile(c.configFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("failed to create config file %s: %w", c.configFile, err)
	}
	defer file.Close()

	return nil
}

func DefaultLocal() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return strings.Join([]string{home, ".backhome", ""}, string(os.PathSeparator))
}
