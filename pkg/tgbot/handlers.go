package tgbot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	context2 "github.com/kbats183/CTStickersBot/pkg/core/context"
	"github.com/kbats183/CTStickersBot/pkg/ocrapi"
	request_tokenizer "github.com/kbats183/CTStickersBot/pkg/request-tokenizer"
	"go.uber.org/zap"
	"strconv"
)

func (b *Bot) answerInline(ctx context2.Context, updateID int, inlineQuery *tgbotapi.InlineQuery) {
	userRequestToken := request_tokenizer.Tokenize(inlineQuery.Query)
	err := b.storage.SaveUserRequest(ctx, inlineQuery.From.ID, inlineQuery.From.UserName, inlineQuery.Query)
	if err != nil {
		ctx.Logger.Error("Can't save user request", zap.Error(err))
	} else {
		ctx.Logger.Info("User search request", zap.String("user_login", inlineQuery.From.UserName), zap.Any("user_request", userRequestToken))
	}
	var queryResults []interface{}
	stickers, err := b.storage.SearchStickers(context.Background(), userRequestToken, b.config.InlineStickerLimit, ctx.Logger)
	if err != nil {
		ctx.Logger.Error("Can't search stickers", zap.Error(err))
	}
	for _, sticker := range stickers {
		queryResults = append(queryResults,
			tgbotapi.NewInlineQueryResultCachedSticker(strconv.Itoa(sticker.ID), sticker.FileID, sticker.StickerTitle),
		)
	}

	//ctx.Logger.Info("Preparing inline answer", zap.Any("Answer", queryResults))
	query, err := b.tgBotApi.AnswerInlineQuery(tgbotapi.InlineConfig{InlineQueryID: inlineQuery.ID, Results: queryResults})
	if err != nil {
		ctx.Logger.Info("Can't send inline query result", zap.Error(err))
	} else if !query.Ok {
		ctx.Logger.Info("Can't send inline query result", zap.Any("response", query.Result))
	}
}

func (b *Bot) answerChosenInlineResult(ctx context2.Context, updateID int, chosenInlineResult *tgbotapi.ChosenInlineResult) {
	ctx.Logger.Info("ChosenInlineResult", zap.Any("result", chosenInlineResult))
}

func (b *Bot) answerMessageSticker(ctx context2.Context, updateID int, sticker *tgbotapi.Sticker) {
	url, err := b.tgBotApi.GetFileDirectURL(sticker.FileID)
	if err != nil {
		panic(err)
	}
	stickerLocalPath, err := prepareStickerToParsing(url)
	if err != nil {
		ctx.Logger.Error("Can't download sticker to parsing", zap.Error(err), zap.Any("sticker", sticker))
		return
	}
	parseAnswer, err := ctx.OCRClient.ParseImage(stickerLocalPath)
	if err != nil {
		ctx.Logger.Error("Can't parse image", zap.Error(err), zap.Any("sticker", sticker))
		return
	}
	stickerText := ocrapi.GetStringByParseAnswer(parseAnswer)
	createdStickerID, err := b.storage.CreateSticker(ctx, sticker, stickerText)
	if err != nil {
		ctx.Logger.Error("Can't create sticker", zap.Error(err), zap.Any("sticker", sticker))
		return
	}
	ctx.Logger.Info("Create sticker", zap.Any("sticker", sticker), zap.String("sticker_text", stickerText), zap.Int("sticker_id", createdStickerID))
}

func (b *Bot) answerMessage(ctx context2.Context, updateID int, message *tgbotapi.Message) {
	_, err := b.tgBotApi.Send(tgbotapi.NewMessage(message.Chat.ID, "Hello, "+message.Chat.UserName+"!"))
	if err != nil {
		ctx.Logger.Info("Can't answer message", zap.Error(err))
	}
}
