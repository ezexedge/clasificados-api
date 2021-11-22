package routers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ezexedge/clasificados-2/bd"
	"github.com/ezexedge/clasificados-2/helpers"
	"github.com/ezexedge/clasificados-2/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {

	var t models.UsuariosRegistro

	usuario := &models.UsuariosRegistro{}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 400)

		return
	}

	println("eee", t.Email)

	json.Unmarshal(reqBody, &t)

	if result := bd.DB.Where("email = ?", t.Email).First(usuario); result.RowsAffected == 1 {
		http.Error(w, "el usuario existe", 400)
		return
	}

	t.Password, _ = helpers.EncriptarPassword(t.Password)

	if e := bd.DB.Create(&t).Error; e != nil {
		log.Println("no se puede crear error")
		http.Error(w, err.Error(), 400)

		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func Signin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("context-type", "application/json")

	var t models.UsuarioLogin

	usuario := &models.UsuariosRegistro{}

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña inválidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido ", 400)
		return
	}

	if result := bd.DB.Where("email = ?", t.Email).First(usuario); result.RowsAffected == 0 {
		http.Error(w, "el usuario no se encuentra registrado", 400)
		return
	}

	passwordBytes := []byte(t.Password)
	passwordBD := []byte(usuario.Password)

	respuesta := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if respuesta != nil {
		http.Error(w, "el password no coincide", 400)
		return
	}

	jwtKey, err := helpers.GeneroJWT(*usuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar general el Token correspondiente "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
