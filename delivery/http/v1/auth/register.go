package auth

import (
	"errors"

	"github.com/cnson19700/pkg/apperror"
	"github.com/cnson19700/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/user_service/usecase/auth"
)

func (r *Route) Register(c echo.Context) error {
	var (
		ctx      = &utils.CustomEchoContext{Context: c}
		appError = apperror.AppError{}
		req      = auth.RegisterRequest{}
	)
	if err := c.Bind(&req); err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	res, err := r.authUseCase.Register(ctx, req)
	if err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, appError)
	}

	return utils.Response.Success(ctx, res)
}
