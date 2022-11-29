package main

import (
	"fmt"
	"time"

	tele "gopkg.in/telebot.v3"
	"main.go/config"
	"main.go/logger"
)

type User struct {
	ID          int64
	FirstName   string
	LastName    string
	Username    string
	PhoneNumber string
	Language    string
	Age         string
	City        string
}
type SessionUser struct {
	step string
	user *User
}

type Bot struct {
	bot   *tele.Bot
	users map[int64]*SessionUser
	log   logger.Logger
	cfg   config.Config
	// builder builder.BuilderI
}

type BotI interface {
	NewBotWithPolling() (*tele.Bot, error)
}

func NewBot(log logger.Logger, cfg config.Config) BotI {
	return &Bot{log: log, cfg: cfg, users: map[int64]*SessionUser{}}
}

func (b *Bot) NewBotWithPolling() (*tele.Bot, error) {

	pref := tele.Settings{
		Token:     b.cfg.Token,
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		OnError:   botOnError,
		ParseMode: "HTML",
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	b.bot = bot

	// Register midlewares
	// midlewares := midleware.New(b.log, b.cfg, bot)

	bot.Use(func(next tele.HandlerFunc) tele.HandlerFunc {

		return func(ctx tele.Context) error {
			var (
				sender = ctx.Sender()
				menu   = &tele.ReplyMarkup{
					OneTimeKeyboard: true,
					ResizeKeyboard:  true,
				}
			)

			menu.Reply(tele.Row{tele.Btn{Text: "O'zbek Tili 🇺🇿"}, tele.Btn{Text: "Русккий Язык 🇷🇺"}})

			_, ok := b.users[sender.ID]
			if !ok {
				b.users[sender.ID] = &SessionUser{
					step: "lang",
					user: &User{
						ID:        sender.ID,
						FirstName: sender.FirstName,
						LastName:  sender.LastName,
						Username:  sender.Username,
					},
				}

				return ctx.Send("Выберите удобный вам язык", menu)
			}

			return next(ctx)
		}
	})
	// bot.Use(midlewares.Logger, midlewares.CheckUser)

	// Bot settings
	// settings := settings.New(bot, b.log)

	// err = settings.SetCommands()
	// if err != nil {
	// 	b.log.Error("error on SetCommands", logger.Any("err:", err))
	// }

	bot.Handle(tele.OnText, b.Text)

	b.bot.Send(&tele.Chat{ID: -1001805067522}, "", &tele.SendOptions{})

	// go bot.Start()
	bot.Start()

	return bot, nil
}

func botOnError(err error, ctx tele.Context) {

}

func (b *Bot) Text(ctx tele.Context) error {

	var (
		text   = ctx.Text()
		button = &tele.ReplyMarkup{
			OneTimeKeyboard: true,
			ResizeKeyboard:  true,
		}
	)

	switch b.users[ctx.Sender().ID].step {
	case "lang":

		b.users[ctx.Sender().ID].user.Language = text

		b.users[ctx.Sender().ID].step = "name"

		return ctx.Send("Напишите ваше имя")
	case "name":
		b.users[ctx.Sender().ID].user.FirstName = text

		b.users[ctx.Sender().ID].step = "surename"
		return ctx.Send("Напишите вашу фамилию")

	case "surename":
		b.users[ctx.Sender().ID].user.LastName = text

		b.users[ctx.Sender().ID].step = "phone_number"

		button.Reply(tele.Row{tele.Btn{Contact: true, Text: "Поделиться номером"}})

		return ctx.Send("📞 Напишите номер телефона", button)
	case "phone_number":
		if ctx.Message().Contact != nil {
			b.users[ctx.Sender().ID].user.PhoneNumber = ctx.Message().Contact.PhoneNumber
		} else {
			b.users[ctx.Sender().ID].user.PhoneNumber = text
		}

		b.users[ctx.Sender().ID].step = "age"

		return ctx.Send("🕑 Сколько вам лет?")
	case "age":
		b.users[ctx.Sender().ID].user.Age = text

		b.users[ctx.Sender().ID].step = "city"

		return ctx.Send("📍 В каком городе вы родились?")

	case "city":
		b.users[ctx.Sender().ID].user.City = text
		b.users[ctx.Sender().ID].step = ""

		b.Sender(fmt.Sprintf("Анкета📋\n<b>Имя </b>: %s\n\n<b>Фамилия </b>: %s\n\n<b>Телефон Номер📞 </b>: %s\n\n<b>Возраст🕑 </b>: %s \n\n<b>Город Рождения📍</b>:%s \n\n",
			b.users[ctx.Sender().ID].user.FirstName,
			b.users[ctx.Sender().ID].user.LastName,
			b.users[ctx.Sender().ID].user.PhoneNumber,
			b.users[ctx.Sender().ID].user.Age,
			b.users[ctx.Sender().ID].user.City,
		))
		

		return ctx.Send("Спасибо за ваш отзыв! В скором времени наши сотрудники свяжутся с вами.")
	}

	return ctx.Send("Спасибо за ваш отзыв! В скором времени наши сотрудники свяжутся с вами.")

	// Instead, prefer a context short-hand:
}

func (b *Bot) Sender(messsage string) error {

	_, err := b.bot.Send(&tele.Chat{ID: -1001805067522}, messsage, &tele.SendOptions{
		ParseMode: "HTML",
	})

	if err != nil {
		return err
	}

	return nil
}
func main() {
	bot := NewBot(logger.NewLogger("asdf", "info"), config.Load())

	bot.NewBotWithPolling()
}
