package register

import (
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

type RepositoryUser interface {
	CreateUser(ctx context.Context, user *UserDataComplete) error
	GetUser(ctx context.Context, id string) (UserMotorcycles, error)
	UpdateUser(ctx context.Context, user *UserDataComplete) error
	GetAllUsers(ctx context.Context) ([]UserMotorcycles, error)
}

type Repository struct {
	DBRepository *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) Repository {
	return Repository{
		DBRepository: db,
	}
}

func (r Repository) CreateUser(ctx context.Context, userCompleted *UserDataComplete) error {
	log.Println("initial transaction create user")
	tx := r.DBRepository.WithContext(ctx).Begin()
	if tx.Error != nil {
		log.Println("failed to begin transaction: ", tx.Error)
		return tx.Error
	}

	log.Println("creating user: ", userCompleted)
	if err := tx.Create(&userCompleted.User).Error; err != nil {
		log.Println("error in repository create user. error: ", err)
		tx.Rollback()
		return errors.New("error when creating the user or the user already exists")
	}

	var existingMotorcycle Motorcycle
	err := tx.Where("plate = ? AND user_id = ?", userCompleted.Motorcycle.Plate, userCompleted.User.ID).Find(&existingMotorcycle).Error
	if err != nil || ValidMotorcycle(existingMotorcycle) {
		log.Println("Motorcycle already exists for the user: ", userCompleted.User.ID)
		tx.Rollback()
		return errors.New("motorcycle already exists for the user")
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("error checking motorcycle existence: ", err)
		tx.Rollback()
		return errors.New("error checking motorcycle existence")
	}

	log.Println("creating motorcycle: ", userCompleted.Motorcycle)
	if err := tx.Create(&userCompleted.Motorcycle).Error; err != nil {
		log.Println("error in repository create motorcycle. error: ", err)
		log.Println("transaction rollback: ", tx.Rollback().Error)
		return errors.New("error repository create the motorcycle")
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in panic of repository create user or motorcycle", r)
			log.Println("transaction rollback: ", tx.Rollback().Error)
			panic(r)
		}
	}()

	log.Println("finish transaction create user")
	if err := tx.Commit().Error; err != nil {
		log.Println("error in transaction commit: ", err)
		return err
	}
	return nil
}

func ValidMotorcycle(motorcycle Motorcycle) bool {
	return motorcycle != Motorcycle{}
}

func (r Repository) GetUser(ctx context.Context, id string) (UserMotorcycles, error) {
	var user User
	var userMotorcycles UserMotorcycles

	log.Println("fetching user with id: ", id)
	if err := r.DBRepository.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		log.Println("error in repository find user. error: ", err)
		return userMotorcycles, errors.New("error repository find the user")
	}

	var motorcycles []Motorcycle
	if err := r.DBRepository.Where("user_id = ?", user.ID).Find(&motorcycles).Error; err != nil {
		log.Println("error in repository find motorcycles for user. error: ", err)
		return userMotorcycles, errors.New("error repository find motorcycles for user")
	}

	userMotorcycles = UserMotorcycles{
		User:       user,
		Motorcycle: motorcycles,
	}

	return userMotorcycles, nil
}

func (r Repository) UpdateUser(ctx context.Context, userCompleted *UserDataComplete) error {
	log.Println("initial transaction update user")
	tx := r.DBRepository.WithContext(ctx).Begin()
	if tx.Error != nil {
		log.Println("failed to begin transaction: ", tx.Error)
		return tx.Error
	}

	log.Println("updating user: ", userCompleted.User)
	if err := tx.Save(&userCompleted.User).Error; err != nil {
		log.Println("error in repository update user. error: ", err)
		log.Println("transaction rollback: ", tx.Rollback().Error)
		return errors.New("error update the user. Error" + err.Error())
	}

	var existingMotorcycle Motorcycle
	err := tx.Where("plate = ? AND user_id = ?", userCompleted.Motorcycle.Plate, userCompleted.User.ID).First(&existingMotorcycle).Error
	if err == nil && ValidMotorcycle(existingMotorcycle) {
		userCompleted.Motorcycle.ID = existingMotorcycle.ID
		log.Println("updating motorcycle: ", userCompleted.Motorcycle)
		if err := tx.Save(&userCompleted.Motorcycle).Error; err != nil {
			log.Println("error in repository update motorcycle. error: ", err)
			log.Println("transaction rollback: ", tx.Rollback().Error)
			return errors.New("error repository update the motorcycle")
		}
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("error checking motorcycle to update: ", err)
		tx.Rollback()
		return errors.New("error checking motorcycle to update")
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("creating motorcycle: ", userCompleted.Motorcycle)
		if err := tx.Create(&userCompleted.Motorcycle).Error; err != nil {
			log.Println("error in repository create motorcycle. error: ", err)
			log.Println("transaction rollback: ", tx.Rollback().Error)
			return errors.New("error repository create the motorcycle")
		}
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in panic of repository update user", r)
			log.Println("transaction rollback: ", tx.Rollback().Error)
			panic(r)
		}
	}()

	log.Println("finish transaction update user")
	if err := tx.Commit().Error; err != nil {
		log.Println("error in transaction commit: ", err)
		return err
	}
	return nil
}

func (r Repository) GetAllUsers(ctx context.Context) ([]UserMotorcycles, error) {
	var users []User
	var userMotorcycles []UserMotorcycles

	log.Println("fetching all users")
	if err := r.DBRepository.WithContext(ctx).Find(&users).Error; err != nil {
		log.Println("error in repository find all users. error: ", err)
		return nil, errors.New("error repository find all users")
	}

	for _, user := range users {
		var motorcycles []Motorcycle
		if err := r.DBRepository.Where("user_id = ?", user.ID).Find(&motorcycles).Error; err != nil {
			log.Println("error in repository find motorcycles for user. error: ", err)
			return nil, errors.New("error repository find motorcycles for user")
		}
		userMotorcycles = append(userMotorcycles, UserMotorcycles{
			User:       user,
			Motorcycle: motorcycles,
		})
	}

	return userMotorcycles, nil
}
