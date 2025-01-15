package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BuildingService struct {
	db *gorm.DB
}

func NewBuildingService(db *gorm.DB) *BuildingService {
	return &BuildingService{db: db}
}

func (s *BuildingService) CreateBuilding(c *gin.Context) (*GetBuildingDto, error) {
	var createBuildingDto CreateBuildingDto
	if err := c.ShouldBindJSON(&createBuildingDto); err != nil {
		return nil, err
	}
	building := &Building{
		Name:   createBuildingDto.Name,
		City:   createBuildingDto.City,
		Year:   createBuildingDto.Year,
		Floors: createBuildingDto.Floors,
	}
	result := s.db.Create(&building)
	if result.Error != nil {
		return nil, result.Error
	}

	getBuildingDto := &GetBuildingDto{
		ID:        building.ID,
		CreatedAt: building.CreatedAt,
		UpdatedAt: building.UpdatedAt,
		DeletedAt: building.DeletedAt.Time,
		Name:      building.Name,
		City:      building.City,
		Year:      building.Year,
		Floors:    building.Floors,
	}

	return getBuildingDto, nil
}

func (s *BuildingService) GetBuildings(c *gin.Context) ([]GetBuildingDto, error) {
	var buildings []Building
	city := c.Query("city")
	year := c.Query("year")
	floors := c.Query("floors")

	query := s.db.Model(&buildings)

	if city != "" {
		query = query.Where("city = ?", city)
	}
	if year != "" {
		query = query.Where("year = ?", year)
	}
	if floors != "" {
		query = query.Where("floors = ?", floors)
	}
	var getBuildingDto []GetBuildingDto
	query.Find(&buildings)

	for _, building := range buildings {
		getBuildingDto = append(getBuildingDto, GetBuildingDto{
			ID:        building.ID,
			CreatedAt: building.CreatedAt,
			UpdatedAt: building.UpdatedAt,
			DeletedAt: building.DeletedAt.Time,
			Name:      building.Name,
			City:      building.City,
			Year:      building.Year,
			Floors:    building.Floors,
		})
	}

	return getBuildingDto, query.Error
}
