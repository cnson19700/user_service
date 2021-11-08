package myerror

import (
	"net/http"

	"github.com/cnson19700/pkg/apperror"
)

func ErrEmailFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100010,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Email Format is not valid",
		Message:   "Email Format is not valid",
	}
}

func ErrFullNameFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100011,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Full Name Format is not valid",
		Message:   "Full Name Format is not valid",
	}
}

func ErrPasswordFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100012,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Password Format is not valid",
		Message:   "Password Format is not valid",
	}
}

func ErrLogin(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100020,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Fail to login",
		Message:   "Fail to login",
	}
}

func ErrValidate(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100030,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Email or Password is wrong",
		Message:   "Email or Password is wrong",
	}
}

func ErrToken(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100030,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Generate token fail",
		Message:   "Generate token fail",
	}
}

//register fail
func ErrHashPassword(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100040,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Hash password fail",
		Message:   "Hash password fail",
	}
}

func ErrEmailExist(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100050,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "The email is already existed",
		Message:   "The email is already existed",
	}
}

func ErrImage(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100060,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "File is not image",
		Message:   "File is not image",
	}
}
