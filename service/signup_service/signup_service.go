package signup_service

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/usermodel"
)

type SignupService interface {
	Signup(req *usermodel.UserCreateReq) (*idgenmodel.IDGenResp, errorlib.AppError)
}
