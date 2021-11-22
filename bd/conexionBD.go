package bd

import (
	"fmt"

	"github.com/ezexedge/clasificados-2/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MongoCN = ConectarBD()

var DB *gorm.DB

func ConectarBD() int {
	connection, err := gorm.Open(mysql.Open("root:root@tcp(localhost:8889)/test?parseTime=true"), &gorm.Config{})
	if err != nil {
		return 0
	}

	DB = connection

	connection.AutoMigrate(&models.UsuariosRegistro{})

	fmt.Println("conexion exitosa")
	return 1

}
