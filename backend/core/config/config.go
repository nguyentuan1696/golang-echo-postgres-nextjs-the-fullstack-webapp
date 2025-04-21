package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port    int    `mapstructure:"port"`
	Host    string `mapstructure:"host"`
	BaseURL string `mapstructure:"base_url"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	ExpireTime int    `mapstructure:"expire_time"` // in hours
}

// Add new SMTP config struct
type SMTPConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	FromName string `mapstructure:"from_name"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Config struct {
	Environment Environment
	Server      ServerConfig
	Database    DatabaseConfig
	JWT         JWTConfig
	SMTP        SMTPConfig
	Redis       RedisConfig `mapstructure:"redis"`
}

var (
	instance *Config
	once     sync.Once
)

// Get returns the singleton config instance
func Get() *Config {
	if instance == nil {
		panic("Config not initialized. Call Init() first")
	}
	return instance
}

// Init initializes the configuration
// Add new environment type
type Environment string

const (
	DevEnvironment  Environment = "dev"
	ProdEnvironment Environment = "prod"
)

// Add environment field to Config struct

// Add SMTP environment bindings in Init function
func Init(env Environment) error {
	var err error
	once.Do(func() {
		// Load environment-specific .env file first
		envFile := fmt.Sprintf(".env.%s", env)
		if err = godotenv.Load(envFile); err != nil {
			// Try to load default .env if environment-specific file not found
			if err = godotenv.Load(); err != nil {
				fmt.Printf("Warning: no .env files found: %v\n", err)
			}
		}

		v := viper.New()

		// Set default values
		v.SetDefault("environment", string(env))
		v.SetDefault("server.port", 8080)
		v.SetDefault("server.host", "localhost")
		v.SetDefault("database.port", 5432)

		// Read from environment variables
		v.AutomaticEnv()
		v.SetEnvPrefix("APP")
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		// Bind environment variables
		v.BindEnv("server.port", "APP_SERVER_PORT")
		v.BindEnv("server.host", "APP_SERVER_HOST")
		v.BindEnv("server.base_url", "APP_SERVER_BASE_URL")
		v.BindEnv("database.host", "APP_DATABASE_HOST")
		v.BindEnv("database.port", "APP_DATABASE_PORT")
		v.BindEnv("database.user", "APP_DATABASE_USER")
		v.BindEnv("database.password", "APP_DATABASE_PASSWORD")
		v.BindEnv("database.dbname", "APP_DATABASE_DBNAME")
		v.BindEnv("jwt.secret", "APP_JWT_SECRET")
		v.BindEnv("jwt.expire_time", "APP_JWT_EXPIRE_TIME")
		// Add SMTP default values
		v.SetDefault("smtp.port", 587)
		v.SetDefault("smtp.host", "smtp.gmail.com")
		// Bind SMTP environment variables
		v.BindEnv("smtp.host", "APP_SMTP_HOST")
		v.BindEnv("smtp.port", "APP_SMTP_PORT")
		v.BindEnv("smtp.username", "APP_SMTP_USERNAME")
		v.BindEnv("smtp.password", "APP_SMTP_PASSWORD")
		v.BindEnv("smtp.from_name", "APP_SMTP_FROM_NAME")

		// Read from config file
		// Load environment-specific config file
		v.SetConfigName(fmt.Sprintf("config.%s", env))
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		v.AddConfigPath("./config")

		// Try environment-specific config first
		if err = v.ReadInConfig(); err != nil {
			// If not found, try default config
			v.SetConfigName("config")
			if err = v.ReadInConfig(); err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
					return
				}
			}
		}

		instance = &Config{}
		if err = v.Unmarshal(instance); err != nil {
			err = fmt.Errorf("unable to decode config into struct: %w", err)
			return
		}
	})

	return err
}
