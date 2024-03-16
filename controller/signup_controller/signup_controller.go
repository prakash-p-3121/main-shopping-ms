package signup_controller

import "github.com/prakash-p-3121/restlib"

type SignupController interface {
	Signup(restCtx restlib.RestContext)
}
