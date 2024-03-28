package server

import (
	"github.com/gofiber/template/mustache/v2"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/adharshmk96/shAuth/server/infra"
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
)

func StartHttpServer(port string) (*fiber.App, chan bool) {
	infra.LoadDefaultConfig()
	logger := infra.GetLogger()
	engine := mustache.NewFileSystem(http.FS(indexHTML), ".html")
	server := fiber.New(fiber.Config{
		Views: engine,
	})

	// middlewares
	server.Use(fiberLogger.New())

	setupRoutes(server)

	// Start the server
	if err := server.Listen(port); err != nil {
		logger.Error(err.Error())
	}

	// graceful shutdown
	done := make(chan bool)

	// A go routine that listens for os signals
	// it will block until it receives a signal
	// once it receives a signal, it will shut down close the done channel
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		if err := server.Shutdown(); err != nil {
			logger.Error(err.Error())
		}

		close(done)
	}()

	return server, done
}
