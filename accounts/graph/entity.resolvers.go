package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/alvinradeka/federated-gql/accounts/graph/generated"
	"github.com/alvinradeka/federated-gql/accounts/graph/model"
)

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	user := r.userUC.GetUserInfo(id)
	return &model.User{
		ID:       user.ID,
		Username: user.Name,
		UserType: user.UserType,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
