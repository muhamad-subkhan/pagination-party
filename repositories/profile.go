package repositories

import (
	"party/models"

	"gorm.io/gorm"
)

type ProfileRepositories interface {
	GetProfile(ID int) (models.Profile, error)
	UpdateProfile(profile models.Profile) (models.Profile, error)
}

func RepositoriesProfile(db *gorm.DB) *repositories{
	return &repositories{db}
}

func (r *repositories) GetProfile(ID int) (models.Profile, error){
	var profile models.Profile

	err := r.db.First(&profile, ID).Error
	return profile, err
}


func (r *repositories) UpdateProfile(profile models.Profile) (models.Profile, error){
	err := r.db.Save(&profile).Error

	return profile, err
}
