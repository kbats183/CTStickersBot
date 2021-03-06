package tgbot

import (
	"context"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	botcontext "github.com/kbats183/CTStickersBot/pkg/core/context"
	"github.com/kbats183/CTStickersBot/pkg/ocrapi"
	request_tokenizer "github.com/kbats183/CTStickersBot/pkg/request-tokenizer"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

func (b *Bot) answerInline(ctx botcontext.Context, updateID int, inlineQuery *tgbotapi.InlineQuery) {
	startUpdateProcessing := time.Now()

	userRequestToken := request_tokenizer.Tokenize(inlineQuery.Query)
	updateIDStr := strconv.Itoa(updateID)
	err := b.storage.SaveUserRequest(ctx, inlineQuery.From.ID, inlineQuery.From.UserName, updateID, inlineQuery.Query)
	if err != nil {
		ctx.Logger.Error("Can't save user request",
			zap.Int("update_id", updateID),
			zap.Error(err))
	}
	stickers, err := b.storage.SearchStickers(context.Background(), userRequestToken, b.config.InlineStickerLimit)
	if err != nil {
		ctx.Logger.Error("Can't search stickers",
			zap.Int("update_id", updateID),
			zap.Any("user_request", userRequestToken),
			zap.Error(err))
		return
	}

	var queryResults []interface{}
	for _, sticker := range stickers {
		queryResults = append(queryResults,
			tgbotapi.NewInlineQueryResultCachedSticker(updateIDStr+"_sticker_"+strconv.Itoa(sticker.ID), sticker.FileID, updateIDStr+"_"+sticker.StickerTitle),
		)
	}

	_, err = b.tgBotApi.Send(tgbotapi.InlineConfig{
		InlineQueryID: inlineQuery.ID,
		Results:       queryResults,
		CacheTime:     1})
	_, isUnmarshalTypeError := err.(*json.UnmarshalTypeError)
	if err != nil && !isUnmarshalTypeError {
		ctx.Logger.Info("Can't send inline message result",
			zap.Int("update_id", updateID),
			zap.Error(err))
	} else {
		ctx.Logger.Info("Answer user's search request",
			zap.Int("update_id", updateID),
			zap.Int64("duration", time.Now().Sub(startUpdateProcessing).Nanoseconds()),
			zap.Any("user_request", userRequestToken))
	}
}

func (b *Bot) answerChosenInlineResult(ctx botcontext.Context, updateID int, chosenInlineResult *tgbotapi.ChosenInlineResult) {
	ctx.Logger.Debug("Chosen inline result",
		zap.Int("update_id", updateID),
		zap.Any("result", chosenInlineResult))
	resultIDFragment := strings.Split(chosenInlineResult.ResultID, "_")
	if len(resultIDFragment) != 3 || resultIDFragment[1] != "sticker" {
		ctx.Logger.Info("Dont save chosen inline result", zap.Any("result", chosenInlineResult))
		return
	}
	tgRequestID, err := strconv.Atoi(resultIDFragment[0])
	if err != nil {
		ctx.Logger.Error("Cant parse chosen inline result id", zap.Any("result", chosenInlineResult), zap.Error(err))
		return
	}
	stickerID, err := strconv.Atoi(resultIDFragment[2])
	if err != nil {
		ctx.Logger.Error("Cant parse chosen inline result id", zap.Any("result", chosenInlineResult), zap.Error(err))
		return
	}
	err = b.storage.SaveUserRequestChosenSticker(ctx, tgRequestID, stickerID)
	if err != nil {
		ctx.Logger.Error("Cant chosen inline result", zap.Any("result", chosenInlineResult), zap.Error(err))
		return
	}
	ctx.Logger.Info("Save chosen inline result id", zap.Any("result", chosenInlineResult), zap.Error(err))
}

func (b *Bot) answerMessageSticker(ctx botcontext.Context, updateID int, message *tgbotapi.Message) {
	isThisUserAdmin, err := b.storage.CheckAdminTelegram(ctx, message.From.ID, message.From.UserName)
	sticker := message.Sticker
	if err != nil {
		ctx.Logger.Error("Can't check telegram user permission",
			zap.Int("update_id", updateID),
			zap.Any("message", message),
			zap.Error(err))
		return
	} else if isThisUserAdmin == 0 {
		_, err := b.tgBotApi.Send(tgbotapi.NewMessage(message.Chat.ID, "???????????????????? ????????????! ??????????????????, ???????? ???? ?? ???????? ???????????"))
		if err != nil {
			ctx.Logger.Error("Can't answer message",
				zap.Int("update_id", updateID),
				zap.Error(err))
		}
		return
	}

	url, err := b.tgBotApi.GetFileDirectURL(sticker.FileID)
	if err != nil {
		panic(err)
	}
	stickerLocalPath, err := prepareStickerToParsing(url)
	if err != nil {
		ctx.Logger.Error("Can't download sticker to parsing",
			zap.Int("update_id", updateID),
			zap.Any("sticker", sticker),
			zap.Error(err))
		return
	}
	parseAnswer, err := ctx.OCRClient.ParseImage(stickerLocalPath)
	if err != nil {
		ctx.Logger.Error("Can't parse image",
			zap.Int("update_id", updateID),
			zap.Any("sticker", sticker),
			zap.Error(err))
		return
	}
	ctx.Logger.Debug("OCR api answer", zap.Any("answer", parseAnswer))
	stickerText := ocrapi.GetStringByParseAnswer(parseAnswer)
	stickerSetInfo, err := b.tgBotApi.GetStickerSet(tgbotapi.GetStickerSetConfig{Name: sticker.SetName})
	if err != nil {
		ctx.Logger.Error("Can't get sticker set info",
			zap.Int("update_id", updateID),
			zap.Any("sticker", sticker),
			zap.Error(err))
		return
	}

	err = b.storage.CreateStickerSet(ctx, sticker.SetName, stickerSetInfo.Title)
	if err != nil {
		ctx.Logger.Error("Can't create sticker set",
			zap.Int("update_id", updateID),
			zap.Any("sticker", sticker),
			zap.Error(err))
		return
	}
	createdStickerID, err := b.storage.CreateSticker(ctx, sticker, stickerText)
	if err != nil {
		ctx.Logger.Error("Can't create sticker",
			zap.Int("update_id", updateID),
			zap.Any("sticker", sticker),
			zap.Error(err))
		return
	}

	ctx.Logger.Info("Create sticker",
		zap.Int("update_id", updateID),
		zap.Any("sticker", sticker),
		zap.String("sticker_text", stickerText),
		zap.Int("sticker_id", createdStickerID))

	_, err = b.tgBotApi.Send(tgbotapi.NewMessage(message.Chat.ID, "ok"))
	if err != nil {
		ctx.Logger.Error("Can't answer message",
			zap.Int("update_id", updateID),
			zap.Error(err))
	}
}

func (b *Bot) answerMessage(ctx botcontext.Context, updateID int, message *tgbotapi.Message) {
	_, err := b.tgBotApi.Send(tgbotapi.NewMessage(message.Chat.ID, "Hello, "+message.Chat.UserName+"!"))
	if err != nil {
		ctx.Logger.Info("Can't answer message",
			zap.Int("update_id", updateID),
			zap.Error(err))
	}
}
