package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/PlantWaterMe/GardenMonitor/handlers"
	"github.com/PlantWaterMe/GardenMonitor/sensor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/allwinner"
)

func main() {

	envVar := loadEnvVariables()

	err := Init()
	if err != nil {
		log.Fatal(err)
	}

	ds := sensor.New(allwinner.PA16, allwinner.PA1)
	log.Printf("sensor created: +%v", ds)

	// create a new bot
	bot, err := tgbotapi.NewBotAPI(envVar["BOT_TOKEN"])
	if err != nil {
		log.Panic(err)
	}

	cfg := tgbotapi.NewSetMyCommands(handlers.Commands...)

	rsp, err := bot.Request(cfg)
	if err != nil {
		log.Panic(err)
	}
	log.Println(rsp)

	if envVar["ENVIRONMENT"] == "development" {
		bot.Debug = true
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	c := handlers.NewChat(bot, ds).
		Start()

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			c.Handle(&update)
		}
	}
}

func Init() error {
	// Make sure periph is initialized.
	state, err := host.Init()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to initialize periph: %v", err))
	}

	// Prints the loaded driver.
	fmt.Printf("Using drivers:\n")
	for _, driver := range state.Loaded {
		fmt.Printf("- %s\n", driver)
	}

	// Prints the driver that were skipped as irrelevant on the platform.
	fmt.Printf("Drivers skipped:\n")
	for _, failure := range state.Skipped {
		fmt.Printf("- %s: %s\n", failure.D, failure.Err)
	}

	// Having drivers failing to load may not require process termination. It
	// is possible to continue to run in partial failure mode.
	fmt.Printf("Drivers failed to load:\n")
	for _, failure := range state.Failed {
		fmt.Printf("- %s: %v\n", failure.D, failure.Err)
	}

	return nil
}

func loadEnvVariables() map[string]string {

	envVariables := make(map[string]string)

	// load en variables from ".env" file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	botToken, exists := os.LookupEnv("BOT_TOKEN")
	if !exists {
		panic("BOT_TOKEN undefined")
	}

	envVariables["BOT_TOKEN"] = botToken

	return envVariables
}
