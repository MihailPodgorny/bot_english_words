package handlers

import (
	"encoding/json"
	"log"

	tele "gopkg.in/telebot.v3"

	"github.com/MihailPodgorny/bot_english_words/internal/models"
)

const (
	adminUser = 12345678
)

func SendWord(c tele.Context) error {
	log.Printf("get message from: %v", c.Chat().ID)
	return c.Send("word")
}

func AddWord(c tele.Context) error {
	user := c.Chat().ID
	log.Printf("get json from: %v", user)

	// validate user
	if user != adminUser {
		log.Print("wrong user tried to add new words")
	}
	message := c.Text()[5:]

	var targets []models.Word
	err := json.Unmarshal([]byte(message), &targets)
	if err != nil {
		log.Printf("got error: %v", err)
		return err
	}
	for _, word := range targets {
		log.Printf("get new word: %v", word.Text)
	}

	return c.Send("added")
}
