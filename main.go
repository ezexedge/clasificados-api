package main

import (
	"log"

	"github.com/ezexedge/clasificados-2/bd"
	"github.com/ezexedge/clasificados-2/handlers"
)

func main() {

	//	fmt.Println(usuario)
	if bd.ConectarBD() == 0 {
		log.Fatal("sin conexion a la bd")
		return
	}
	handlers.Manejadores()
}
