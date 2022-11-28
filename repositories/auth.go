package repositories

import (
	"party/models"

	"gorm.io/gorm"
)

type AuthRepositories interface {
	Register(profile models.Profile) (models.Profile, error)
	Login(email string) (models.Profile, error)
}

func RepositoriesAuth(db *gorm.DB) *repositories{
	return &repositories{db}
}

func (r *repositories) Register(profile models.Profile) (models.Profile, error){
	
	err := r.db.Create(&profile).Error
	return profile, err
}

func (r *repositories) Login(email string) (models.Profile, error) {
	var profile models.Profile
	err := r.db.First(&profile, "email=?", email).Error

	return profile, err
}
