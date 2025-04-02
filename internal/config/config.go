//go:generate mockery --all --case snake

package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"slices"

	"github.com/fernandogiovanini/backhome/internal/filesystem"
	"github.com/fernandogiovanini/backhome/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const DefaultConfigFilename = "backhome.yaml"

var (
	// LocalPath is public so it can be set by cobra command flags
	// but it should be read from config.GetLocalPath()
	LocalPath string
)

type ConfigStorage interface {
	AddFile(filename string) error
	CreateConfigFile() error
	GetConfig() *ConfigData
	MakeLocalRepository() error
	Save() error
}

type configStorage struct {
	fs  filesystem.FileSystem
	cfg *ConfigData
	v   *viper.Viper
}

type Config interface {
	GetConfigFilePath() string
	GetFilenames() []string
	GetRemote() string
	GetLocalPath() (string, error)
}

type ConfigData struct {
	Filenames []string `mapstructure:"files"`
	Remote    string   `mapstructure:"remote"`

	localPath  string
	configFile string
}

func NewConfigStorage(
	localPath string,
	cfgFilename string,
	fs filesystem.FileSystem,
	v *viper.Viper) (*configStorage, error) {
	cfg, err := NewConfig(LocalPath, cfgFilename)
	if err != nil {
		return nil, err
	}

	configStorage := &configStorage{
		cfg: cfg,
		fs:  fs,
		v:   v,
	}

	v.AddConfigPath(filepath.Dir(cfg.GetConfigFilePath()))

	v.SetConfigName(strings.TrimSuffix(filepath.Base(cfg.GetConfigFilePath()), filepath.Ext(cfg.GetConfigFilePath())))
	v.SetConfigType("yaml")

	v.SetDefault("files", []string{})
	v.SetDefault("remote", nil)

	if err := v.ReadInConfig(); err != nil {
		return configStorage, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := v.UnmarshalExact(&cfg); err != nil {
		return nil, fmt.Errorf("invalid config file %s: %w", v.ConfigFileUsed(), err)
	}

	return configStorage, nil
}

func (cs *configStorage) AddFile(filename string) error {
	if filename == "" {
		return errors.New("ilename cannot be empty")
	}

	path, err := utils.ResolvePath(filename)
	if err != nil {
		return fmt.Errorf("failed to resolve %s: %w", filename, err)
	}

	file, err := cs.fs.Open(path)
	if err != nil {
		if cs.fs.IsNotExist(err) {
			return fmt.Errorf("file %s does not exist", path)
		}
		if cs.fs.IsPermission(err) {
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

	if slices.Contains(cs.cfg.GetFilenames(), filename) {
		return fmt.Errorf("file %s already exists in config", filename)
	}

	filenames := append(cs.cfg.GetFilenames(), filename)
	cs.v.Set("files", filenames)

	return nil
}

func (cs *configStorage) GetConfig() *ConfigData {
	return cs.cfg
}

func (cs configStorage) Save() error {
	if err := cs.v.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write config file %s: %w", viper.ConfigFileUsed(), err)
	}
	return nil
}

func (cs configStorage) MakeLocalRepository() error {
	localPath, err := cs.cfg.GetLocalPath()
	if err != nil {
		return fmt.Errorf("failed to get local path %s: %w", localPath, err)
	}
	if err := cs.fs.MkdirAll(localPath, 0755); err != nil {
		return fmt.Errorf("failed to create local directory %s: %w", localPath, err)
	}

	return nil
}

func (cs configStorage) CreateConfigFile() error {
	file, err := cs.fs.OpenFile(cs.cfg.GetConfigFilePath(), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("failed to create config file %s: %w", cs.cfg.configFile, err)
	}
	defer file.Close()

	return nil
}

func NewConfig(localPath string, configFilename string) (*ConfigData, error) {
	config := &ConfigData{}
	if err := config.initLocalPath(localPath); err != nil {
		return nil, fmt.Errorf("failed to set local path: %w", err)
	}

	if err := config.initConfigPath(DefaultConfigFilename); err != nil {
		return nil, fmt.Errorf("failed to set local path: %w", err)
	}

	return config, nil
}

func (c *ConfigData) initLocalPath(localPath string) error {
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

func (c *ConfigData) initConfigPath(configFilename string) error {
	configPath, err := filepath.Abs(
		strings.Join([]string{c.localPath, configFilename}, string(os.PathSeparator)),
	)
	if err != nil {
		return fmt.Errorf("failed to resolve config file path: %w", err)
	}
	c.configFile = configPath

	return nil
}

func (c ConfigData) GetConfigFilePath() string {
	return c.configFile
}

func (c ConfigData) GetFilenames() []string {
	if c.Filenames == nil {
		return []string{}
	}
	return c.Filenames
}

func (c ConfigData) GetRemote() string {
	if c.Remote == "" {
		return ""
	}
	return c.Remote
}

func (c ConfigData) GetLocalPath() (string, error) {
	if c.localPath == "" {
		return "", errors.New("local path is not set, call InitLocalPath() first")
	}
	return c.localPath, nil
}

func DefaultLocal() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return strings.Join([]string{home, ".backhome", ""}, string(os.PathSeparator))
}
