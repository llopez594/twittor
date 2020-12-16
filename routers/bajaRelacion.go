package routers

import (
	"github.com/llopez594/twittor/bd"
	"github.com/llopez594/twittor/models"
	"net/http"
)

/*BajaRelacion realiza el borrado de la relacion entre usuarios*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al borrar relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado borrar la relacion", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

