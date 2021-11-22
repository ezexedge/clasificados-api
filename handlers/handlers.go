package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ezexedge/clasificados-2/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", routers.Signup).Methods("POST")
	router.HandleFunc("/login", routers.Signin).Methods("POST")

	//falta inicio de sesion
	//inicio de sesion es imposible como si fuese nodejs
	//iniciar sesesion desde front y verifycar el token
	//probar endpoint en front con nextjs
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "7000"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
