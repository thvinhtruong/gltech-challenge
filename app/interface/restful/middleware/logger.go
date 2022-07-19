package middleware

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	loggerOnce        sync.Once
	logger_middleware fiber.Handler
)

func saveLogFile(name string) *os.File {
	// create and save log based on current day
	path := "./log/" + name + ".log"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func LoggerMiddleware() fiber.Handler {
	loggerOnce.Do(func() {
		t := time.Now()
		name := t.Format("2006-01-02 15:04:05")
		logger_middleware = logger.New(logger.Config{
			Format:     "method=${method}, uri=${uri}, status=${status}, latency=${latency}, remote_ip=${remoteIP}, user_agent=${userAgent}",
			TimeFormat: "02-Jan-2006 15:04:05",
			TimeZone:   "local",
			Output:     saveLogFile(name),
		})
	})

	return logger_middleware
}
