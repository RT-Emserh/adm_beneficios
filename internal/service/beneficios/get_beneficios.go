package beneficios

import (
	"beneficios.com/internal/entities"
	"beneficios.com/pkg/config"
)

func GetBeneficios() ([]entities.Beneficios, error) {
	// refazer aqyui
	db := config.DB
	var beneficios []entities.Beneficios

	if err := db.Find(&beneficios).Error; err != nil {
		return nil, err
	}

	return beneficios, nil
}
