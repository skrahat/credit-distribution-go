package config

import (
	"github.com/spf13/viper"
)

type Config struct {
    Server struct {
        Address string
    }
    Database struct {
        DSN string
    }
}

func LoadConfig() Config {
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()

    var cfg Config
    if err := viper.ReadInConfig(); err != nil {
        panic(err)
    }

    if err := viper.Unmarshal(&cfg); err != nil {
        panic(err)
    }

    return cfg
}
