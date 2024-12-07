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
		log.Println(desc_err, err)
	}

	userPages := make(map[int64]int)

	commands := []BotCommand{
		{Command: "start", Description: start_cmd},
		{Command: "about", Description: about_cmd},
	}

	
	var page_three = &tele.ReplyMarkup{}
	page_three.Inline(
		page_three.Row(zar_app_btn),
		page_three.Row(sina_app_btn),
		page_three.Row(app_demo_btn, q_and_a_btn),
		page_three.Row(contact_btn, address_btn),
		page_three.Row(prev_page),
	)

	var page_two = &tele.ReplyMarkup{}
	page_two.Inline(
		page_two.Row(chap_factor_app_btn),
		page_two.Row(modian_maliat_app_btn),
		page_two.Row(anbardari_app_btn, chap_check_app_btn),
		page_two.Row(khazane_app_btn, sina_matab_app_btn),
		page_two.Row(prev_page, next_page),
	)

	var page_one = &tele.ReplyMarkup{}
	page_one.Inline(
		page_one.Row(sina_website_btn),
		page_one.Row(hesabdari_app_btn),
		page_one.Row(automation_app_btn),
		page_one.Row(hesabdari_sherkati_btn),
		page_one.Row(next_page),
	)

	var back_to_main = &tele.ReplyMarkup{}
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
		return c.Send(bot_intro, page_one)
	})

	b.Handle(&next_page, func(c tele.Context) error {
		chatID := c.Sender().ID

		
		if _, exists := userPages[chatID]; !exists {
			userPages[chatID] = 1
		}

		
		switch userPages[chatID] {
		case 1:
			userPages[chatID] = 2
			return c.Edit(bot_intro, page_two)
		case 2:
			userPages[chatID] = 3
			return c.Edit(bot_intro, page_three)
		default:
			return c.Respond(&tele.CallbackResponse{
				Text: "You are already on the last page.",
			})
		}
	})

	b.Handle(&prev_page, func(c tele.Context) error {
		chatID := c.Sender().ID

		
		if _, exists := userPages[chatID]; !exists {
			userPages[chatID] = 1
		}

		
		switch userPages[chatID] {
		case 3:
			userPages[chatID] = 2
			return c.Edit(bot_intro, page_two)
		case 2:
			userPages[chatID] = 1
			return c.Edit(bot_intro, page_one)
		default:
			return c.Respond(&tele.CallbackResponse{
				Text: "You are already on the first page.",
			})
		}
	})

	b.Handle(&backBtn, func(c tele.Context) error {
		return c.Edit(bot_intro, page_one)
	})

	// Handle other buttons
	b.Handle(&contact_btn, func(c tele.Context) error {
		return c.Edit(contact_support_res, back_to_main)
	})
	b.Handle(&address_btn, func(c tele.Context) error {
		return c.Edit(company_address_res, back_to_main)
	})

	fmt.Println(on_start)
	b.Start()
	fmt.Println(on_crash)
}
