package beneficios

import (
	"errors"

	"beneficios.com/internal/entities"
	"beneficios.com/pkg/config"
)

func DeleteBeneficios(id string) error {
	db := config.DB
	var beneficio entities.Beneficios

	// Verificar se o benefício existe
	if err := db.Where("ID = ?", id).First(&beneficio).Error; err != nil {
		return errors.New("Beneficio não encontrado")
	}

	// Excluir o benefício
	if err := db.Delete(&beneficio).Error; err != nil {
		return errors.New("Erro ao excluir o beneficio")
	}

	return nil
}
