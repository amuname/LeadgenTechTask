package main

import (
	_ "LeadgenTechTask/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	bs := NewBuildingService(db)
	bc := NewBuildingController(bs)

	r.Use(ErrorHandler)

	r.POST("/building", bc.CreateBuilding)
	r.GET("/building", bc.GetBuildings)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
