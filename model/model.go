package model

import (
	"time"

	"github.com/google/uuid"
)

type RedisMsg struct {
	Topic       string
	SensorLevel string
	CreatedAt   time.Time
}

type WaterStatusRecordRequest struct {
	HasWater bool `json:"has_water"`
}

type WaterStatusRecord struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid();"`
	HasWater  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewWaterStatusRecordFromRequest(req WaterStatusRecordRequest) WaterStatusRecord {
	return WaterStatusRecord{
		HasWater: req.HasWater,
	}
}
