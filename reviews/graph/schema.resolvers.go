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

func (r *queryResolver) AllReviews(ctx context.Context) ([]*model.Review, error) {
	var res []*model.Review

	reviews := r.reviewUC.GetAllReviews()
	for i := range reviews {
		review := &model.Review{
			Body: reviews[i].Body,
		}
		res = append(res, review)
	}

	return res, nil
}

func (r *productResolver) Reviews(ctx context.Context, obj *model.Product) ([]*model.Review, error) {
	var res []*model.Review

	log.Println("product resolver", *obj)
	reviews := r.reviewUC.GetAllReviews()
	for _, review := range reviews {
		if review.ProductID == obj.ID {
			reviewedProduct := &model.Review{
				Body:    review.Body,
				Product: &model.Product{ID: obj.ID},
			}

			res = append(res, reviewedProduct)
		}
	}

	return res, nil
}

func (r *userResolver) Reviews(ctx context.Context, obj *model.User) ([]*model.Review, error) {
	var res []*model.Review

	log.Println("user resolver", *obj)
	reviews := r.reviewUC.GetAllReviews()
	for _, review := range reviews {
		if review.ReviewerID == obj.ID {
			reviewingUser := &model.Review{
				Body:     review.Body,
				Reviewer: &model.User{ID: obj.ID},
			}

			res = append(res, reviewingUser)
		}
	}

	return res, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{
		&Resolver{
			reviewUC: usecase.InitUsecase(),
		},
	}
}

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver {
	return &productResolver{
		&Resolver{
			reviewUC: usecase.InitUsecase(),
		},
	}
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver {
	return &userResolver{
		&Resolver{
			reviewUC: usecase.InitUsecase(),
		},
	}
}

type queryResolver struct{ *Resolver }

type productResolver struct{ *Resolver }

type userResolver struct{ *Resolver }
