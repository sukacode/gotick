package event

import (
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
)

func WatchBuyButton(page playwright.Page) {

	log.Println("watching buy button...")

	for {

		button := page.Locator(BuyButton).First()

		count, _ := button.Count()

		if count > 0 {

			enabled, _ := button.IsEnabled()

			log.Println("button found")

			if enabled {

				log.Println("BUTTON ENABLED")

				err := button.Click()

				if err != nil {

					log.Println(err)

					continue
				}

				log.Println("BUY BUTTON CLICKED")

				break
			}
		}

		time.Sleep(1 * time.Second)

		page.Reload()
	}
}