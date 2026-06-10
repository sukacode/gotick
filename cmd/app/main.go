package main

import (
	"log"

	"gotick/internal/auth"
	"gotick/internal/browser"
	"gotick/internal/event"
)

func main() {

	pw, context, page := browser.Start()

	log.Println("browser started")

	_, err := page.Goto("https://www.tiket.com/id-id/to-do/my-chemical-romance-live-in-jakarta-2026?utm_page=toDoSearchResult")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("website opened")

	event.Monitor(page)

	event.WatchBuyButton(page)

	auth.SaveSession(page)

	select {}

	_ = pw
	_ = context
}