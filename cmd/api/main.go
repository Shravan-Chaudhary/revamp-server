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

	"github.com/Shravan-Chaudhary/revamp-server/internal/http/user"
	"github.com/Shravan-Chaudhary/revamp-server/internal/pkg/config"
	"github.com/Shravan-Chaudhary/revamp-server/internal/pkg/database"
	"github.com/Shravan-Chaudhary/revamp-server/internal/pkg/errors"
	"github.com/Shravan-Chaudhary/revamp-server/internal/pkg/health"
	"github.com/Shravan-Chaudhary/revamp-server/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.MustLoad()

	// Connect to MongoDB
	client, err := database.ConnectMongoDB(cfg.MONGO_URI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %s", err.Error())
	}
	slog.Info("Connected to MongoDB")

	responseHandler := response.NewResponseHandler(*cfg)
	mongoUserRepository := user.NewMongoUserRepository(client, cfg)
	userHandler := user.NewUserHandler(mongoUserRepository, responseHandler)

	r := gin.Default()

	isDev := true

	// Middlewares
	r.Use(errors.ErrorHandler(isDev))

	// Other routes
	r.NoRoute(func(c *gin.Context) {
		c.Error(errors.HttpErrors.NotFound("Route not found"))
	})

	r.GET("/health", func(c *gin.Context) {
		healthData ,err := health.HealthCheck(cfg.Env)
		if err != nil {
			c.Error(errors.HttpErrors.InternalServer(err.Error()))
			return
		}
		responseHandler.Send(c, http.StatusOK, response.Messages.Success, healthData)
	})

	// User routes
	r.GET("/", userHandler.HandleCreateUser)

	r.GET("/ping", func(c *gin.Context) {
		responseHandler.Send(c, http.StatusOK, response.Messages.Success, gin.H{
			"message": "pong",
		})
	})

	//  Start server
	server := &http.Server{
		Addr: cfg.Addr,
		Handler: r,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Graceful shutdown
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