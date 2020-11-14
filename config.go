package cachedservice

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	cacheEnabled = "cache.enabled"
	redisUrl     = "redis.url"
	projectId    = "project.id"
	credentials  = "project.credentials"
)

type ConfigPath string

type ConfigHolder struct {
	config *viper.Viper
}

func NewConfig(path ConfigPath) *viper.Viper {
	return newConfigHolder(string(path)).getConfig()
}

func newConfigHolder(path string) *ConfigHolder {
	if path == "" {
		log.Fatalln("path is empty")
	}
	if !fileExists(path) {
		log.Fatalln("file doesn't exist", path)
	}
	v := viper.New()
	v.SetConfigType("json")
	v.SetConfigFile(path)
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln("error reading config", path)
	}
	return &ConfigHolder{config: v}
}

func (cfg ConfigHolder) getConfig() *viper.Viper {
	return cfg.config
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
