package main

import (
	"context"
	"log"
	"os"
	"strconv"

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
		Password: "",
		DB:       dbNumber,
	})

	subs := rdb.Subscribe(context.Background(), envVar[REDIS_TOPIC])
	// Close the subscription when we are done.
	defer subs.Close()

	log.Printf("env vars: %+v", envVar)
	for {
		msgStr, err := subs.ReceiveMessage(context.Background())
		if err != nil {
			panic(err)
		}

		log.Printf("msgStr: %+v\n", msgStr)
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

	envVariables[REDIS_TOPIC], exists = os.LookupEnv(REDIS_TOPIC)
	if !exists {
		panic("missing env var: " + REDIS_TOPIC)
	}

	return envVariables
}
