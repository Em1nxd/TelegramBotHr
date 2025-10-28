package main

import (
	"fmt"
	"regexp"
	"time"

	tele "gopkg.in/telebot.v3"
	"main.go/config"
	"main.go/logger"
)

type User struct {
	ID            int64
	FirstName     string
	Username      string
	PhoneNumber   string
	Language      string
	Age           string
	City          string
	WorkingAs     string
	Student       string
	Degree        string
	Photo         *tele.Photo
	WorkingAdress string
	Gender        string
	FreeWork      string
	AboutUs       string
	OurAdresses   string
	Contact       string
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

	fmt.Println(bot.Me)
	b.bot = bot

	// Register midlewares
	// midlewares := midleware.New(b.log, b.cfg, bot)

	bot.Use(func(next tele.HandlerFunc) tele.HandlerFunc {

		return func(ctx tele.Context) error {
			var (
				sender = ctx.Sender()
			)

			_, ok := b.users[sender.ID]
			if !ok {

				var menu = &tele.ReplyMarkup{
					OneTimeKeyboard: true,
					ResizeKeyboard:  true,
				}

				menu.Reply(tele.Row{tele.Btn{Text: "O'zbek Tili 🇺🇿"}, tele.Btn{Text: "Русcкий Язык 🇷🇺"}})
				b.users[sender.ID] = &SessionUser{
					step: "lang",
					user: &User{
						ID:        sender.ID,
						FirstName: sender.FirstName,
						Username:  sender.Username,
					},
				}

				return ctx.Send("Tilni tanlang!", menu)
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
	bot.Handle(tele.OnPhoto, b.Photo)

	b.bot.Send(&tele.Chat{ID: -1001805067522}, "Bot o'z ishini boshladi😉", &tele.SendOptions{})

	// go bot.Start()
	bot.Start()

	return bot, nil
}

func botOnError(err error, ctx tele.Context) {

}

func (b *Bot) Text(ctx tele.Context) error {

	var (
		text = ctx.Text()
		ques = &tele.ReplyMarkup{
			OneTimeKeyboard: true,
			ResizeKeyboard:  true,
		}
		ques1 = &tele.ReplyMarkup{
			OneTimeKeyboard: true,
			ResizeKeyboard:  true,
		}
		ques2 = &tele.ReplyMarkup{
			RemoveKeyboard: true,
			ResizeKeyboard: true,
		}
	)

	if text == "🏠Menyu" {
		b.users[ctx.Sender().ID].step = "about_us"
		ques.Reply(tele.Row{tele.Btn{Text: "💼Bo'sh ish o'rinlari"}, tele.Btn{Text: "🏢Biz haqimizda"}}, tele.Row{tele.Btn{Text: "📍Manzillarimiz"}, tele.Btn{Text: "📞Aloqa"}})

		return ctx.Send("Ozingizga kerakli menyuni tanlang:", ques)
	}

	switch b.users[ctx.Sender().ID].step {
	case "lang":

		b.users[ctx.Sender().ID].user.Language = text
		b.users[ctx.Sender().ID].step = "about_us"
		ques.Reply(tele.Row{tele.Btn{Text: "💼Bo'sh ish o'rinlari"}, tele.Btn{Text: "🏢Biz haqimizda"}}, tele.Row{tele.Btn{Text: "📍Manzillarimiz"}, tele.Btn{Text: "📞Aloqa"}})

		return ctx.Send("Ozingizga kerakli menyuni tanlang:", ques)

	case "about_us":
		switch text {
		case "💼Bo'sh ish o'rinlari":

			b.users[ctx.Sender().ID].user.AboutUs = text
			b.users[ctx.Sender().ID].step = "working_place"

			ques.Reply(tele.Row{tele.Btn{Text: "Tiin Sayram"}}, tele.Row{tele.Btn{Text: "🔙Orqaga"}, tele.Btn{Text: "🏠Menyu"}})

			return ctx.Send("Qaysi filialda ishlamoqchisiz?", ques)
		case "🏢Biz haqimizda":
			return ctx.Send("Tiin ulgurji market")
		case "📍Manzillarimiz":
			return ctx.Send("Tiin Sayram 5/92")
		case "Aloqa":
			return ctx.Send("📞Aloqa uchun: +998935559562")

		}

	case "working_place":
		if text == "🔙Orqaga" {
			b.users[ctx.Sender().ID].step = "about_us"
			ques.Reply(tele.Row{tele.Btn{Text: "💼Bo'sh ish o'rinlari"}, tele.Btn{Text: "🏢Biz haqimizda"}}, tele.Row{tele.Btn{Text: "📍Manzillarimiz"}, tele.Btn{Text: "📞Aloqa"}})

			return ctx.Send("Ozingizga kerakli menyuni tanlang:", ques)
		}
		b.users[ctx.Sender().ID].user.WorkingAdress = text
		isMatch, err := regexp.Match("Tiin Sayram", []byte(text))
		if err != nil || !isMatch {

			ques.Reply(tele.Row{tele.Btn{Text: "Tiin Sayram"}}, tele.Row{tele.Btn{Text: "🔙Orqaga"}, tele.Btn{Text: "🏠Menyu"}})

			return ctx.Send("Filialni to'g'ri tanlang!", ques)
		}
		b.users[ctx.Sender().ID].step = "working_as"

		ques1.Reply(tele.Row{tele.Btn{Text: "Kassir"}, tele.Btn{Text: "Sotuvchi"}}, tele.Row{tele.Btn{Text: "Oxrana"}, tele.Btn{Text: "Ofis hodimi"}}, tele.Row{tele.Btn{Text: "🔙Orqaga"}, tele.Btn{Text: "🏠Menyu"}})

		return ctx.Send("Qaysi lavozimga topshiryapsiz?", ques1)

	case "working_as":
		if text == "🔙Orqaga" {
			b.users[ctx.Sender().ID].step = "working_place"
			ques.Reply(tele.Row{tele.Btn{Text: "Tiin Sayram"}}, tele.Row{tele.Btn{Text: "🔙Orqaga"}, tele.Btn{Text: "🏠Menyu"}})

			return ctx.Send("Qaysi filialda ishlamoqchisiz?", ques)
		}
		b.users[ctx.Sender().ID].user.WorkingAs = text
		isMatch, err := regexp.Match("Kassir|Sotuvchi|Oxrana|Ofis hodimi", []byte(text))
		if err != nil || !isMatch {

			ques1.Reply(tele.Row{tele.Btn{Text: "Kassir"}, tele.Btn{Text: "Sotuvchi"}}, tele.Row{tele.Btn{Text: "Oxrana"}, tele.Btn{Text: "Ofis hodimi"}}, tele.Row{tele.Btn{Text: "🔙Orqaga"}, tele.Btn{Text: "🏠Menyu"}})

			return ctx.Send("Lavozimni to'g'ri kiriting!!!", ques)
		}
		b.users[ctx.Sender().ID].step = "name"

		return ctx.Send("To'liq ismingizni kiriting (Murodjon Tursunov Husanboy o'g'li):")
	case "name":
		b.users[ctx.Sender().ID].user.FirstName = text
		b.users[ctx.Sender().ID].step = "age"

		return ctx.Send("Tug'ilgan sanangiz (masalan: 18.03.1995):")
	case "age":
		b.users[ctx.Sender().ID].user.Age = text
		b.users[ctx.Sender().ID].step = "gender"
		ques.Reply(tele.Row{tele.Btn{Text: "🧑Erkak"}, tele.Btn{Text: "👩Ayol"}}, tele.Row{tele.Btn{Text: "🔙Orqaga"}, tele.Btn{Text: "🏠Menyu"}})

		return ctx.Send("Jinsingiz:", ques)
	case "gender":
		if text == "🔙Orqaga" {
			b.users[ctx.Sender().ID].step = "age"
			return ctx.Send("Tug'ilgan sanangiz (masalan: 18.03.1995):")
		}
		b.users[ctx.Sender().ID].user.Gender = text

		isMatch, err := regexp.Match("🧑Erkak|👩Ayol", []byte(text))
		if err != nil || !isMatch {

			ques.Reply(tele.Row{tele.Btn{Text: "🧑Erkak"}, tele.Btn{Text: "👩Ayol"}}, tele.Row{tele.Btn{Text: "🔙Orqaga"}, tele.Btn{Text: "🏠Menyu"}})

			return ctx.Send("Jinsingizni xato kiritdingiz!!!", ques)
		}
		b.users[ctx.Sender().ID].step = "city"

		// ques.Reply(tele.Row{tele.Btn{Text: "Xa"}, tele.Btn{Text: "Yoq"}})

		return ctx.Send("Yashash manzilingiz? (masalan: Toshkent shahar, Chilonzor tuman)")
	// case "student":
	// 	b.users[ctx.Sender().ID].user.Student = text
	// 	b.users[ctx.Sender().ID].step = "city"

	// 	return ctx.Send("Yashash manzilingiz:")
	case "city":
		b.users[ctx.Sender().ID].user.City = text
		b.users[ctx.Sender().ID].step = "phone_number"

		return ctx.Send("Telefon raqamingizni kiriting (masalan: +998991234567):")
	case "phone_number":
		b.users[ctx.Sender().ID].user.PhoneNumber = text
		b.users[ctx.Sender().ID].step = "degree"
		ques2.Reply(tele.Row{tele.Btn{Text: "Oliy"}, tele.Btn{Text: "O'rta"}}, tele.Row{tele.Btn{Text: "O'rta maxsus"}, tele.Btn{Text: "Talaba"}}, tele.Row{tele.Btn{Text: "🔙Orqaga"}, tele.Btn{Text: "🏠Menyu"}})

		return ctx.Send("Ma'lumotingiz qanday?", ques2)
	case "degree":
		if text == "🔙Orqaga" {
			b.users[ctx.Sender().ID].step = "phone_number"

			return ctx.Send("Telefon raqamingizni kiriting (masalan: +998991234567):")
		}
		b.users[ctx.Sender().ID].user.Degree = text
		isMatch, err := regexp.Match("Oliy|O'rta|O'rta maxsus|Talaba", []byte(text))
		if err != nil || !isMatch {

			ques2.Reply(tele.Row{tele.Btn{Text: "Oliy"}, tele.Btn{Text: "O'rta"}}, tele.Row{tele.Btn{Text: "O'rta maxsus"}, tele.Btn{Text: "Talaba"}}, tele.Row{tele.Btn{Text: "🔙Orqaga"}, tele.Btn{Text: "🏠Menyu"}})

			return ctx.Send("Ma'lumotingizni to'g'ri kiriting!!!", ques)
		}
		b.users[ctx.Sender().ID].step = "photo"

		return ctx.Send("Suratingizni yuboring (telefoningizdan selfi olishingiz mumkin):")
	}

	return ctx.Send("Ariza faqat bir marta jonatiladi!")

	// Instead, prefer a context short-hand:
}

func (b *Bot) Photo(ctx tele.Context) error {

	var (
		photo = ctx.Message().Photo
		ques  = &tele.ReplyMarkup{
			OneTimeKeyboard: true,
			ResizeKeyboard:  true,
		}
	)

	switch b.users[ctx.Sender().ID].step {
	case "photo":

		if photo != nil {
			b.users[ctx.Sender().ID].user.Photo = photo
			fmt.Println(photo)
			b.PhotoSender(
				tele.Album{
					&tele.Photo{
						File: tele.File{FileID: photo.FileID, UniqueID: photo.UniqueID},
						Caption: fmt.Sprintf("📋Rezyume\n\n<b>📍Filial</b>:%s<b>\n👨‍💼Lavozim</b>:%s<b>\n📇Ism va Familiya</b>:%s<b>\n🔢Yosh</b>:%s<b>\n👥Jinsi</b>:%s<b>\n🏡Yashash manzili</b>:%s<b>\n📞Telefon Raqami</b>:%s<b>\n📃Ma'lumoti</b>:%s",
							b.users[ctx.Sender().ID].user.WorkingAdress,
							b.users[ctx.Sender().ID].user.WorkingAs,
							b.users[ctx.Sender().ID].user.FirstName,
							b.users[ctx.Sender().ID].user.Age,
							b.users[ctx.Sender().ID].user.Gender,
							b.users[ctx.Sender().ID].user.City,
							b.users[ctx.Sender().ID].user.PhoneNumber,
							b.users[ctx.Sender().ID].user.Degree,
						),
					},
				},
			)
			fmt.Println(photo)
		} else {
			return ctx.Send("Rasm xato jo'natildi!")
		}

		b.users[ctx.Sender().ID].step = "about_us"
		ques.Reply(tele.Row{tele.Btn{Text: "💼Bo'sh ish o'rinlari"}, tele.Btn{Text: "🏢Biz haqimizda"}}, tele.Row{tele.Btn{Text: "📍Manzillarimiz"}, tele.Btn{Text: "📞Aloqa"}})

		return ctx.Send("Rahmat. Siz ko'rib chiqiladigan nomzodlar ro'yxatidasiz.Hurmat bilan Tiin kadrlar bo'limi!", ques)
	}

	return ctx.Send("Ariza faqat bir marta jonatiladi!")

	// Instead, prefer a context short-hand:
}

func (b *Bot) MessageSender(messsage string) error {

	_, err := b.bot.Send(&tele.Chat{ID: -1001805067522}, messsage, &tele.SendOptions{
		ParseMode: "HTML",
	})

	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) PhotoSender(album tele.Album) error {

	_, err := b.bot.SendAlbum(&tele.Chat{ID: -1001805067522}, album, &tele.SendOptions{
		ParseMode: "HTML",
	})

	if err != nil {
		return err
	}

	return nil
}
func main() {
	bot := NewBot(logger.NewLogger("asdf", "info"), config.Load())

	// google.New()

	bot.NewBotWithPolling()
}
