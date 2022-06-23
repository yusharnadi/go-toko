package repository

import (
	"github.com/yusharnadi/go-toko/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Insert(product entity.Product) (entity.Product, error)
	GetAll() ([]entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Insert(product entity.Product) (entity.Product, error) {
	err := r.db.Create(&product)
	return product, err.Error
}

func (r *productRepository) GetAll() ([]entity.Product, error) {
	var product []entity.Product
	err := r.db.Find(&product)
	return product, err.Error
}
