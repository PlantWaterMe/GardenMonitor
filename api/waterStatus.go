package api

import (
	"log"
	"net/http"

	"github.com/PlantWaterMe/GardenMonitor/model"
	"github.com/gin-gonic/gin"
)

func (a *Api) CreateWaterLevelRecord(c *gin.Context) {
	var requestBody model.WaterStatusRecordRequest

	if err := c.BindJSON(&requestBody); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "invalid request params")
	}

	record := model.NewWaterStatusRecordFromRequest(requestBody)

	created, err := a.repo.CreateWaterLevelRecord(record)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	log.Printf("created: %+v", created)
	c.JSON(201, created)
}

func (a *Api) GetWaterLevelRecords(c *gin.Context) {
	var records []model.WaterStatusRecord
	dbQueries := make(map[string]interface{})

	reqQueries := c.Request.URL.Query()
	for key, value := range reqQueries {
		dbQueries[key] = value[0]
	}

	log.Printf("req queries: %+v", reqQueries)
	log.Printf("db queries: %+v", dbQueries)

	records, err := a.repo.GetWaterLevelRecords(dbQueries)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(200, records)
}
