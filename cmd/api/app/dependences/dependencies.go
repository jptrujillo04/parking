package dependences

import (
	"log"
	"parking/internal/config"
	"parking/internal/database"
	"parking/internal/register"

	"gorm.io/gorm"
)

type Dependencies struct {
	UserRepository register.Repository
}

func NewDependencies() (*Dependencies, error) {
	db, err := ConnectionDataBase()
	if err != nil {
		return nil, err
	}
	repositoryUser := register.NewRepositoryUser(db)
	return &Dependencies{
		UserRepository: repositoryUser,
	}, nil
}

func ConnectionDataBase() (*gorm.DB, error) {
	dbConfig := config.ReadDBConfig()
	db, err := database.ConnectDB(dbConfig)
	if err != nil {
		log.Println("Error conectando a la base de datos:", err)
	}
	return db, nil
}
