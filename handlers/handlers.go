package handlers

import (
	// "github.com/creativeboy1999/botftaher/botfather/bot/builder"
	// "github.com/creativeboy1999/botftaher/config"
	// "github.com/creativeboy1999/botftaher/pkg/logger"
	// "github.com/creativeboy1999/botftaher/storage"
	tele "gopkg.in/telebot.v3"
	"main.go/config"
	"main.go/logger"
)

type BotHandler struct {
	bot *tele.Bot
	log logger.Logger
	cfg config.Config
	// builder builder.BuilderI
}

type BotHandlerI interface {
	RegisterAllHandlers()

	Text(tele.Context) error
}

// func New(log logger.Logger, cfg config.Config, builder builder.BuilderI, bot *tele.Bot) BotHandlerI {
func New(log logger.Logger, cfg config.Config, bot *tele.Bot) BotHandlerI {

	return &BotHandler{
		log: log,
		cfg: cfg,
		// builder: builder,
		bot: bot,
	}
}

func (h *BotHandler) RegisterAllHandlers() {
	h.bot.Handle(tele.OnText, h.Text)
}

	func (m *BotHandler) Text(c tele.Context) error {

		var (
			// user = c.Sender()
			text = c.Text()
		)

		if text == "/start" {

			// return c.Send(text, m.builder.Keyboard().Markup().Menu())
		}

		// Instead, prefer a context short-hand:
		return c.Send(text)
	}
