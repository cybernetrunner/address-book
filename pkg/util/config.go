package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerHost     string `mapstructure:"SERVER_HOST"`
	ServerGRPCPort string `mapstructure:"SERVER_GRPC_PORT"`
	ServerHTTPPort string `mapstructure:"SERVER_HTTP_PORT"`

	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBSSLMode  string `mapstructure:"DB_SSLMODE"`
	DBTimezone string `mapstructure:"DB_TIMEZONE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
