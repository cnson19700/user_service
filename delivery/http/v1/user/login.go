package user

import (
	"errors"

	"github.com/cnson19700/auth_service/proto"
	"github.com/cnson19700/pkg/apperror"
	"github.com/cnson19700/pkg/utils"
	authservice "github.com/cnson19700/user_service/client/authService"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Email    string
	Password string
}

func (r *Route) Login(c echo.Context) error {
	var (
		ctx      = &utils.CustomEchoContext{Context: c}
		appError = apperror.AppError{}
		req      = LoginRequest{}
	)

	// Bind order by
	if err := c.Bind(&req); err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	reqLogin := &proto.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	client := authservice.GetClient()

	res, err := client.Login(ctx, reqLogin)

	if err != nil {
		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	return utils.Response.Success(ctx, res)
}
