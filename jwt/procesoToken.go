package jwt

import (
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	//"github.com/francoleyes/twitterGo/bd"
	"github.com/francoleyes/twitterGo/models"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string, JWTSign string) (*models.Claim, bool, string, error){
	miClave := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("Formato de token inválido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err != nil {
		// Rutina chequea bd
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Token inválido")
	}

	return &claims, false, string(""), err
}