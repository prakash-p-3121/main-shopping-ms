package signup_service

import "github.com/prakash-p-3121/main-shopping-ms/service/signup_service/impl"

func NewSignupService() SignupService {
	return &impl.SignupServiceImpl{}
}
