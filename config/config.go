package config

import "github.com/spf13/viper"

type config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}

func New(path string) *config {

	viper.SetConfigName("app_config")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	cfg := new(config)
	err = viper.Unmarshal(cfg)

	if err != nil {
		panic(err)
	}

	return cfg
}
