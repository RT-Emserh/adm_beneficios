package beneficios

import (
	"beneficios.com/internal/entities"
	"beneficios.com/pkg/config"
)

func PostBeneficios(beneficio *entities.Beneficios) error {
	db := config.DB

	if err := db.Create(beneficio).Error; err != nil {
		return err
	}

	return nil
}
