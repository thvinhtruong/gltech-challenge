package main

import (
	"context"
	"time"

	"github.com/gofiber/fiber"
	"github.com/thvinhtruong/legoha/app/config"
)

type application struct {
	cfg *config.Config
	app *fiber.App
}

type operation func(ctx context.Context) error

func (a *application) CleanAndShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) error {
	return nil
}
