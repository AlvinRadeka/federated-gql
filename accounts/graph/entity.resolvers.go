package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"accounts/graph/generated"
	"accounts/graph/model"
	"context"
	"log"
)

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	log.Println("FindUserByID with ID: " + id)
	name := "User " + id
	if id == "1234" {
		name = "Me"
	}

	return &model.User{
		ID:       id,
		Username: name,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver {
	log.Println("Entity Resolver for accounts")
	return &entityResolver{r}
}

type entityResolver struct{ *Resolver }
