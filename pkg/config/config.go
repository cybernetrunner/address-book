package config

import (
	"github.com/spf13/viper"
)

var Yaml = []byte(`
server:
  hostname: localhost
  port:
    grpc: 9090
    http: 8081

database:
  user: gorm
  password: gorm12345
  dbname: gorm
  port: 9920
  sslmode: disable
  timezone: Europe/Moscow
`)

type Config = viper.Viper
