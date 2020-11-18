package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"accounts/graph/generated"
	"accounts/graph/model"
	"context"
)

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	user := r.userUC.GetUserInfo("user-001")
	return &model.User{
		ID:       user.ID,
		Username: user.Name,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
