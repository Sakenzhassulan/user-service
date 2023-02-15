package routes

import (
	"github.com/Sakenzhassulan/user-service/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request, dbi *db.DB) {
	email := chi.URLParam(r, "email")
	res, err := dbi.GetUserByEmail(email)
	if err != nil {
		render.JSON(w, r, ErrorResponse{
			Error:  err.Error(),
			Status: 404,
		})
		return
	}
	render.JSON(w, r, res)
}
