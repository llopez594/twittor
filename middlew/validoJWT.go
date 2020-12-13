package middlew

import (
	"github.com/llopez594/twittor/routers"
	"net/http"
)

/*ValidoJWT permite validar el JWT que nos viene de la peticion*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err :=routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token ! "+err.Error(), http.StatusBadRequest)
		}
		next.ServeHTTP(w, r)
	}
}
