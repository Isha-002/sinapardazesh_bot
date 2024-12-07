package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

func main() {
	fmt.Printf("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)

	err := godotenv.Load()
	if err != nil {
		log.Fatal(envError)
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	b, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = b.SetMyDescription(bot_description, "fa")
	if err != nil {
		log.Println("Error setting description:", err)
	}

	var menu = &tele.ReplyMarkup{}
	var back_to_main = &tele.ReplyMarkup{}

	commands := []BotCommand{
		{Command: "start", Description: start_cmd},
		{Command: "about", Description: about_cmd},
	}

	menu.Inline(
		menu.Row(zar_app_btn),
		menu.Row(sina_app_btn),
		menu.Row(app_demo_btn, q_and_a_btn),
		menu.Row(contact_btn, address_btn),
	)

	back_to_main.Inline(
		back_to_main.Row(backBtn),
	)

	teleCommands := make([]tele.Command, len(commands))
	for i, cmd := range commands {
		teleCommands[i] = tele.Command{
			Text:        cmd.Command,
			Description: cmd.Description,
		}
	}

	if err := b.SetCommands(teleCommands); err != nil {
		log.Fatalf(cmd_set_err, err)
	}

	var about_menu = &tele.ReplyMarkup{}
	about_menu.Inline(
		about_menu.Row(about_btn...),
		about_menu.Row(backBtn),
	)

	b.Handle("/about", func(c tele.Context) error {
		return c.Send(about_bot, about_menu)
	})

	b.Handle("/start", func(c tele.Context) error {
		return c.Send(bot_intro, menu)
	})

	b.Handle(&contact_btn, func(c tele.Context) error {
		return c.Edit(contact_support_res, back_to_main)
	})
	b.Handle(&address_btn, func(c tele.Context) error {
		return c.Edit(company_address_res, back_to_main)
	})

	b.Handle(&backBtn, func(c tele.Context) error {
		return c.Edit(bot_intro, menu)
	})

	fmt.Println(on_start)
	b.Start()
	fmt.Println(on_crash)
}
