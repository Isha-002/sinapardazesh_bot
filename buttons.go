package main

import (
	tele "gopkg.in/telebot.v3"
)

var (
	zar_app_btn = tele.Btn{Unique: "zar_app", Text: zar_app, URL: zar_link}
	sina_app_btn = tele.Btn{Unique: "sina_app", Text: sina_app, URL: sina_link}
	q_and_a_btn = tele.Btn{Unique: "q&a_app", Text: q_and_a, URL: q_and_a_link}
	app_demo_btn = tele.Btn{Unique: "zar_app", Text: apps_demo, URL: apps_demo_link}
	address_btn = tele.Btn{Unique: "address_app", Text: company_address_label}
	contact_btn = tele.Btn{Unique: "contact_app", Text: contact_support_label}

	backBtn = tele.Btn{Unique: "back_btn", Text: "↩️ بازگشت"}
)