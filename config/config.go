package config

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Server *Server `mapstructure:"server" validate:"required"`
}

type Server struct {
	Port         int           `mapstructure:"port" validate:"required"`
	AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
	Timeout      time.Duration `mapstructure:"timeout" validate:"required"`
	BodyLimit    string        `mapstructure:"bodyLimit" validate:"required"`
}

var (
	once           sync.Once
	configInstance *Config
)

func ConfigGetting() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
		validate := validator.New()
		if err := validate.Struct(configInstance); err != nil {
			panic(err)
		}
	})
	fmt.Println("configInstance was loaded successfully")
	return configInstance
}
