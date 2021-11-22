package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ezexedge/clasificados-2/models"
)

func GeneroJWT(t models.UsuariosRegistro) (string, error) {

	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"nombre":    t.Nombre,
		"apellidos": t.Apellido,
		"edad":      t.Edad,
		"id":        t.ID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
