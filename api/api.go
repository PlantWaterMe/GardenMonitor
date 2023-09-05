package api

import (
	"log"

	"github.com/PlantWaterMe/GardenMonitor/api/repository"
	"github.com/PlantWaterMe/GardenMonitor/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Api struct {
	gin  *gin.Engine
	port string
	repo repository.IWaterLevelInterface
}

func NewApi(db *gorm.DB, port string) *Api {

	var api = Api{}

	sqlDB, _ := db.DB()
	_, err := sqlDB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(model.WaterStatusRecord{})
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	v1 := r.Group("/api/v1")

	// create repositories
	repo := repository.NewWaterLevelRepository(db)

	// create handler
	v1.POST("/waterlevel", api.CreateWaterLevelRecord)
	v1.GET("/waterlevel", api.GetWaterLevelRecords)

	api.gin = r
	api.port = port
	api.repo = repo

	return &api
}

func (a *Api) Start() {
	log.Println("Start server on: ", a.gin.BasePath())

	if err := a.gin.Run(":" + a.port); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
