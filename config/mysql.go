package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	DB_Host string
	DB_Port string
	DB_User string
	DB_Pass string
	DB_Name string
}

func (e *Environment) Init() {
	e.DB_Host = getEnv("DB_HOST")
	e.DB_Port = getEnv("DB_PORT")
	e.DB_User = getEnv("DB_USER")
	e.DB_Pass = getEnv("DB_PASS")
	e.DB_Name = getEnv("DB_NAME")
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Missing or invalid environment key")
	}
	return value
}

func LoadEnvironmentFile(file string) []string {
	if err := godotenv.Load(file); err != nil {
		fmt.Printf("Error on load environment file: %s", file)
	}
	e := &Environment{}
	e.Init()
	return []string{e.DB_Host, e.DB_Port, e.DB_User, e.DB_Pass, e.DB_Name}
}
