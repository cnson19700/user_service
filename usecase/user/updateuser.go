package user

import (
	"context"
	"log"
	"mime/multipart"
	"strconv"
	"strings"

	checkform "github.com/cnson19700/book_service/package/checkForm"
	imgvalid "github.com/cnson19700/book_service/package/fileValid"
	authError "github.com/cnson19700/auth_service/util/myerror"
	"github.com/cnson19700/pkg/middleware"
	"github.com/cnson19700/user_service/model"
	userError "github.com/cnson19700/user_service/util/myerror"
)

type UpdateRequest struct {
	Form *multipart.Form
}

func (u *Usecase) UpdateUser(ctx context.Context, req UpdateRequest) (*model.User, error) {
	payload := middleware.GetClaim(ctx)
	id := payload.UserID

	user, err := u.userRepo.GetById(ctx, id)
	if err != nil {
		return &model.User{}, userError.ErrGetUser(err)
	}

	if len(req.Form.Value["email"]) != 0 {
		formEmail := req.Form.Value["email"][0]
		if user.Email != formEmail {
			isMail, email := checkform.CheckFormatValue("email", formEmail)
			if !isMail {
				return nil, authError.ErrEmailFormat(nil)
			}

			if u.userRepo.CheckEmailExist(ctx, email) {
				return nil, authError.ErrEmailExist(nil)
			}
			user.Email = email
		}
	}

	if len(req.Form.Value["full_name"]) != 0 {
		formName := req.Form.Value["full_name"][0]
		if user.Email != formName {
			isName, fullname := checkform.CheckFormatValue("full_name", formName)
			if !isName {
				return &model.User{}, authError.ErrFullNameFormat(nil)
			}
			user.FullName = fullname
		}
	}

	if len(req.Form.Value["age"]) != 0 {
		formAge := req.Form.Value["age"][0]
		if user.Email != formAge {
			isAge, ageStr := checkform.CheckFormatValue("age", formAge)
			if !isAge {
				return &model.User{}, userError.ErrAgeFormat(nil)
			}
			age, _ := strconv.Atoi(ageStr)
			user.Age = age
		}
	}

	if len(req.Form.File["avatar"]) != 0 {
		file := req.Form.File["avatar"][0]
		var initAva = "blank.png"
		var pathFile = "/images/avatar"

		fileType, err := imgvalid.CheckImage(file)
		if err != nil {
			return &model.User{}, err
		}

		initAva = strconv.FormatInt(user.ID, 10) + "." + strings.Split(fileType, "/")[1]

		err = imgvalid.CopyFile(file, initAva, "."+pathFile)
		if err != nil {
			log.Fatal(err)
		}

		user.Avatar = pathFile + initAva
	}

	res, err := u.userRepo.Update(ctx, user)
	if err != nil {
		return &model.User{}, userError.ErrUpdateUser(nil)
	}

	return res, nil
}
