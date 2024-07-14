package dependences

import (
	"log"
	"parking/internal/config"
	"parking/internal/database"

	"gorm.io/gorm"
)

type Dependencies struct {
	//LocationRepository location.Repository
}

func NewDependencies() (*Dependencies, error) {
	_, err := ConnectionDataBase()
	if err != nil {
		return nil, err
	}
	//dbRepository := location.NewRepositoryLocation(db)
	return &Dependencies{
		//LocationRepository: dbRepository,
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
