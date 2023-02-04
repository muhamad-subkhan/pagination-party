package repositories

import (
	"party/models"

	"gorm.io/gorm"
)

type ProductRepositories interface {
	GetProduct(ID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	FindProduct(limit int, page int) ([]models.Product, error)
}
func RepositoriesProduct(db *gorm.DB) *repositories {
	return &repositories{db}
}


func (r *repositories) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.First(&product, ID).Error

	return product, err
}

func (r *repositories) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}

	func (r *repositories) UpdateProduct(product models.Product) (models.Product, error) {
		err := r.db.Save(&product).Error

		return product, err
}

func (r *repositories) FindProduct(limit int, page int) ([]models.Product, error) {
	var product []models.Product
	if page <= 1 {
		page = 0
	}

	err := r.db.Limit(limit).Offset(page).Find(&product).Error

	return product, err
}