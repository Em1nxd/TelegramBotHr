package main

// import (
// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// 	"main.go/config"
// 	"main.go/logger"
// )

// var Token string = "5926355930:AAFNJ6xy7yasaXcHBhIUMyejTvnOyqDQaTo"

// // func main() {
// // 	bot, err := tgbotapi.NewBotAPI(Token)
// // 	if err != nil {
// // 		log.Panic(err)
// // 	}

// // 	bot.Debug = true

// // 	log.Printf("Authorized on account %s", bot.Self.UserName)

// // 	u := tgbotapi.NewUpdate(0)
// // 	u.Timeout = 60

// // 	updates := bot.GetUpdatesChan(u)

// // 	for update := range updates {
// // 		if update.Message != nil { // If we got a message
// // 			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

// // 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
// // 			msg.ReplyToMessageID = update.Message.MessageID

// // 			bot.Send(msg)
// // 		}
// // 	}

// // }

// // var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
// // 	tgbotapi.NewInlineKeyboardRow(
// // 		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
// // 		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
// // 		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
// // 	),
// // 	tgbotapi.NewInlineKeyboardRow(
// // 		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
// // 		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
// // 		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
// // 	),
// // )

// // func main() {
// // 	bot, err := tgbotapi.NewBotAPI(Token)
// // 	if err != nil {
// // 		log.Panic(err)
// // 	}

// // 	bot.Debug = true

// // 	log.Printf("Authorized on account %s", bot.Self.UserName)

// // 	u := tgbotapi.NewUpdate(0)
// // 	u.Timeout = 60

// // 	updates := bot.GetUpdatesChan(u)

// // 	for update := range updates {
// // 		if update.Message == nil { // ignore any non-Message updates
// // 			continue
// // 		}

// // 		if !update.Message.IsCommand() { // ignore any non-command Messages
// // 			continue
// // 		}

// // 		// Create a new MessageConfig. We don't have text yet,
// // 		// so we leave it empty.
// // 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

// // 		// Extract the command from the Message.
// // 		switch update.Message.Command() {
// // 		case "language":
// // 			msg.Text = "Выберите удобный вам язык. Русский Uzbek"
// // 		case "Русский":
// // 			msg.Text = "Вы выбрали Русский язык"
// // 		case "Uzbek":
// // 			msg.Text = "Siz O'zbek tilini tanladingiz"
// // 		default:
// // 			msg.Text = "Как вас зовут?"
// // 		}

// // 		if _, err := bot.Send(msg); err != nil {
// // 			log.Panic(err)
// // 		}
// // 	}
// // }

// var numericKeyboard = tgbotapi.NewReplyKeyboard(
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("O'zbek Tili 🇺🇿"),
// 		tgbotapi.NewKeyboardButton("Русккий Язык 🇷🇺"),
// 	),
// )

// func main() {

// 	bot := NewBot(logger.NewLogger("asdf", "info"), config.Load())

// 	bot.NewBotWithPolling(Token)
// 	// bot, err := tgbotapi.NewBotAPI(Token)
// 	// if err != nil {
// 	// 	log.Panic(err)
// 	// }

// 	// bot.Debug = true

// 	// log.Printf("Authorized on account %s", bot.Self.UserName)

// 	// u := tgbotapi.NewUpdate(0)
// 	// u.Timeout = 60

// 	// updates := bot.GetUpdatesChan(u)

// 	// for update := range updates {
// 	// 	if update.Message == nil { // ignore non-Message updates
// 	// 		continue
// 	// 	}

// 	// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Здравствуйте, это HR бот Tiin. Чтобы продолжить, напишите свое имя")

// 	// 	switch update.Message.Text {
// 	// 	case "Muhammadamin":
// 	// 		msg.Text = "Выберите удобный вам язык"
// 	// 		msg.ReplyMarkup = numericKeyboard
// 	// 		// msg.Text = "Сколько вам лет?"
// 	// 	case "O'zbek Tili 🇺🇿":
// 	// 		msg.Text = "Yoshingiz nechida?"
// 	// 	case "19":
// 	// 		msg.Text = "Tug'ilgan shahringiz?"
// 	// 	case "Toshkent":
// 	// 		msg.Text = "Ma'lumot uchun rahmat! Yaqin orada siz bilan bog'lanamiz"
// 	// 	case "Русккий Язык 🇷🇺":
// 	// 		msg.Text = "Сколько вам лет?"
// 	// 	case "20":
// 	// 		msg.Text = "Укажите город рождения"
// 	// 	case "Ташкент":
// 	// 		msg.Text = "Спасибо за информацию! Свяжемся с вами в скором времени"

// 	// 	}

// 	// 	if _, err := bot.Send(msg); err != nil {
// 	// 		log.Panic(err)
// 	// 	}
// 	// }
// }
