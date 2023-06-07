package config

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

var _ ConfigurationGetter = (*YamlConfigurationGetter) (nil)

type Config struct {
	Http struct {
		Port string `mapstructure:"port"`
		BasicAuth map[string] string `mapstructure:"basicAuth"`
	} `mapstructure:"http"`
	Logging struct {
		Level int `mapstructure:"level"`
	} `mapstructure:"logging"`
}

type ConfigurationGetter interface {
	GetConfig() (Config, error)
}

type YamlConfigurationGetter struct {
}

func NewYamlConfigurationGetter() *YamlConfigurationGetter {
	return &YamlConfigurationGetter{}
}

func (y *YamlConfigurationGetter) GetConfig() (Config, error) {
	conf := Config{}
	fileName, err:= filepath.Abs("./config/config.yaml")
	if err != nil {
		logrus.WithError(err).Fatal("Failed File path")
		return conf, err
	}
	vpr, err := ConfigureViper(fileName, "app", "yaml")
	if err != nil {
		logrus.WithError(err).Fatal("Failed Config viper")
		return conf, err
	}
	vpr.Unmarshal(&conf)
	return conf, nil
}

func ConfigureViper(path, envPrefix, configType string) (*viper.Viper, error) {
	vpr := viper.New()
	defaultConfig := bytes.NewReader([]byte{})
	vpr.SetConfigType("yaml")
	if err := vpr.MergeConfig(defaultConfig); err != nil {
		return nil, err
	}

	// Override config
	vpr.SetConfigFile(path)
	vpr.SetConfigType(configType)
	if err := vpr.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigParseError); ok {
			return nil, err
		}
		// Dont return error if file is missing as overwrite is optional
	}

	// Override env variables
	vpr.AutomaticEnv()
	vpr.SetEnvPrefix(envPrefix)
	vpr.AddConfigPath(".")
	vpr.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Manually set all config values, only needed when using viper.Unmarshal()
	for _, key := range vpr.AllKeys() {
		val := vpr.Get(key)
		vpr.Set(key, val)
	}

	return vpr, nil
}