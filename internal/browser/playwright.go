package browser

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

func Start() (*playwright.Playwright, playwright.BrowserContext, playwright.Page) {

	pw, err := playwright.Run()

	if err != nil {
		log.Fatal(err)
	}

	context, err := pw.Chromium.LaunchPersistentContext(
		"C:/playwright-profile",
		playwright.BrowserTypeLaunchPersistentContextOptions{
			Channel:  playwright.String("chrome"),
			Headless: playwright.Bool(false),
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	page := context.Pages()[0]

	return pw, context, page
}