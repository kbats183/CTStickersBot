package config

import (
	botserver "github.com/kbats183/CTStickersBot/pkg/bot-server"
	externalserverticker "github.com/kbats183/CTStickersBot/pkg/external-server-ticker"
	"github.com/kbats183/CTStickersBot/pkg/ocrapi"
	"github.com/kbats183/CTStickersBot/pkg/storage"
	"github.com/kbats183/CTStickersBot/pkg/tgbot"
)

type AppConfig struct {
	APPName string `default:"CTStickerBot"`

	DB storage.StorageConfig `yaml:"db"`

	TelegramBot tgbot.BotConfig `yaml:"telegram_bot"`

	OCR ocrapi.OCRClientConfig `yaml:"ocr"`

	ServerConfig botserver.ServerConfig `yaml:"server"`

	ServerTicker externalserverticker.ServerTickerConfig `yaml:"server_ticker"`
}
