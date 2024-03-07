package http

import "household-dashboard/src/controller/user"

type Controllers struct {
	UserController *user.UserController
}

func NewUserHttpHandler(controllers *Controllers) *user.UserController {
	return controllers.UserController
}
