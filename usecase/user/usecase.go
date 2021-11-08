package user

import (
	"github.com/user_service/repository"
	"github.com/user_service/repository/user"
)

type Usecase struct {
	userRepo user.Repository
}

func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		userRepo: repo.User,
	}
}
