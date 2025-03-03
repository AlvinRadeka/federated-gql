package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/alvinradeka/federated-gql/products/business/usecase"
	"github.com/alvinradeka/federated-gql/products/graph/generated"
	"github.com/alvinradeka/federated-gql/products/graph/model"
)

func (r *entityResolver) FindProductByID(ctx context.Context, id string) (*model.Product, error) {
	product := r.productUC.GetProductInfo(id)
	return &model.Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: int(product.Price),
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver {
	return &entityResolver{
		&Resolver{
			productUC: usecase.InitUsecase(),
		},
	}
}

type entityResolver struct{ *Resolver }
