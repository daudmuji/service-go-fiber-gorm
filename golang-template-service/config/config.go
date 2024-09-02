package config

import (
	"fmt"
	"log"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Config struct {
	Http     HttpConfig
	Database DatabaseConfig
	Redis    RedisConfig
}

// HttpConfig represents struct for http server-related config
type HttpConfig struct {
	Port string `env:"HTTP_PORT"`
}

type DatabaseConfig struct {
	Host     string `env:"DB_HOST"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DbName   string `env:"DB_NAME"`
	Port     int    `env:"DB_PORT"`
}

type RedisConfig struct {
	RedistHost string `env:"REDIS_HOST"`
	RedisPort  string `env:"REDIS_PORT"`
	RedisPass  string `env:"REDIS_PASS"`
	RedisDb    int    `env:"REDIS_DB"`
}

func Read(filePath string) (cfg Config, err error) {
	// use Overload to prioritize config file over OS' env vars
	if err = godotenv.Load(filePath); err != nil {
		log.Println("error reading file in", filePath, ", defaulting to OS environment variables")
	}

	if err = envdecode.StrictDecode(&cfg); err != nil {
		err = fmt.Errorf("error decoding config: %+v", err)
		return
	}

	cfg.Http.Port = fmt.Sprintf("0.0.0.0:%s", cfg.Http.Port)

	return cfg, nil
}
