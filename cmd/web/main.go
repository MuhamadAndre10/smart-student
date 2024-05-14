package main

import (
	"context"
	"flag"
	"github.com/MuhamadAndre10/smartStudentApp/internal/config"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	var wait time.Duration
	flag.DurationVar(
		&wait,
		"graceful-timeout",
		time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m",
	)
	flag.Parse()

	viper := config.NewViper()
	log := config.NewLog()
	db := config.NewDatabase(viper, log)
	mux := config.NewMuxRouter()
	validate := config.NewValidator(viper)

	config.Bootstrap(&config.BootstrapConfig{
		DB:        db,
		Mux:       mux,
		Log:       log,
		Config:    viper,
		Validator: validate,
	})

	// new server
	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      mux,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.

	err := srv.Shutdown(ctx)

	if err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.

	log.Info("shutting down")
	os.Exit(0)

}
