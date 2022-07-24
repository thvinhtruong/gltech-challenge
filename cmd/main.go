package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/thvinhtruong/legoha/app/api"
	"github.com/thvinhtruong/legoha/app/config"
	"github.com/thvinhtruong/legoha/pkg/sqlconn"
	"golang.org/x/sync/errgroup"
)

type operation func(ctx context.Context) error

func CleanOperations(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})

	go func() {
		interruptCtx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		defer stop()

		// set timeout for long request to prevent sys hang

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()
		g, gCtx := errgroup.WithContext(interruptCtx)

		for _, op := range ops {
			g.Go(func() error {
				<-gCtx.Done()
				return op(gCtx)
			})
		}

		g.Wait()
		close(wait)

	}()

	return wait

}

func init() {
	log.Println("Server Started Successfully")
}

func main() {
	cfg := config.Config{
		DB_Host: "localhost",
		DB_Port: "3306",
		DB_User: "root",
		DB_Pass: "root",
		DB_Name: "todolist ",
	}
	sqlconn.Init(cfg)
	app := api.Restful()
	log.Fatal(app.Listen(":8080"))

	wait := CleanOperations(context.Background(), time.Duration(time.Second), map[string]operation{
		"database": func(ctx context.Context) error {
			return sqlconn.DB.Close()
		},

		"server": func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	<-wait

}
