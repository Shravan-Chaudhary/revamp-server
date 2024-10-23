package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shravan-Chaudhary/revamp-server/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// custom logger if any
	// database setup
	// redis setup if any

	cfg := config.MustLoad()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message": "pong",
		})
	})

	server := &http.Server{
		Addr: cfg.Addr,
		Handler: r,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		slog.Info("server started on" ,slog.String("addr", cfg.Addr))
		err := server.ListenAndServe()
		if err != nil {
		log.Fatal("Failed to start server", err.Error())
	}
	}()

	<- done

	slog.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server shutdown failed", slog.String("error",err.Error()))
	}

	slog.Info("server shutdown successfully")

}