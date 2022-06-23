package repository

import (
	"github.com/yusharnadi/go-toko/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Insert(product entity.Product) (entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) Insert(product entity.Product) (entity.Product, error) {
	err := r.db.Create(&product)
	return product, err.Error
}
