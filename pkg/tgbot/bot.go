package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	botcontext "github.com/kbats183/CTStickersBot/pkg/core/context"
	"github.com/kbats183/CTStickersBot/pkg/storage"
	"go.uber.org/zap"
	"time"
)

type BotConfig struct {
	BotAuthToken       string `yaml:"auth_token" env:"BOT_AUTH_TOKEN"`
	EnableDebug        bool   `yaml:"enable_debug" env:"BOT_ENABLE_DEBUG" default:"true"`
	InlineStickerLimit int    `yaml:"inline_stickers_limit" env:"BOT_INLINE_STICKER_LIMIT" default:"10"`
}

type Bot struct {
	config   BotConfig
	tgBotApi *tgbotapi.BotAPI
	storage  *storage.Storage
}

func NewBot(config BotConfig, storage *storage.Storage) (*Bot, error) {
	tgBotApi, err := tgbotapi.NewBotAPI(config.BotAuthToken)
	if err != nil {
		return nil, err
	}
	return &Bot{config: config, tgBotApi: tgBotApi, storage: storage}, nil
}

func (b *Bot) StartListening(ctx botcontext.Context) error {
	ctx.Logger.Info("Starting telegram bot ...")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.tgBotApi.GetUpdatesChan(u)

	for update := range updates {
		b.updateHandler(ctx, &update)
	}
	return nil
}

func (b *Bot) updateHandler(ctx botcontext.Context, update *tgbotapi.Update) {
	upID := update.UpdateID
	upAction := "unknown"
	ctx.Logger.Debug("Start update", zap.Int("update_id", upID), zap.Any("update", update))
	startUpdateProcessing := time.Now()
	if update.InlineQuery != nil {
		upAction = "InlineQuery"
		b.answerInline(ctx, upID, update.InlineQuery)
	} else if update.ChosenInlineResult != nil {
		upAction = "ChosenInlineResult"
		b.answerChosenInlineResult(ctx, upID, update.ChosenInlineResult)
	} else if update.Message != nil && update.Message.Sticker != nil {
		upAction = "MessageWithSticker"
		b.answerMessageSticker(ctx, upID, update.Message)
	} else if update.Message != nil {
		upAction = "Message"
		b.answerMessage(ctx, upID, update.Message)
	} else {
		ctx.Logger.Debug("Unknown update", zap.Any("update", update))
	}
	ctx.Logger.Debug("End update",
		zap.Int("update_id", upID),
		zap.Any("update", update),
		zap.String("update_action", upAction),
		zap.Int64("duration", int64(time.Now().Sub(startUpdateProcessing))))
}

func (b *Bot) GetBotUserName() string {
	return b.tgBotApi.Self.UserName
}
