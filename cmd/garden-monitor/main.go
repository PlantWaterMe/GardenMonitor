package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/PlantWaterMe/GardenMonitor/model"
	"github.com/PlantWaterMe/GardenMonitor/sensor"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/allwinner"
)

const REDIS_HOST = "REDIS_HOST"
const REDIS_PORT = "REDIS_PORT"
const REDIS_DB = "REDIS_DB"
const REDIS_PASSWORD = "REDIS_PASSWORD"
const REDIS_TOPIC = "REDIS_TOPIC"

func main() {
	envVar := loadEnvVariables()

	err := Init()
	if err != nil {
		log.Fatal(err)
	}

	ds := sensor.New(allwinner.PA16, allwinner.PA1)
	log.Printf("sensor created: +%v", ds)

	dbNumber, err := strconv.Atoi(envVar[REDIS_DB])
	if err != nil {
		log.Fatal(err)
	}

	// Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     envVar[REDIS_HOST] + ":" + envVar[REDIS_PORT],
		Password: envVar[REDIS_PASSWORD],
		DB:       dbNumber,
	})

	ticker := time.NewTicker(5 * time.Second)

	// this channel will signal when ctrl-c or a kill happen
	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)

	ctx := context.Background()
	for {
		var msg model.RedisMsg
		select {
		case <-exitSignal:
			fmt.Println("exit signal received, exiting...")
			return
		case t := <-ticker.C:
			msg = model.RedisMsg{
				Topic:       envVar[REDIS_TOPIC],
				CreatedAt:   time.Now(),
				SensorLevel: ds.Probe().String(),
			}

			json, err := json.Marshal(msg)
			if err != nil {
				log.Printf("error marshalling json: %s", err)
			}

			rp := rdb.Publish(ctx, envVar[REDIS_TOPIC], json)
			fmt.Printf("Tick at %s with rp: %s and message %+v \n", t, rp, msg)
		}
	}
}

func Init() error {
	// Make sure periph is initialized.
	state, err := host.Init()
	if err != nil {
		return fmt.Errorf("failed to initialize periph: %v", err)
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
	var exists bool

	// load en variables from ".env" file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	envVariables[REDIS_HOST], exists = os.LookupEnv(REDIS_HOST)
	if !exists {
		panic("missing env var: " + REDIS_HOST)
	}

	envVariables[REDIS_PORT], exists = os.LookupEnv(REDIS_PORT)
	if !exists {
		panic("missing env var: " + REDIS_PORT)
	}

	envVariables[REDIS_DB], exists = os.LookupEnv(REDIS_DB)
	if !exists {
		panic("missing env var: " + REDIS_DB)
	}

	envVariables[REDIS_PASSWORD], exists = os.LookupEnv(REDIS_PASSWORD)
	if !exists {
		envVariables[REDIS_PASSWORD] = ""
	}

	envVariables[REDIS_TOPIC], exists = os.LookupEnv(REDIS_TOPIC)
	if !exists {
		panic("missing env var: " + REDIS_TOPIC)
	}

	return envVariables
}
