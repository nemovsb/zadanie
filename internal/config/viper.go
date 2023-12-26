package config

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

const (
	ENV_PROD = "Prod"
	ENV_DEV  = "Dev"
)

var ErrUnmarshalConfig = errors.New("viper failed to unmarshal app config")

func ViperConfigurationProvider(env string, writeConfig bool) (cfg *Config, err error) {
	var filename string

	switch env {
	case ENV_PROD:
		filename = "config.prod"
	case ENV_DEV:
		filename = "config.local"
	default:
		filename = "config"
	}

	v := NewViper(filename)

	cfg, err = NewConfiguration(v)
	if err != nil {
		return
	}

	if writeConfig {
		if err = v.WriteConfig(); err != nil {
			log.Println("viper failed to write app config file:", err)
		}
	}

	return cfg, nil
}

func NewViper(filename string) *viper.Viper {
	v := viper.New()

	if filename != "" {
		v.SetConfigName(filename)
		v.AddConfigPath(".")
		v.AddConfigPath(filepath.FromSlash("./config/"))
	}

	// Some defaults will be set just so they are accessible via environment variables
	// (basically so viper knows they exist)

	v.SetDefault("Server.Port", "8085")

	v.SetConfigType("yaml")

	// Set environment variable support:
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix("MYAPP")
	v.AutomaticEnv()

	// ReadInConfig will discover and load the configuration file from disk
	// and key/value stores, searching in one of the defined paths.
	if err := v.ReadInConfig(); err != nil {
		log.Println("viper failed to read app config file:", err)
	}

	return v
}

func NewConfiguration(v *viper.Viper) (*Config, error) {
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrUnmarshalConfig, err)
	}

	return &c, nil
}

func GetConfig() (conf Config, err error) {

	var configPath string
	flag.StringVar(&configPath, "cfgPath", "", "path to file")

	flag.Parse()
	if !flag.Parsed() {
		log.Fatal("Flag not parsed")
	}

	binFile, err := os.ReadFile(configPath)
	if err != nil {
		return conf, err
	}

	switch strings.ToLower(path.Ext(configPath)) {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(binFile, &conf)
	case ".json":
		err = json.Unmarshal(binFile, &conf)
	default:
		return conf, errors.New("unknown config format")
	}

	return conf, err
}
