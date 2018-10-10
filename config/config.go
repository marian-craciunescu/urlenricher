package config

import (
	"errors"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

//Config is used to hold properties read from property file
type Config struct {
	ServerPort int          `mapstructure:"server_port"`
	LogLevel   logrus.Level `mapstructure:"log_level"`
	ElkLogging bool         `mapstructure:"use_elk"`
	ApiKey     string       `mapstructure:"api_key"`
	ApiSecret  string       `mapstructure:"api_secret"`
	DataPath   string       `mapstructure:"path"`
}

//ErrAppPropertiesNotFound is used when the property file is missing at the reading time
var ErrAppPropertiesNotFound = errors.New("Could not read from application_properties.json")

func init() {
	pflag.String("profile", "dev", "set active profile")
	pflag.String("use_elk", "false", "use remote ELK logging")
	pflag.String("log_level", "debug", "set log level")
	pflag.String("api_key", "", "set api key for brightcloud")
	pflag.String("api_secret", "", "set api secret for brightcloud")
	pflag.String("path", "datadir", "set path to on-disk saving ")

	pflag.Parse()
}

//ReadConfig read properties from file and maps them to Config struct
func ReadConfig(propertyFile string) (*Config, error) {
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		logger.WithError(err).Error("cannot bind cmdline flags")
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigFile(propertyFile)
	err = viper.ReadInConfig()
	if err != nil {
		logger.Error(ErrAppPropertiesNotFound.Error())
		return nil, ErrAppPropertiesNotFound
	}

	var p Config
	err = viper.Unmarshal(&p)
	if err != nil {
		logger.Error("Could not deserialize properties from application properties file")
		return nil, err
	}

	return &p, nil
}
