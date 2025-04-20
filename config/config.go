package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	PostgresConfig
	ServerConfig
	SchedulerConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

type ServerConfig struct {
	Port string
	Env  string
}

type SchedulerConfig struct {
	Enabled             bool
	RssSyncIntervalCron string
}

func (cfg *PostgresConfig) ToMap() map[string]string {
	return map[string]string{
		"HOST":     cfg.Host,
		"PORT":     cfg.Port,
		"USER":     cfg.User,
		"PASSWORD": cfg.Password,
		"DB":       cfg.DB,
	}
}

func (cfg *ServerConfig) ToMap() map[string]string {
	return map[string]string{
		"PORT": cfg.Port,
		"ENV":  cfg.Env,
	}
}

func (cfg *SchedulerConfig) ToMap() map[string]string {
	return map[string]string{
		"ENABLED":                strconv.FormatBool(cfg.Enabled),
		"RSS_SYNC_INTERVAL_CRON": cfg.RssSyncIntervalCron,
	}
}

func NewConfig() (*Config, error) {
	postgresConf := PostgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DB:       os.Getenv("POSTGRES_DB"),
	}

	serverConf := ServerConfig{
		Port: os.Getenv("PORT"),
		Env:  os.Getenv("ENV"),
	}

	// Default scheduler configuration
	schedulerEnabled := true
	if enabled, err := strconv.ParseBool(os.Getenv("SCHEDULER_ENABLED")); err == nil {
		schedulerEnabled = enabled
	}

	// Default to every hour if not specified
	rssSyncCron := "0 0 * * * *"
	if cronExpr := os.Getenv("RSS_SYNC_INTERVAL_CRON"); cronExpr != "" {
		rssSyncCron = cronExpr
	}

	schedulerConf := SchedulerConfig{
		Enabled:             schedulerEnabled,
		RssSyncIntervalCron: rssSyncCron,
	}

	cfg := &Config{
		PostgresConfig:  postgresConf,
		ServerConfig:    serverConf,
		SchedulerConfig: schedulerConf,
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
	// Scheduler has defaults, so we don't check for missing values

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
