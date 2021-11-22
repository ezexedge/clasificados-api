package models

type UsuarioLogin struct {
	Email    string `bson:"email" json:"email,omitempty"`
	Password string `bson:"password" json:"password,omitempty"`
}
