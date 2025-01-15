package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BuildingController struct {
	bs *BuildingService
}

func NewBuildingController(bs *BuildingService) *BuildingController {
	return &BuildingController{bs: bs}
}

// CreateBuilding godoc
// @Summary      Create building
// @Tags         building
// @Accept       json
// @Produce      json
// @Param        body  body  CreateBuildingDto  true  "Building object"
// @Success      201  {object}  GetBuildingDto
// @Router       /building [post]
func (bc *BuildingController) CreateBuilding(c *gin.Context) {
	building, err := bc.bs.CreateBuilding(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, building)
}

// GetBuildings godoc
// @Summary      Get buildings by params
// @Description  Retrieve buildings filtered by city, year, and floors
// @Tags         building
// @Accept       json
// @Produce      json
// @Param        city   query  string  false  "City name"
// @Param        year   query  int     false  "Year of construction"
// @Param        floors query  int     false  "Number of floors"
// @Success      200  {array}  GetBuildingDto
// @Router       /building [get]
func (bc *BuildingController) GetBuildings(c *gin.Context) {
	building, err := bc.bs.GetBuildings(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, building)
}
