package authhandler

import "github.com/umardev500/eschool/domain"

type authHandler struct{}

func NewAuthHandler() domain.AuthHandler {
	return &authHandler{}
}
