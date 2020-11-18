package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/alvinradeka/federated-gql/accounts/business/usecase"
	"github.com/alvinradeka/federated-gql/accounts/graph/generated"
	"github.com/alvinradeka/federated-gql/accounts/graph/model"
)

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	user := r.userUC.GetUserInfo("user-001")
	return &model.User{
		ID:       user.ID,
		Username: user.Name,
		UserType: user.UserType,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{
		&Resolver{
			userUC: usecase.InitUsecase(),
		},
	}
}

type queryResolver struct{ *Resolver }
