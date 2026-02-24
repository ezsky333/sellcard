package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
	JWT struct {
		Secret     string `yaml:"secret"`
		TTLMinutes int    `yaml:"ttl_minutes"`
	} `yaml:"jwt"`
	Swagger struct {
		Enabled bool `yaml:"enabled"`
	} `yaml:"swagger"`
	Turnstile struct {
		Enabled   bool   `yaml:"enabled"`
		SiteKey   string `yaml:"site_key"`
		SecretKey string `yaml:"secret_key"`
	} `yaml:"turnstile"`
}

func LoadConfig(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}
	// 应用环境变量覆盖
	cfg.applyEnvOverrides()
	return &cfg, nil
}

// applyEnvOverrides 使用环境变量覆盖配置值
func (c *Config) applyEnvOverrides() {
	// Server
	if host := os.Getenv("SERVER_HOST"); host != "" {
		c.Server.Host = host
	}
	if port := os.Getenv("SERVER_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			c.Server.Port = p
		}
	}

	// Database
	if host := os.Getenv("DATABASE_HOST"); host != "" {
		c.Database.Host = host
	}
	if port := os.Getenv("DATABASE_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			c.Database.Port = p
		}
	}
	if user := os.Getenv("DATABASE_USER"); user != "" {
		c.Database.User = user
	}
	if password := os.Getenv("DATABASE_PASSWORD"); password != "" {
		c.Database.Password = password
	}
	if dbname := os.Getenv("DATABASE_NAME"); dbname != "" {
		c.Database.DBName = dbname
	}

	// JWT
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		c.JWT.Secret = secret
	}
	if ttl := os.Getenv("JWT_TTL_MINUTES"); ttl != "" {
		if t, err := strconv.Atoi(ttl); err == nil {
			c.JWT.TTLMinutes = t
		}
	}

	// Swagger
	if enabled := os.Getenv("SWAGGER_ENABLED"); enabled != "" {
		c.Swagger.Enabled = enabled == "true" || enabled == "1"
	}

	// Turnstile
	if enabled := os.Getenv("TURNSTILE_ENABLED"); enabled != "" {
		c.Turnstile.Enabled = enabled == "true" || enabled == "1"
	}
	if siteKey := os.Getenv("TURNSTILE_SITE_KEY"); siteKey != "" {
		c.Turnstile.SiteKey = siteKey
	}
	if secretKey := os.Getenv("TURNSTILE_SECRET_KEY"); secretKey != "" {
		c.Turnstile.SecretKey = secretKey
	}
}

func (c *Config) DatabaseDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.DBName)
}

func (c *Config) ServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}
