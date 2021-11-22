package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ezexedge/clasificados-2/bd"
	"github.com/ezexedge/clasificados-2/models"
)

func VerPefil(w http.ResponseWriter, r *http.Request) {

	usuario := &models.Usuario{}
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar un paramatero id", http.StatusBadRequest)
		return
	}

	if result := bd.DB.Where("id = ?", ID).First(usuario); result.RowsAffected == 0 {
		http.Error(w, "el usuario no existe", 400)
		return
	}

	usuario.Password = ""

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}
