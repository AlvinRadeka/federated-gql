package usecase

import "products/business/entity"

// ProductItf ...
type ProductItf interface {
	GetProductInfo(productID string) entity.Product
	GetAllProducts() []entity.Product
}

type productUC struct{}

// InitUsecase ...
func InitUsecase() ProductItf {
	return &productUC{}
}

var products = []entity.Product{
	{
		ID:    "product-001",
		Name:  "Faber-Castell Utensils",
		Price: 100000,
	},
	{
		ID:    "product-002",
		Name:  "IKEA Furnitures",
		Price: 1500000,
	},
	{
		ID:    "product-003",
		Name:  "Modena Kitchenwares",
		Price: 5000000,
	},
}

func (p *productUC) GetProductInfo(productID string) entity.Product {
	for i := range products {
		if products[i].ID == productID {
			return products[i]
		}
	}

	return entity.Product{
		ID:   "no product found",
		Name: "no product found",
	}
}

func (p *productUC) GetAllProducts() []entity.Product {
	return products
}
