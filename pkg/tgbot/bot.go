package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kbats183/CTStickersBot/pkg/core"
	"github.com/kbats183/CTStickersBot/pkg/core/context"
	"github.com/kbats183/CTStickersBot/pkg/storage"
	"go.uber.org/zap"
	"time"
)

type Bot struct {
	config   core.BotConfig
	tgBotApi *tgbotapi.BotAPI
	storage  *storage.Storage
}

func NewBot(config core.BotConfig, storage *storage.Storage) (*Bot, error) {
	tgBotApi, err := tgbotapi.NewBotAPI(config.BotAuthToken)
	if err != nil {
		return nil, err
	}
	return &Bot{config: config, tgBotApi: tgBotApi, storage: storage}, nil
}

func (b *Bot) StartListening(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.tgBotApi.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		b.updateHandler(ctx, &update)
	}
	return nil
}

func (b *Bot) updateHandler(ctx context.Context, update *tgbotapi.Update) {
	upID := update.UpdateID
	ctx.Logger.Debug("Start update", zap.Int("update_id", upID), zap.Any("update", update))
	startUpdateProcessing := time.Now()
	if update.InlineQuery != nil {
		b.answerInline(ctx, upID, update.InlineQuery)
	} else if update.ChosenInlineResult != nil {
		b.answerChosenInlineResult(ctx, upID, update.ChosenInlineResult)
	} else if update.Message != nil && update.Message.Sticker != nil {
		b.answerMessageSticker(ctx, upID, update.Message.Sticker)
	} else if update.Message != nil {
		b.answerMessage(ctx, upID, update.Message)
	} else {
		ctx.Logger.Info("Unknown update", zap.Any("update", update))
	}
	ctx.Logger.Debug("End update", zap.Int("update_id", upID), zap.Int64("duration", int64(time.Now().Sub(startUpdateProcessing))))
}

func (b *Bot) GetBotUserName() string {
	return b.tgBotApi.Self.UserName
}
