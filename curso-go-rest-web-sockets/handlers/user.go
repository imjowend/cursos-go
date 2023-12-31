package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/imjowend/cursos-go/curso-go-rest-web-sockets/models"
	"github.com/imjowend/cursos-go/curso-go-rest-web-sockets/repository"
	"github.com/imjowend/cursos-go/curso-go-rest-web-sockets/server"
	"github.com/segmentio/ksuid"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignUpRequest{}
		// Aca primero el handler decodifica el cuerpo de la Request (r) y lo mete en request
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Luego genera una id random
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// crea una variable user de tipo models.User y le mete el mail y contra de la Request y la id generada
		var user = models.User{
			Email:    request.Email,
			Password: request.Password,
			Id:       id.String(),
		}
		// Verifica si hay un error con el metodo "Insert User" de la bd donde se este consultando
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Devuelve un json con el ID y el email de la persona con el metodo "Sign Up Response"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SignUpResponse{
			Id:    user.Id,
			Email: user.Email,
		})
	}
}
