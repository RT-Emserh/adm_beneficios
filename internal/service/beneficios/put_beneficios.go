package beneficios

import (
	"errors"

	"beneficios.com/internal/entities"
	"beneficios.com/pkg/config"
)

func PutBeneficios(id string, beneficio *entities.Beneficios) error {
	db := config.DB

	var existingBeneficio entities.Beneficios
	if err := db.Where("ID = ?", id).First(&existingBeneficio).Error; err != nil {
		return errors.New("Benefício não encontrado")
	}

	if err := db.Model(&existingBeneficio).Updates(beneficio).Error; err != nil {
		return err
	}

	return nil
}
