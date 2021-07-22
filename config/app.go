package config

import "github.com/tkanos/gonfig"

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Config {
	conf := Config{}
	gonfig.GetConf("config/app.json", &conf)

	return conf
}
