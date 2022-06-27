package repository

import (
	"github.com/yusharnadi/go-toko/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Insert(product *entity.Product) error
	GetAll() (*[]entity.Product, error)
	FindId(id int) (entity.Product, error)
	Update(product *entity.Product, id int) error
	Delete(id int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Insert(product *entity.Product) error {
	err := r.db.Create(&product)
	return err.Error
}

func (r *productRepository) GetAll() (*[]entity.Product, error) {
	var product *[]entity.Product
	err := r.db.Find(&product)
	return product, err.Error
}

func (r *productRepository) FindId(id int) (entity.Product, error) {
	var product entity.Product
	err := r.db.Find(&product, id)
	return product, err.Error
}

func (r *productRepository) Update(product *entity.Product, id int) error {
	var Product *entity.Product
	err := r.db.Model(&Product).Where("id", id).Updates(product)
	return err.Error
}

func (r *productRepository) Delete(id int) error {
	var Product *entity.Product
	err := r.db.Delete(&Product, id)
	return err.Error
}
