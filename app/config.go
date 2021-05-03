package app

import (
	"github.com/batroff/schedule-back/models/config"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func LoadConfig(path string) (*config.AppConfig, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg = &config.AppConfig{}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, errors.Wrap(err, "can not unmarshal config from file to struct")
	}

	return cfg, nil
}
