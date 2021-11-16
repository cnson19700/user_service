package user

import (
	"github.com/cnson19700/user_service/usecase"
	"github.com/cnson19700/user_service/usecase/user"
	"github.com/labstack/echo/v4"
)

type Route struct {
	userUseCase user.IUsecase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{
		userUseCase: useCase.User,
	}
	group.GET("/me", r.GetMe)
	group.PUT("/change-password", r.UpdatePassword)
	group.POST("/update", r.UpdateUser)
	group.POST("/login", r.Login)
}
