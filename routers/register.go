package routers

import (
	"encoding/json"
	"net/http"
	"github.com/esneyder/practica-golan/bd"
	"github.com/esneyder/practica-golan/models"
)

/*Register register from users*/
func Register(w http.ResponseWriter,r *http.Request){
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err !=  nil{
		http.Error(w,"Error al procesar datos recibidos" + err.Error(),400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w,"Email de usuario es  requerido",409)
		return
	}
	if len(user.Password) < 6{
		http.Error(w,"Debe espesificar una contraseña de al  menos 6 caracteres",409)
		return
	}
	_,exist,_ := bd.CheckExistUser(user.Email)
	if  exist {
		http.Error(w,"Ya existe  un usuario con este email",409)
		return
	}

	_,status,err := bd.InsertRegister(user)

	if err != nil  {
		http.Error(w,"Error al intentar crear  el usuario " + err.Error(),409)
		return
	}

	if !status {
		http.Error(w,"No se ha logrado crear  el usuario " ,409)
		return
	} 

	w.WriteHeader(http.StatusCreated)
}