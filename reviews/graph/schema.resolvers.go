package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"reviews/graph/generated"
	"reviews/graph/model"
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
