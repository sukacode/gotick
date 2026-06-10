package browser

import (
	"log"
	"os"

	"github.com/playwright-community/playwright-go"
)

func NewContext(browser playwright.Browser) playwright.BrowserContext {

	sessionPath := "sessions/tiket.json"

	_, err := os.Stat(sessionPath)

	if err == nil {

		log.Println("loading saved session")

		context, err := browser.NewContext(
			playwright.BrowserNewContextOptions{
				StorageStatePath: playwright.String(sessionPath),
			},
		)

		if err != nil {
			log.Fatal(err)
		}

		return context
	}

	log.Println("creating fresh session")

	context, err := browser.NewContext()

	if err != nil {
		log.Fatal(err)
	}

	return context
}