package routers

import (
	"github.com/llopez594/twittor/bd"
	"io"
	"net/http"
	"os"
)

/*ObtenerBanner envia el banner al HTTP*/
func ObtenerBanner(w http.ResponseWriter, r *http.Request){

	ID := r.URL.Query().Get("id")
	if len(ID) > 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	Openfile, err := os.Open("uploads/banners/"+perfil.Banner)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, Openfile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}
}
