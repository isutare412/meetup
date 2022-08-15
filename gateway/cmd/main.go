package main

import (
	"flag"

	"github.com/isutare412/meetup/gateway/pkg/config"
	"github.com/isutare412/meetup/gateway/pkg/logger"
)

var cfgPath = flag.String("config", "configs/local/config.yaml", "path to yaml config file")

func main() {
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		panic(err)
	}

	logger.Init(cfg.Logger)
	defer logger.Sync()

	logger.S().Info("Hello, world!")
}
