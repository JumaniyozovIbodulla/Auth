package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresPassword string
	PostgresUser     string
	PostgresDatabase string
	ServiceName      string
	RedisHost        string
	RedisPort        string
	RedisPassword    string
	GinMode          string
	ProdPort         string
	ProdHost         string
}

func Load() Config {

	if err := godotenv.Load(); err != nil {
		fmt.Println("no .env file: ", err)
	}

	cfg := Config{}

	cfg.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "sun"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "golang"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "golang"))
	cfg.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost"))
	cfg.RedisPort = cast.ToString(getOrReturnDefault("REDIS_PORT", "6379"))
	cfg.RedisPassword = cast.ToString(getOrReturnDefault("REDIS_PASSWORD", ""))
	cfg.ServiceName = cast.ToString(getOrReturnDefault("SERVICE_NAME", ""))
	cfg.GinMode = cast.ToString(getOrReturnDefault("GIN_MODE", ""))
	cfg.ProdPort = cast.ToString(getOrReturnDefault("PORT", "3000"))
	cfg.ProdHost = cast.ToString(getOrReturnDefault("HOST", ""))

	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {

	if os.Getenv(key) == "" {
		return defaultValue
	}
	return os.Getenv(key)
}
