package config

import "github.com/sahalazain/go-common/config"

var Map = map[string]interface{}{
	"APP_PORT":       "3030",
	"APP_PREFIX":     "",
	"APP_SECRET":     "",
	"REDIS_HOST":     "localhost:6379",
	"REDIS_PASSWORD": "",
	"REDIS_DB":       0,
	"MYSQL_USER":     "",
	"MYSQL_PASS":     "",
	"MYSQL_HOST":     "",
	"MYSQL_PORT":     "",
	"MYSQL_DBNAME":   "",
	"PG_USER":        "",
	"PG_PASS":        "",
	"PG_HOST":        "",
	"PG_PORT":        "",
	"PG_DBNAME":      "",
}

var Cfg config.Getter
var ConfigUrl string

func Load() error {
	cfgClient, err := config.Load(Map, ConfigUrl)
	if err != nil {
		return err
	}

	Cfg = cfgClient

	return nil
}
