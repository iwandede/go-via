package config

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-playground/validator"
	"gopkg.in/yaml.v2"
)

/* -------------------------------------------------------------------------- */
/*                           DEFINE CONFIG FROM YAML                          */
/* -------------------------------------------------------------------------- */
type Config struct {
	Server   *Server   `yaml:"server"`
	DB       *Database `yaml:"database"`
	Security *Security `yaml:"security"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int16  `yaml:"port"`
}

type Security struct {
	SecretKey string `yaml:"secret_key"`
}

/* -------------------------------------------------------------------------- */
/*                               DATABASE CONFIG                              */
/* -------------------------------------------------------------------------- */
type Database struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int16  `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"dbname"`
	SSL      string `yaml:"ssl"`
}

func NewConfigFromYAML(src io.Reader) (*Config, error) {
	var conf Config
	buf, err := ioutil.ReadAll(src)

	if err != nil {
		return nil, fmt.Errorf("failed to read data: %s", err)
	}

	if err := yaml.Unmarshal(buf, &conf); err != nil {
		return nil, err
	}

	if err := checkConfig(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func checkConfig(config *Config) error {
	v := validator.New()
	if err := v.Struct(*config); err != nil {
		return fmt.Errorf("config missing required fields: %s", err)
	}
	return nil
}
