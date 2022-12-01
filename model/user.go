package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id          int    `json:"id"`
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	Question    string `json:"question"`
	Answer      string `json:"answer"`
	TokenString string `json:"tokenString"`
	Person      string `json:"person"`
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
