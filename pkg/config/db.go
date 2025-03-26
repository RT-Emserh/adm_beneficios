package config

import (
	"log"
	"time"

	"beneficios.com/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	if DB != nil {
		return // Evitar conexão duplicada
	}

	var err error
	// Opção PreferSimpleProtocol: true desabilita o uso de prepared statements
	dsn := "user=postgres.xqkdlawjgrfptndhkvtv password=Evilasio22?? host=aws-0-sa-east-1.pooler.supabase.com port=6543 dbname=postgres sslmode=disable"
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // Desabilita cache de declarações preparadas
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("failed to get database object: %v", err)
		return
	}

	// Configurando o pool de conexões
	sqlDB.SetMaxOpenConns(25)                 // número máximo de conexões abertas
	sqlDB.SetMaxIdleConns(25)                 // número máximo de conexões ociosas
	sqlDB.SetConnMaxLifetime(time.Second * 3) // tempo máximo de vida das conexões

	// Só auto-migra a entidade Beneficios
	if err := DB.AutoMigrate(&entities.Beneficios{}); err != nil {
		log.Printf("failed to run automigrate for Beneficios: %v", err)
	} else {
		log.Printf("AutoMigrate executed successfully for Beneficios")
	}
}
