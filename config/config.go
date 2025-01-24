package config

import (
	"fmt"
	"os"
)

type Config struct {
	PostgresConfig
	ServerConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Db       string
}

type ServerConfig struct {
	Port string
	Env  string
}

func (cfg *PostgresConfig) ToMap() map[string]string {
	return map[string]string{
		"HOST":     cfg.Host,
		"PORT":     cfg.Port,
		"USER":     cfg.User,
		"PASSWORD": cfg.Password,
		"DB":       cfg.Db,
	}
}

func (cfg *ServerConfig) ToMap() map[string]string {
	return map[string]string{
		"PORT": cfg.Port,
		"ENV":  cfg.Env,
	}
}

func NewConfig() (*Config, error) {
	postgresConf := PostgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Db:       os.Getenv("POSTGRES_DB"),
	}

	serverConf := ServerConfig{
		Port: os.Getenv("PORT"),
		Env:  os.Getenv("ENV"),
	}

	cfg := &Config{
		PostgresConfig: postgresConf,
		ServerConfig:   serverConf,
	}

	if err := validateEnvs(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func validateEnvs(cfg *Config) error {
	missingEnvs := []string{}

	missingEnvs = append(missingEnvs, findEmptyValueKeys(cfg.PostgresConfig.ToMap())...)
	missingEnvs = append(missingEnvs, findEmptyValueKeys(cfg.ServerConfig.ToMap())...)

	if len(missingEnvs) > 0 {
		return fmt.Errorf("missing envs: %v", missingEnvs)
	}
	return nil
}

func findEmptyValueKeys(m map[string]string) []string {
	keys := []string{}
	for k, v := range m {
		if v == "" {
			keys = append(keys, k)
		}
	}
	return keys
}
