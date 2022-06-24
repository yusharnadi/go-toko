package service

import (
	"github.com/yusharnadi/go-toko/entity"
	"github.com/yusharnadi/go-toko/repository"
)

type productService struct {
	productRepository repository.ProductRepository
}

type ProductService interface {
	Insert(product *entity.Product) error
	GetAll() (*[]entity.Product, error)
	FindId(id int) (entity.Product, error)
	Update(product *entity.Product, id int) error
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{productRepository}
}

func (s *productService) Insert(product *entity.Product) error {

	err := s.productRepository.Insert(product)

	return err
}

func (s *productService) GetAll() (*[]entity.Product, error) {
	products, err := s.productRepository.GetAll()

	return products, err
}

func (s *productService) FindId(id int) (entity.Product, error) {
	data, err := s.productRepository.FindId(id)

	return data, err
}

func (s *productService) Update(product *entity.Product, id int) error {
	err := s.productRepository.Update(product, id)

	return err
}
