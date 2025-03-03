package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/alvinradeka/federated-gql/products/business/usecase"
	"github.com/alvinradeka/federated-gql/products/graph/generated"
	"github.com/alvinradeka/federated-gql/products/graph/model"
)

func (r *queryResolver) AllProducts(ctx context.Context) ([]*model.Product, error) {
	var result []*model.Product
	products := r.productUC.GetAllProducts()
	for i := range products {
		product := &model.Product{
			ID:    products[i].ID,
			Name:  products[i].Name,
			Price: int(products[i].Price),
		}
		result = append(result, product)
	}

	return result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{
		&Resolver{
			productUC: usecase.InitUsecase(),
		},
	}
}

type queryResolver struct{ *Resolver }
