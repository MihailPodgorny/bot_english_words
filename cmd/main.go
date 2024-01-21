package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tele "gopkg.in/telebot.v3"

	"github.com/MihailPodgorny/bot_english_words/config"
	"github.com/MihailPodgorny/bot_english_words/internal/handlers"

	"github.com/MihailPodgorny/bot_english_words/internal/scheduler"
	"github.com/MihailPodgorny/bot_english_words/pkg/postgres"
)

func main() {
	log.Println("starting bot...")

	// load config
	cfg, err := config.GetConfig("../config")
	if err != nil {
		log.Fatalf("Loading config: %v", err)
	}

	//  start db
	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		log.Panic("cannot start DB")
	}
	defer psqlDB.Close()

	ctx, _ := context.WithCancel(context.Background())
	// bot settings
	pref := tele.Settings{
		Token:  cfg.Telegram.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	bot.Handle("/word", handlers.SendWord)
	bot.Handle("/add", handlers.AddWord)

	s := scheduler.GetScheduler()
	chat := &TestSendable{chat_id: "1343890469"}

	go func() {
		bot.Start()
	}()

	go func() {
		for {
			_, _ = s.Every(1).Day().At("12:28").Do(bot.Send(chat, "that word 3"))
			_, t := s.NextRun()
			log.Print(t)
			time.Sleep(60 * time.Minute)
		}

	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Printf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Printf("ctx.Done: %v", done)
	}

}

type TestSendable struct {
	chat_id string
}

func (ts *TestSendable) Recipient() string {
	return ts.chat_id
}
