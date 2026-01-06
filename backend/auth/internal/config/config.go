package config

import (
	"github.com/joho/godotenv"
	"github.com/numbergroup/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	ServerConfig ServerConfig `yaml:"server"`
	GRPCConfig   GRPCConfig   `yaml:"grpc"`
}

type ServerConfig struct {
	Port            string        `yaml:"port"`
	Timeout         time.Duration `yaml:"timeout"`
	SessTTl         time.Duration `yaml:"sess_ttl"`
	AbsoluteSessTTl time.Duration `yaml:"absolute_sess_ttl"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

type GRPCConfig struct {
	Host    string        `yaml:"host"`
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoadConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		return getDefaultConfig()
	}

	path := os.Getenv("PATH_CONFIG")
	if len(path) == 0 {
		return getDefaultConfig()
	}

	var config Config
	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("Ошибка загрузки конфига")
	}

	return &config
}

func getDefaultConfig() *Config {
	log.Println("\n\n\n\nИспользуется дефолтный кофиг, установите корректный путь до yaml/.env файла\n\n\n\n")
	return &Config{
		ServerConfig: ServerConfig{
			Port:            "8080",
			Timeout:         2 * time.Minute,
			SessTTl:         30 * time.Minute,
			AbsoluteSessTTl: 24 * time.Hour,
			ShutdownTimeout: 2 * time.Second,
		},
		GRPCConfig: GRPCConfig{
			Host:    "0.0.0.0",
			Port:    "50051",
			Timeout: 2 * time.Minute,
		},
	}
}
