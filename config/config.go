package config

import "github.com/spf13/viper"

type Config struct {
	Port         string `mapstructure:"PORT"`
	Host         string `mapstructure:"HOST"`
	BasePath     string `mapstructure:"BASE_PATH"`
	SaltUrl      string `mapstructure:"SALT_URL"`
	DbName       string `mapstructure:"DB_NAME"`
	DbCollection string `mapstructure:"DB_COLLECTION"`
	DbUri        string `mapstructure:"DB_URI"`
}

func NewConfig() (config Config, err error) {
	viper.SetConfigFile("./.env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
