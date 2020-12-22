package config

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

// Config configure application parameters
type Config struct {
	RunMode string `yaml:"runMode"`
	HTTP    HTTP   `yaml:"http"`
	Log     Log    `yaml:"log"`
	DB      DB     `yaml:"db"`
	API     API    `yaml:"api"`
}

func (c *Config) String() string {
	ret, _ := json.Marshal(c)
	return string(ret)
}

// NewConfig is a constructor of config with config path
func NewConfig(path string) (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// HTTP configure HTTP server parameters
type HTTP struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// GetAddress combine host and port, format is `host:port`
func (h HTTP) GetAddress() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

// Log configure log parameters
type Log struct {
	Format string `yaml:"format"`
	Level  string `yaml:"level"`
}

// DB is configuration of database
type DB struct {
	URL   string `yaml:"url"`
	Debug bool   `yaml:"debug"`
}

// API configure api settings
type API struct {
	Endpoint string `yaml:"endpoint"`
}
