package main

import (
	"context"
	context2 "github.com/kbats183/CTStickersBot/pkg/core/context"
	"github.com/kbats183/CTStickersBot/pkg/ocrapi"
	"go.uber.org/zap/zapcore"

	"github.com/jinzhu/configor"
	"github.com/kbats183/CTStickersBot/pkg/core"
	"github.com/kbats183/CTStickersBot/pkg/storage"
	"github.com/kbats183/CTStickersBot/pkg/tgbot"
	"go.uber.org/zap"
	"log"
)

func main() {
	loggerConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseLevelEncoder,
		},
	}
	logger, _ := loggerConfig.Build()
	defer func() { _ = logger.Sync() }()

	var appConfig core.AppConfig
	err := configor.Load(&appConfig, "config_dev.yaml")
	if err != nil {
		logger.Error("Can't parse app config", zap.Error(err))
	}
	logger.Info("config", zap.Any("any", appConfig))

	ctx := context2.Context{Context: context.Background(), Logger: logger, OCRClient: ocrapi.NewOCRClient(appConfig.OCR)}

	st, err := storage.NewStorage(ctx, appConfig.DB)

	if err != nil {
		ctx.Logger.Fatal("Can't create storage", zap.Error(err))
	}

	bot, err := tgbot.NewBot(appConfig.TelegramBot, st)
	if err != nil {
		panic(err)
	}

	logger.Info("Bot user name", zap.String("bot_login", bot.GetBotUserName()))

	log.Fatal(bot.StartListening(ctx))

}
