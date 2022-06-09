package models

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	ID       uint
	Name     string
	NickName string
	jwt.StandardClaims
}
