package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"chaos/extension"
	"chaos/proxy"
)

var (
	extensionName = filepath.Base(os.Args[0])
	logPrefix     = fmt.Sprintf("[%s] ", extensionName)
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	logger := log.New(os.Stdout, logPrefix, 0)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sigs
		cancel()
		logger.Print("Received", s)
		logger.Print("Exiting")
	}()

	client := extension.NewClient(os.Getenv("AWS_LAMBDA_RUNTIME_API"))

	res, err := client.Register(ctx, extensionName)
	if err != nil {
		logger.Fatal("could not register:", err)
	}
	logger.Print("Register response:", res)

	port := os.Getenv("EXTENSION_HTTP_PORT")
	if len(port) == 0 {
		port = "8888"
	}

	proxy.Start(logger, port)

	processEvents(ctx, logger, client)
}

func processEvents(ctx context.Context, logger *log.Logger, client *extension.Client) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			logger.Print("Waiting for event...")
			res, err := client.NextEvent(ctx)
			if err != nil {
				log.Print("Error:", err)
				log.Print("Exiting")
				return
			}

			log.Print("Received event:", res)
			if res.EventType == extension.Shutdown {
				log.Print("Received SHUTDOWN event")
				log.Print("Exiting")
				return
			}
		}
	}
}
