package repository

import (
	"github.com/PlantWaterMe/GardenMonitor/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IWaterLevelInterface interface {
	CreateWaterLevelRecord(model.WaterStatusRecord) (model.WaterStatusRecord, error)
	GetWaterLevelRecordById(uuid.UUID) (model.WaterStatusRecord, error)
	GetWaterLevelRecords(map[string]interface{}) ([]model.WaterStatusRecord, error)
}

type WaterLevelRepository struct {
	db *gorm.DB
}

func NewWaterLevelRepository(db *gorm.DB) *WaterLevelRepository {
	return &WaterLevelRepository{
		db: db,
	}
}

func (w *WaterLevelRepository) CreateWaterLevelRecord(record model.WaterStatusRecord) (model.WaterStatusRecord, error) {
	result := w.db.Create(&record)
	if result.Error != nil {
		return model.WaterStatusRecord{}, result.Error
	}

	return record, nil
}

func (w *WaterLevelRepository) GetWaterLevelRecordById(id uuid.UUID) (model.WaterStatusRecord, error) {
	var record model.WaterStatusRecord

	result := w.db.First(&record, id)
	if result.Error != nil {
		return model.WaterStatusRecord{}, result.Error
	}

	return record, nil
}

func (w *WaterLevelRepository) GetWaterLevelRecords(query map[string]interface{}) ([]model.WaterStatusRecord, error) {
	var records []model.WaterStatusRecord

	q := w.db
	if query["after"] != nil {
		q = q.Where("created_at > ?", query["after"])
	}
	if query["before"] != nil {
		q = q.Where("created_at < ?", query["before"])
	}
	result := q.Find(&records)
	if result.Error != nil {
		return []model.WaterStatusRecord{}, result.Error
	}

	return records, nil
}
