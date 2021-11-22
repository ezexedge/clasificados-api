package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Email    string `bson:"email" json:"email,omitempty"`
	Nombre   string `bson:"nombre" json:"nombre,omitempty"`
	Apellido string `bson:"apellido" json:"apellido,omitempty"`
	Edad     string `bson:"edad" json:"edad,omitempty"`
	Password string `bson:"password" json:"password,omitempty"`
	Foto     string `bson:"foto" json:"foto,omitempty"`
}
