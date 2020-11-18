package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"reviews/graph/generated"
	"reviews/graph/model"
)

func (r *entityResolver) FindProductByID(ctx context.Context, id string) (*model.Product, error) {
	product := r.reviewUC.GetProductByID(id)
	return &model.Product{
		ID: product.ID,
	}, nil
}

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	user := r.reviewUC.GetUserByID(id)
	return &model.User{
		ID: user.ID,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
