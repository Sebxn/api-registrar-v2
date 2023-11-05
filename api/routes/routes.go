package routes

import (
	"backend/api/handlers"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
)

func ConfigureRoutes(r *mux.Router, app *firebase.App) {
	// allowedOrigins := []string{"http://facturacion.lumonidy.studio", "http://localhost:3000"}

	// c := middleware.CorsMiddleware(allowedOrigins)
	// r.Use(c)
	r.Handle("/user/register", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterUser(w, r, app)
	})).Methods("POST")

	r.Handle("/user/login", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginUser(w, r, app)
	})).Methods("POST")

	r.Handle("/user/reset-password", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.ResetPassword(w, r)
	})).Methods("POST")

	// Agrega más configuraciones de rutas aquí si es necesario
}
