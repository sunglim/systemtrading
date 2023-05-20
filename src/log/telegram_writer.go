package log

// Implementation of
// type Writer interface {
//	Write(p []byte) (n int, err error)
//}

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateTelegramWriter(token string, chatId int64) *TelegramWriter {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil
	}

	return &TelegramWriter{
		api:    bot,
		chatId: chatId,
	}
}

type TelegramWriter struct {
	api    *tgbotapi.BotAPI
	chatId int64
}

func (writer TelegramWriter) Write(p []byte) (n int, err error) {
	msg := tgbotapi.NewMessage(writer.chatId, string(p))
	_, err = writer.api.Send(msg)
	return len(p), err
}
