package config

import (
	"embed"
	_ "embed"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

//go:embed app.json
//go:embed app-test.json
//go:embed app-debug.json
var f embed.FS

type Config struct {
	AppName     string         `json:"app_name"`
	AppHost     string         `json:"app_host"`
	AppPort     string         `json:"app_port"`
	JWTKey      string         `json:"jwt_key"`
	Database    DatabaseConfig `json:"database"`
	RedisConfig RedisConfig    `json:"redis_config"`
	SMTPConfig  SMTPConfig     `json:"smtp_config"`
	MailConfig  MailConfig     `json:"mail_config"`
	FileConfig  FileConfig     `json:"file_config"`
}

type DatabaseConfig struct {
	DSN string `json:"dsn"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

type SMTPConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	SSLPort  string `json:"ssl_port"`
	Username string `json:"username"`
	Password string `json:"password"`
	SSL      bool   `json:"ssl"`
}
type MailConfig struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Url     string `json:"url"`
}
type FileConfig struct {
	Prefix string `json:"prefix"`
	Avatar string `json:"avatar"`
	Photo  string `json:"photo"`
	Editor string `json:"editor"`
}

func GetConfig() *Config {
	return cfg
}

var cfg *Config = nil

func ParseConfig() (*Config, error) {
	configPath := "app-debug.json"
	if gin.Mode() == gin.TestMode {
		configPath = "app-test.json"
	}
	if gin.Mode() == gin.ReleaseMode {
		configPath = "app.json"
	}
	data, _ := f.ReadFile(configPath)
	err := json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
