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
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

const REDIS_HOST = "REDIS_HOST"
const REDIS_PORT = "REDIS_PORT"
const REDIS_DB = "REDIS_DB"
const REDIS_PASSWORD = "REDIS_PASSWORD"
const REDIS_TOPIC = "REDIS_TOPIC"

func main() {
	envVar := loadEnvVariables()

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

	ticker := time.NewTicker(30 * time.Second)

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
				SensorLevel: "low",
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
