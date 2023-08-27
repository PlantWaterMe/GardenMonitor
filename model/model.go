package model

import (
	"time"
)

type RedisMsg struct {
	Topic       string
	SensorLevel string
	CreatedAt   time.Time
}
