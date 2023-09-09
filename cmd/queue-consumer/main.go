package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/PlantWaterMe/GardenMonitor/model"
	"github.com/PlantWaterMe/GardenMonitor/repository"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const REDIS_HOST = "REDIS_HOST"
const REDIS_PORT = "REDIS_PORT"
const REDIS_DB = "REDIS_DB"
const REDIS_PASSWORD = "REDIS_PASSWORD"
const REDIS_TOPIC = "REDIS_TOPIC"
const DB_HOST = "DB_HOST"
const DB_USER = "DB_USER"
const DB_PW = "DB_PW"
const DB_NAME = "DB_NAME"
const DB_PORT = "DB_PORT"

func main() {
	envVar := loadEnvVariables()

	dbNumber, err := strconv.Atoi(envVar[REDIS_DB])
	if err != nil {
		log.Fatal(err)
	}

	// connect to db
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		envVar["DB_HOST"],
		envVar["DB_USER"],
		envVar["DB_PW"],
		envVar["DB_NAME"],
		envVar["DB_PORT"])

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		log.Panicf("%s", err)
	}

	// create repositories
	repo := repository.NewWaterLevelRepository(db)

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

		var msg model.RedisMsg
		err = json.Unmarshal([]byte(msgStr.Payload), &msg)
		if err != nil {
			log.Printf("error unmarshalling json: %s", err)
		}

		// create record
		newRec := model.WaterStatusRecord{
			HasWater:  msg.SensorLevel == "high",
			CreatedAt: msg.CreatedAt,
		}

		rec, err := repo.CreateWaterLevelRecord(newRec)
		if err != nil {
			log.Printf("error creating record: %s", err)
		}

		log.Println(rec)
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

	envVariables[DB_HOST], exists = os.LookupEnv(DB_HOST)
	if !exists {
		panic("missing env var: " + DB_HOST)
	}

	envVariables[DB_USER], exists = os.LookupEnv(DB_USER)
	if !exists {
		panic("missing env var: " + DB_USER)
	}

	envVariables[DB_PW], exists = os.LookupEnv(DB_PW)
	if !exists {
		panic("missing env var: " + DB_PW)
	}

	envVariables[DB_NAME], exists = os.LookupEnv(DB_NAME)
	if !exists {
		panic("missing env var: " + DB_NAME)
	}

	envVariables[DB_PORT], exists = os.LookupEnv(DB_PORT)
	if !exists {
		panic("missing env var: " + DB_PORT)
	}

	return envVariables
}
