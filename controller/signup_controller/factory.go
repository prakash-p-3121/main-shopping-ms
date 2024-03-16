package signup_controller

import (
	"github.com/prakash-p-3121/main-shopping-ms/controller/signup_controller/impl"
	"github.com/prakash-p-3121/main-shopping-ms/service/signup_service"
)

func NewSignupController() SignupController {
	return &impl.SignupControllerImpl{Service: signup_service.NewSignupService()}
}
