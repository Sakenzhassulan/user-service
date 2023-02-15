package user

import (
	"github.com/Sakenzhassulan/user-service/internal/salt/handlers"
	"github.com/Sakenzhassulan/user-service/internal/user/routes"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux, userClient *ServiceClient) *chi.Mux {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Post("/generate-salt", userClient.GenerateSalt)
	r.Post("/create-user", userClient.CreateUser)
	r.Get("/get-user/{email}", userClient.GetUser)
	return r
}

func (svc *ServiceClient) GenerateSalt(w http.ResponseWriter, r *http.Request) {
	handlers.GenerateSalt(w, r, svc.SaltServiceClient)
}

func (svc *ServiceClient) CreateUser(w http.ResponseWriter, r *http.Request) {
	routes.CreateUser(w, r, svc.SaltServiceClient, svc.Validator, svc.DBI)
}

func (svc *ServiceClient) GetUser(w http.ResponseWriter, r *http.Request) {
	routes.GetUser(w, r, svc.DBI)
}
