package usecase

import "accounts/business/entity"

// UserItf ...
type UserItf interface {
	GetUserInfo(userID string) entity.User
}

type userUC struct{}

// InitUsecase ...
func InitUsecase() UserItf {
	return &userUC{}
}

func (u *userUC) GetUserInfo(userID string) entity.User {
	if userID == "user-001" {
		return entity.User{
			ID:       "user-001",
			Name:     "Admin",
			UserType: 99,
		}
	}

	return entity.User{
		ID:       "invalid id",
		Name:     "invalid name",
		UserType: 0,
	}
}
