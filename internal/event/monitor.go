package event

import (
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
)

func Monitor(page playwright.Page) {

	log.Println("waiting popup...")

	time.Sleep(5 * time.Second)

	HandlePopup(page)

	log.Println("monitoring started")
}