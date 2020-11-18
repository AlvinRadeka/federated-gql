package usecase

import (
	accounts "github.com/alvinradeka/federated-gql/accounts/business/usecase"
	products "github.com/alvinradeka/federated-gql/products/business/usecase"
	"github.com/alvinradeka/federated-gql/reviews/business/entity"
)

// ReviewItf ...
type ReviewItf interface {
	GetAllReviews() []entity.Review
	GetProductByID(id string) entity.Product
	GetUserByID(id string) entity.User
}

type reviewUC struct {
	accountssvc accounts.UserItf
	productsvc  products.ProductItf
}

// InitUsecase ...
func InitUsecase() ReviewItf {
	return &reviewUC{
		accountssvc: accounts.InitUsecase(),
		productsvc:  products.InitUsecase(),
	}
}

var reviews = []entity.Review{
	{
		Body:       "Top-notch utensils for your writing needs.",
		ProductID:  "product-001",
		ReviewerID: "user-001",
	},
	{
		Body:       "Perfect furnitures to make your dream house.",
		ProductID:  "product-002",
		ReviewerID: "user-001",
	},
	{
		Body:       "All the tools you will ever need to cook your perfect steak.",
		ProductID:  "product-003",
		ReviewerID: "user-002",
	},
}

func (r *reviewUC) GetAllReviews() []entity.Review {
	return reviews
}

func (r *reviewUC) GetProductByID(id string) entity.Product {
	product := r.productsvc.GetProductInfo(id)
	return product
}

func (r *reviewUC) GetUserByID(id string) entity.User {
	user := r.accountssvc.GetUserInfo(id)
	return user
}
