package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	dbUser        string `mapstructure:"DB_USER"`
	dbHost        string `mapstructure:"DB_HOST`
	dbDatabase    string `mapstructure:"DB_DATABASE"`
	dbPassword    string `mapstructure:"DB_PASSWORD"`
	dbPort        string `mapstructure:"DB_PORT"`
	maxOpenDbConn string `mapstructure:"MAX_OPEN_DB_CONN"`
	maxIdleDbConn string `mapstructure:"MAX_IDLE_DB_CONN"`
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
