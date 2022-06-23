package service

import (
	"github.com/yusharnadi/go-toko/entity"
	"github.com/yusharnadi/go-toko/model"
	"github.com/yusharnadi/go-toko/repository"
)

type productService struct {
	productRepository repository.ProductRepository
}

type ProductService interface {
	Insert(CreateProductRequest *model.CreateProductRequest) error
	GetAll() (*[]entity.Product, error)
	FindId(id int) (entity.Product, error)
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{productRepository}
}

func (s *productService) Insert(CreateProductRequest *model.CreateProductRequest) error {
	product := entity.Product{
		Name:  CreateProductRequest.Name,
		Price: CreateProductRequest.Price,
		Stock: CreateProductRequest.Stock,
	}

	err := s.productRepository.Insert(&product)
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
