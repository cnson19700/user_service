package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/user_service/usecase"
	"github.com/user_service/usecase/auth"
)

type Route struct {
	authUseCase auth.IUsecase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{
		authUseCase: useCase.Auth,
	}

	group.POST("/register", r.Register)
	group.POST("/login", r.Login)
}
