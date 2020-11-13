package config

import (
	"github.com/spf13/viper"
	"log"
	"cachedservice/pkg/util"
)
type ConfigPath string
type ConfigHolder struct {
	config *viper.Viper
}

func NewConfigHolder(path string) *ConfigHolder {
	if path == "" {
		log.Fatalln("path is empty")
	}
	if !util.FileExists(path) {
		log.Fatalln("file doesn't exist", path)
	}
	v := viper.New()
	v.SetConfigType("json")
	v.SetConfigFile(path)
	err:=v.ReadInConfig()
	if err!= nil{
		log.Fatalln("error reading config", path)
	}
	return &ConfigHolder{config: v}
}

func (cfg ConfigHolder) getConfig() *viper.Viper {
	return cfg.config
}
