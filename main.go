package main

import (
	"github.com/Dvizio/BackendAPI/tree/main/controllers/campaigncontroller"
	"github.com/Dvizio/BackendAPI/tree/main/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/campaigns", campaigncontroller.Index)
	r.GET("/api/campaign/:id", campaigncontroller.Show)
	r.GET("/api/campaign/click_through", campaigncontroller.MaxClickThrough)
	r.GET("/api/campaign/conversion", campaigncontroller.MaxConversion)
	r.GET("/api/campaign/nilai_akhir", campaigncontroller.MaxNilaiAkhir)
	r.POST("/api/campaign", campaigncontroller.Create)
	// r.POST("/api/campaign/nilai_akhir", campaigncontroller.CreateNilaiAkhir)
	r.PUT("/api/campaign/:id", campaigncontroller.Update)
	r.DELETE("/api/campaign", campaigncontroller.Delete)

	r.Run()
}
