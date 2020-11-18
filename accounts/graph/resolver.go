// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
package graph

import "github.com/alvinradeka/federated-gql/accounts/business/usecase"

// Resolver ...
type Resolver struct {
	userUC usecase.UserItf
}
