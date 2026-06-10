package event

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

func HandlePopup(page playwright.Page) {

	closeBtn := page.Locator(".rr___styles-module__close_button___wI4SI")

	count, _ := closeBtn.Count()

	log.Println("close button count:", count)

	if count > 0 {

		log.Println("popup detected")

		err := closeBtn.Click()

		if err != nil {

			log.Println("click failed:", err)

			return
		}

		log.Println("popup closed")
	}
}