package impl

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/main-shopping-ms/cfg"
	"github.com/prakash-p-3121/usermgtclient"
	"github.com/prakash-p-3121/usermodel"
	"log"
)

type SignupServiceImpl struct {
}

func (service *SignupServiceImpl) Signup(req *usermodel.UserCreateReq) (*idgenmodel.IDGenResp, errorlib.AppError) {
	if req == nil {
		return nil, errorlib.NewBadReqError("req-nil")
	}
	userMgtMsCfg, err := cfg.GetMsConnectionCfg("usermgtms")
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}

	userMgtClient := usermgtclient.NewUserMgtClient(userMgtMsCfg.Host, uint(userMgtMsCfg.Port))
	idResp, appErr := userMgtClient.UserCreate(req)
	if appErr != nil {
		return nil, appErr
	}
	log.Println("idGenResp-Service=", idResp)

	return idResp, nil
}
