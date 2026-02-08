package config

import (
	"github.com/joho/godotenv"
	"github.com/numbergroup/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env  string
	HTTP HTTPConfig `yaml:"http"`
	GRPC GRPCConfig `yaml:"grpc"`
}

type HTTPConfig struct {
	Port            string        `yaml:"port"`
	Timeout         time.Duration `yaml:"timeout"`
	AbsoluteSessTTl time.Duration `yaml:"absolute_sess_ttl"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

type GRPCConfig struct {
	Host    string        `yaml:"host"`
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	var cfg Config

	defer func() *Config {
		if err := recover(); err != nil {
			return getDefaultConfig()
		}
		return &cfg
	}()

	godotenv.Load(".env")
	path := os.Getenv("CONFIG_PATH")
	cleanenv.ReadConfig(path, &cfg)

	cfg.Env = getEnv()
	return &cfg
}

func getDefaultConfig() *Config {
	log.Println("Используется дефолтный конфиг")

	return &Config{
		Env: getEnv(),
		HTTP: HTTPConfig{
			Port:            "8120",
			Timeout:         5 * time.Second,
			ShutdownTimeout: 2 * time.Second,
		},
		GRPC: GRPCConfig{
			Host:    "localhost",
			Port:    "50051",
			Timeout: 4 * time.Second,
		},
	}
}

func getEnv() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "prod"
	}
	return env
}
