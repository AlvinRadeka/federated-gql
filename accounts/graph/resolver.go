// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
package graph

import "accounts/business/usecase"

// Resolver ...
type Resolver struct {
	userUC usecase.UserItf
}
