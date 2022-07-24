package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server config
	Server_Host    string
	Server_Port    string
	Server_Timeout int

	// Database config
	DB_Host         string
	DB_Port         string
	DB_User         string
	DB_Pass         string
	DB_Name         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int // time in minute
	DB_SSLMode      string
	IsEnabledLog    bool

	// Cookie Config
	Cookie_Name     string
	Cookie_Domain   string
	Cookie_MaxAge   int
	Cookie_Secure   bool
	Cookie_HttpOnly bool

	// JWT Config
	JWT_Secret string
	JWT_Expire int

	// API
	API_Version string
	API_Domain  string
}

func (e *Config) ServerConn() {
	e.Server_Host = getEnv("SERVER_HOST")
	e.Server_Port = getEnv("SERVER_PORT")
	e.Server_Timeout, _ = strconv.Atoi(getEnv("SERVER_TIMEOUT"))
}

func (e *Config) DBConn() {
	e.DB_Host = getEnv("DB_HOST")
	e.DB_Port = getEnv("DB_PORT")
	e.DB_User = getEnv("DB_USER")
	e.DB_Pass = getEnv("DB_PASS")
	e.DB_Name = getEnv("DB_NAME")
	e.DB_SSLMode = getEnv("DB_SSL_MODE")
	e.MaxOpenConns, _ = strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS"))
	e.MaxIdleConns, _ = strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS"))
	e.ConnMaxLifetime, _ = strconv.Atoi(getEnv("DB_CONN_MAX_LIFETIME"))
	e.IsEnabledLog, _ = strconv.ParseBool(getEnv("DB_IS_ENABLED_LOG"))
}

func (e *Config) CookieConfig() {
	e.Cookie_Name = getEnv("COOKIE_NAME")
	e.Cookie_Domain = getEnv("COOKIE_DOMAIN")
	e.Cookie_MaxAge, _ = strconv.Atoi(getEnv("COOKIE_MAX_AGE"))
	e.Cookie_Secure, _ = strconv.ParseBool(getEnv("COOKIE_SECURE"))
	e.Cookie_HttpOnly, _ = strconv.ParseBool(getEnv("COOKIE_HTTP_ONLY"))
}

func (e *Config) JWTConfig() {
	e.JWT_Secret = getEnv("JWT_SECRET")
	e.JWT_Expire, _ = strconv.Atoi(getEnv("JWT_EXPIRE"))
}

func (e *Config) API() {
	e.API_Version = getEnv("API_VERSION")
	e.API_Domain = getEnv("API_DOMAIN")
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Missing or invalid environment key")
	}
	return value
}

func ReadEnvFile(path string) Config {
	if err := godotenv.Load(path); err != nil {
		fmt.Printf("Error on load environment file: %s", path)
	}
	e := Config{}
	//e.ServerConn()
	e.DBConn()
	//e.CookieConfig()
	return e
}
