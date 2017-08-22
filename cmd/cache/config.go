package main

import (
	"github.com/moira-alert/moira-alert/cmd"
)

type config struct {
	Redis    cmd.RedisConfig    `yaml:"redis"`
	Graphite cmd.GraphiteConfig `yaml:"graphite"`
	Logger   cmd.LoggerConfig   `yaml:"log"`
	Cache    cacheConfig        `yaml:"cache"`
}

type cacheConfig struct {
	Listen          string `yaml:"listen"`
	RetentionConfig string `yaml:"retention-config"`
}

func getDefault() config {
	return config{
		Redis: cmd.RedisConfig{
			Host: "localhost",
			Port: "6379",
			DBID: 0,
		},
		Logger: cmd.LoggerConfig{
			LogFile:  "stdout",
			LogLevel: "debug",
		},
		Cache: cacheConfig{
			Listen:          ":2003",
			RetentionConfig: "storage-schemas.conf",
		},
		Graphite: cmd.GraphiteConfig{
			URI:      "localhost:2003",
			Prefix:   "DevOps.Moira",
			Interval: "60s0ms",
		},
	}
}
