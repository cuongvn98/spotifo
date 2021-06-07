package main

import (
	"github.com/cuongvn98/golibs/config"
	"github.com/cuongvn98/golibs/logger"
	"go.uber.org/zap"
	"spotifo/configs"
)

func main() {
	var cfg configs.Config
	if err := config.Load([]string{"."}, "config", &cfg); err != nil {
		panic(err)
	}
	log, err := logger.NewZap(cfg.IsDev)
	if err != nil {
		panic(err)
	}
	sugaredLogger := log.Sugar()
	defer func(sugaredLogger *zap.SugaredLogger) {
		_ = sugaredLogger.Sync()
	}(sugaredLogger)

	sugaredLogger.Infow("hhihaah", "ha", "b")

}
