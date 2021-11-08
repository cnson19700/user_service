package user

import (
	"github.com/labstack/echo/v4"
	"github.com/user_service/usecase"
	"github.com/user_service/usecase/user"
)

type Route struct {
	userUseCase user.IUsecase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{
		userUseCase: useCase.User,
	}
	group.GET("/getme", r.GetMe)
	group.PUT("/updatepassword", r.UpdatePassword)
	group.POST("/updateuser", r.UpdateUser)
}
