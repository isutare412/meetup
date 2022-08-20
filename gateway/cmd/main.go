package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/isutare412/meetup/gateway/pkg/config"
	"github.com/isutare412/meetup/gateway/pkg/logger"
)

var cfgPath = flag.String("config", "configs/local/config.yaml", "path to yaml config file")

func main() {
	flag.Parse()

	var cfg *config.Config
	cfg, err := config.Load(*cfgPath)
	if err != nil {
		panic(fmt.Errorf("loading config: %w", err))
	}
	if err := cfg.Validate(); err != nil {
		panic(fmt.Errorf("invalid config: %w", err))
	}

	logger.Init(cfg.Logger)
	defer logger.Sync()

	startupCtx, cancel := context.WithTimeout(context.Background(), cfg.Timeout.Startup)
	defer cancel()

	var comp *components
	comp, err = newComponents(startupCtx, cfg)
	if err != nil {
		logger.S().Fatalf("Injecting dependencies: %v", err)
	}

	if err := comp.init(startupCtx); err != nil {
		logger.S().Fatalf("Initializing components: %v", err)
	}

	runCtx := context.Background()
	runFails := comp.run(runCtx)

	signals := make(chan os.Signal, 3)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case sig := <-signals:
		logger.S().Infof("Caught signal(%s)", sig.String())
	case err := <-runFails:
		logger.S().Error("Error while running components: %v", err)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Timeout.Shutdown)
	defer cancel()

	comp.shutdown(shutdownCtx)
}
