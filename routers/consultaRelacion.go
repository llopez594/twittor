package routers

import (
	"encoding/json"
	"github.com/llopez594/twittor/bd"
	"github.com/llopez594/twittor/models"
	"net/http"
)

/*ConsultaRelacion chequea si hay relacion entre 2 usuarios*/
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
