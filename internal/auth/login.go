package auth

import (
	"encoding/json"
	"log"
	"os"

	"github.com/playwright-community/playwright-go"
)

func SaveSession(page playwright.Page) {

	log.Println("login manually first")

	page.WaitForTimeout(60000)

	state, err := page.Context().StorageState()

	if err != nil {

		log.Println(err)

		return
	}

	data, err := json.Marshal(state)

	if err != nil {

		log.Println(err)

		return
	}

	err = os.WriteFile(
		"sessions/tiket.json",
		data,
		0644,
	)

	if err != nil {

		log.Println(err)

		return
	}

	log.Println("session saved")
}