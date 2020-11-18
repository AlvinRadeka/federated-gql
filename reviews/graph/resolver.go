package graph

import "github.com/alvinradeka/federated-gql/reviews/business/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver ...
type Resolver struct {
	reviewUC usecase.ReviewItf
}
