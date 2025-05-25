package config

import (
	"fmt"
	"net/url"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Port       string
	Host       string
	PathPrefix string
	SecretKey  string
}

type DBConfig struct {
	User     string
	Name     string
	Password string
	Port     int
	Host     string
	PoolSize int
}

func (d *DBConfig) DSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		url.QueryEscape(d.User),
		url.QueryEscape(d.Password),
		d.Host,
		d.Port,
		d.Name,
	)
}

type Config struct {
	App  AppConfig
	DB   DBConfig
	Base BaseConfig
}

func Load() (*Config, error) {
	base := BaseConfig{}
	if err := base.Load(); err != nil {
		return nil, err
	}

	return &Config{
		Base: base,
		App: AppConfig{
			Port:       viper.GetString("APP_PORT"),
			Host:       viper.GetString("APP_HOST"),
			PathPrefix: viper.GetString("PATH_PREFIX"),
			SecretKey:  viper.GetString("SECRET_KEY"),
		},
		
		DB: DBConfig{
			User:     viper.GetString("DATABASE_USERNAME"),
			Name:     viper.GetString("DATABASE_NAME"),
			Password: viper.GetString("DATABASE_PASSWORD"),
			Port:     viper.GetInt("DATABASE_PORT"),
			Host:     viper.GetString("DATABASE_HOST"),
		},
	}, nil
}