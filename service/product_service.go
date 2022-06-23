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
	Insert(CreateProductRequest model.CreateProductRequest) (entity.Product, error)
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	return &productService{productRepository}
}

func (s *productService) Insert(CreateProductRequest model.CreateProductRequest) (entity.Product, error) {
	product := entity.Product{
		Name:  CreateProductRequest.Name,
		Price: CreateProductRequest.Price,
		Stock: CreateProductRequest.Stock,
	}

	newProduct, err := s.productRepository.Insert(product)
	return newProduct, err
}
