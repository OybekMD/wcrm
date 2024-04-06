package config

import (
	"os"

	"github.com/spf13/cast"
)

type Kafka struct {
	Address            string
	UserCreateTopic    string
	CategoryCreateTopic string
	CategoryIconCreateTopic string
	ProductCreateTopic string
	OrderproductCreateTopic string
	CommentCreateTopic string
}

// Config ...
type Config struct {
	Environment string // develop, staging, production

	// User
	UserServiceHost string
	UserServicePort int

	// Post
	PostServiceHost string
	PostServicePort int

	// Comment
	CommentServiceHost string
	CommentServicePort int

	// context timeout in seconds
	CtxTimeout int

	// Casbin
	AuthConfigPath string
	CSVFilePath    string

	// Redis
	RedisHost string
	RedisPort int

	LogLevel string
	HTTPPort string
	HTTPHost string

	// Kafka
	Kafka Kafka

	// JWT
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

	// Post Service
	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "127.0.0.1"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 2220))

	// Comment Service
	c.CommentServiceHost = cast.ToString(getOrReturnDefault("COMMENT_SERVICE_HOST", "127.0.0.1"))
	c.CommentServicePort = cast.ToInt(getOrReturnDefault("COMMENT_SERVICE_PORT", 3330))

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

	// Kafka
	c.Kafka.Address = cast.ToString(getOrReturnDefault("KAFKA_ADDRESS", "localhost:9092"))
	c.Kafka.ProductCreateTopic = cast.ToString(getOrReturnDefault("PRODUCT_CREATE_TOPIC", "product.create"))
	c.Kafka.UserCreateTopic = cast.ToString(getOrReturnDefault("USER_CREATE_TOPIC", "user.create"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
