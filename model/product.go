package model

type CreateProductRequest struct {
	Name  string `form:"name"`
	Price int    `form:"price"`
	Stock int    `form:"stock"`
}

type CreateProductResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type GetProductResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}
