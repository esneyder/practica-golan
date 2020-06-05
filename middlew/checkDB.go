package middlew

import (
	"net/http"

	"github.com/esneyder/practica-golan/bd"
)

/*CheckBD check database*/
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnect() == 0 {
			http.Error(w, "Conexión perdida con la base  de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
