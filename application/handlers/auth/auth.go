package authhandler

import "github.com/umardev500/eschool/domain"

type authHandler struct{}

func NewAuthController() domain.AuthHandler {
	return &authHandler{}
}
