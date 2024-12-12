package main

import (
	tele "gopkg.in/telebot.v3"
)

var (
	zar_app_btn  = tele.Btn{Unique: "zar_app", Text: zar_app, URL: zar_link} // 3
	sina_app_btn = tele.Btn{Unique: "sina_app", Text: sina_app, URL: sina_link}
	q_and_a_btn  = tele.Btn{Unique: "q&a_app", Text: q_and_a, URL: q_and_a_link}
	app_demo_btn = tele.Btn{Unique: "zar_app", Text: apps_demo, URL: apps_demo_link}
	address_btn  = tele.Btn{Unique: "address_app", Text: company_address_label}
	contact_btn  = tele.Btn{Unique: "contact_app", Text: contact_support_label}
	//
	sina_website_btn       = tele.Btn{Unique: "sina_website", Text: sina_website, URL: sina_website_link} // 1
	hesabdari_app_btn      = tele.Btn{Unique: "hesabdari_app", Text: hesabdari_app, URL: hesabdari_link}
	automation_app_btn     = tele.Btn{Unique: "automation_app", Text: automation_app, URL: automation_link}
	hesabdari_sherkati_btn = tele.Btn{Unique: "hesabdari_sherkati", Text: hesabdari_sherkati_app, URL: hesabdari_sherkati_link}
	//
	chap_factor_app_btn    = tele.Btn{Unique: "chap_factor", Text: chap_factor_app, URL: chap_factor_link} // 2
	anbardari_app_btn      = tele.Btn{Unique: "anbardari", Text: anbardari_app, URL: anbardari_link}
	chap_check_app_btn     = tele.Btn{Unique: "chap_check", Text: chap_check_app, URL: chap_check_link}
	khazane_app_btn        = tele.Btn{Unique: "khazane", Text: khazane_app, URL: khazane_link}
	modian_maliat_app_btn  = tele.Btn{Unique: "modian_maliat", Text: modian_maliat_app, URL: modian_maliat_link}
	sina_matab_app_btn     = tele.Btn{Unique: "sina_matab", Text: sina_matab_app, URL: sina_matab_link}

	backBtn = tele.Btn{Unique: "back_btn", Text: "↩️ بازگشت"}
	next_page = tele.Btn{Unique: "next_page", Text: "➡️ صفحه بعد"}
	prev_page = tele.Btn{Unique: "prev_page", Text: "صفحه قبل ⬅️"}

	// about
	about_btn = []tele.Btn{
		{Unique: "isha", Text: isha, URL: isha_git},
	}
)
