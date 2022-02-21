package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	keyLogLevel, defaultLogLevel = "LOG_LEVEL", "0"

	keyHTTPPort, defaultHTTPPort = "HTTP_PORT", "8000"
	keyHTTPHost, defaultHTTPHost = "HTTP_HOST", ""

	keyDBPort, defaultDBPort         = "DB_PORT", "5432"
	keyDBHost, defaultDBHost         = "DB_HOST", "localhost"
	keyDBUser, defaultDBUser         = "DB_USER", "root"
	keyDBPassword, defaultDBPassword = "DB_PAS", "example"
	keyDBName, defaultDBName         = "DB_NAME", "warehouse"

	keySigningKey, defaultSigningKey   = "TOK", "60e09d0d8fa190a9c6edb7bd"
	keyExpiresAt, defaultExpiresAt     = "TOK_EXPIRES_MIN", "30"
	keyRTExpiresAt, defaultRTExpiresAt = "REF_TOK_EXPIRES_MIN", "360"
)

type (
	Config struct {
		Log   LoggerConfig
		HTTP  HTTPConfig
		DB    DBConfig
		Token TokenConfig
	}

	LoggerConfig struct {
		Level int
	}

	HTTPConfig struct {
		Host string
		Port int
	}

	DBConfig struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}

	TokenConfig struct {
		SigningKey string
		// In minutes
		UTExpiresAt int
		// In minutes
		URTExpiresAT int
	}
)

func getEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return defaultValue
}

func getIntEnv(key, defaultValue string) (v int, err error) {
	value, ok := os.LookupEnv(key)
	if ok {
		v, err = strconv.Atoi(value)
		if err != nil {
			err = fmt.Errorf("invalid integer type by key: %v", key)
		}
		return
	}
	v, err = strconv.Atoi(defaultValue)
	return
}

func Init() (cfg *Config, err error) {
	cfg = new(Config)

	cfg.Log.Level, err = getIntEnv(keyLogLevel, defaultLogLevel)
	if err != nil {
		return
	}

	cfg.HTTP.Host = getEnv(keyHTTPHost, defaultHTTPHost)
	cfg.HTTP.Port, err = getIntEnv(keyHTTPPort, defaultHTTPPort)
	if err != nil {
		return
	}

	cfg.DB.Host = getEnv(keyDBHost, defaultDBHost)
	cfg.DB.Port = getEnv(keyDBPort, defaultDBPort)
	cfg.DB.User = getEnv(keyDBUser, defaultDBUser)
	cfg.DB.Password = getEnv(keyDBPassword, defaultDBPassword)
	cfg.DB.Name = getEnv(keyDBName, defaultDBName)

	cfg.Token.SigningKey = getEnv(keyExpiresAt, defaultSigningKey)
	cfg.Token.UTExpiresAt, err = getIntEnv(keyExpiresAt, defaultExpiresAt)
	if err != nil {
		return
	}
	cfg.Token.URTExpiresAT, err = getIntEnv(keyExpiresAt, defaultRTExpiresAt)
	if err != nil {
		return
	}

	return
}
