package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func NewConfig(path ConfigPath) *viper.Viper {
	return NewConfigHolder(string(path)).getConfig()
}

var ConfigProviders = wire.NewSet(NewConfig)
