package config

import (
	"github.com/spf13/viper"
	"fmt" 
)

type Environment string

const (
	EnvProd Environment = "prod"
	EnvDev  Environment = "dev"
)

type BaseConfig struct {
	env Environment
}

func (c *BaseConfig) Load() error {
    possiblePaths := []string{
        "/app/config/.env",  
        "config/.env",       
        ".env",             
    }
    
    for _, path := range possiblePaths {
        viper.SetConfigFile(path)
        if err := viper.ReadInConfig(); err == nil {
            viper.AutomaticEnv()
            c.env = Environment(viper.GetString("ENV"))
            return nil
        }
    }
    
    return fmt.Errorf("failed to find .env file in any of: %v", possiblePaths)
}

func (c *BaseConfig) Env() Environment {
	return c.env
}