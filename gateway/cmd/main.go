package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/isutare412/meetup/gateway/pkg/config"
	"github.com/isutare412/meetup/gateway/pkg/logger"
)

var cfgPath = flag.String("config", "configs/local/config.yaml", "path to yaml config file")

func main() {
	flag.Parse()

	var cfg *config.Config
	cfg, err := config.Load(*cfgPath)
	if err != nil {
		panic(err)
	}

	logger.Init(cfg.Logger)
	defer logger.Sync()

	startupCtx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	var comp *components
	comp, err = newComponents(startupCtx, cfg)
	if err != nil {
		logger.S().Fatalf("Injecting dependencies: %v", err)
	}

	if err := comp.init(startupCtx); err != nil {
		logger.S().Fatalf("Initializing components: %v", err)
	}

	runFails := comp.run(context.Background())

	signals := make(chan os.Signal, 3)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case sig := <-signals:
		logger.S().Infof("Caught signal(%s)", sig.String())
	case err := <-runFails:
		logger.S().Error("Error while running components: %v", err)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	comp.shutdown(shutdownCtx)
}
