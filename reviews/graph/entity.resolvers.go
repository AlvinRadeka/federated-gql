package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/alvinradeka/federated-gql/reviews/business/usecase"
	"github.com/alvinradeka/federated-gql/reviews/graph/generated"
	"github.com/alvinradeka/federated-gql/reviews/graph/model"
)

func (r *entityResolver) FindProductByID(ctx context.Context, id string) (*model.Product, error) {
	log.Println("productbyID with id: ", id)
	product := r.reviewUC.GetProductByID(id)

	reviewForProduct := r.reviewUC.GetReviewByProductID(id)
	var reviews []*model.Review
	for i := range reviewForProduct {
		review := &model.Review{
			Body: reviewForProduct[i].Body,
			Product: &model.Product{
				ID: reviewForProduct[i].ProductID,
			},
			Reviewer: &model.User{
				ID: reviewForProduct[i].ReviewerID,
			},
		}

		reviews = append(reviews, review)
	}

	return &model.Product{
		ID:      product.ID,
		Reviews: reviews,
	}, nil
}

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	log.Println("userbyID with id: ", id)
	user := r.reviewUC.GetUserByID(id)

	reviewForUser := r.reviewUC.GetReviewByUserID(id)
	var reviews []*model.Review
	for i := range reviewForUser {
		review := &model.Review{
			Body: reviewForUser[i].Body,
			Product: &model.Product{
				ID: reviewForUser[i].ProductID,
			},
			Reviewer: &model.User{
				ID: reviewForUser[i].ReviewerID,
			},
		}

		reviews = append(reviews, review)
	}

	return &model.User{
		ID: user.ID,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver {
	return &entityResolver{
		&Resolver{
			reviewUC: usecase.InitUsecase(),
		},
	}
}

type entityResolver struct{ *Resolver }
