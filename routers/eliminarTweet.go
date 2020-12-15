package routers

import (
	"github.com/llopez594/twittor/bd"
	"net/http"
)

/*EliminarTweet permite borrar un Tweet determinado*/
func EliminarTweet(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
