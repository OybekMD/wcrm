package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	// User
	UserServiceHost string
	UserServicePort int

	// Post
	PostServiceHost string
	PostServicePort int

	// context timeout in seconds
	CtxTimeout int

	// Casbin
	AuthConfigPath string
	CSVFilePath    string

	// Redis
	RedisHost  string
	RedisPort  int

	LogLevel string
	HTTPPort string
	HTTPHost string

	SignInKey          string
	AccessTokenTimeout int
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8090"))

	// User Service
	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "127.0.0.1"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 1110))

	// Cource Service
	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "127.0.0.1"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 2220))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	// Redis
	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "redis"))
	c.RedisPort = cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379))

	// casbin...
	c.AuthConfigPath = cast.ToString(getOrReturnDefault("AUTH_CONFIG_PATH", "./config/auth.conf"))
	c.CSVFilePath = cast.ToString(getOrReturnDefault("CSV_FILE_PATH", "./config/auth.csv"))

	// jwt
	c.SignInKey = cast.ToString(getOrReturnDefault("SIGN_IN_KEY", "test-key"))
	c.AccessTokenTimeout = cast.ToInt(getOrReturnDefault("ACCESS_TOKEN_TIMEOUT", 3600))

	// Contex timeout
	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
