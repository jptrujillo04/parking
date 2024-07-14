package database

import (
	"fmt"
	"log"
	"parking/internal/config"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg config.DBConfig) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error al obtener el objeto *sql.DB:", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Error al hacer ping a la base de datos:", err)
	}

	return db, nil
}
