package routes

import (
	"encoding/json"
	"github.com/Sakenzhassulan/user-service/db"
	"github.com/Sakenzhassulan/user-service/internal/pb"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int32  `json:"status"`
}

func CreateUser(w http.ResponseWriter, r *http.Request, svc pb.SaltServiceClient, validate *validator.Validate, dbi *db.DB) {
	body := CreateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		render.JSON(w, r, ErrorResponse{
			Error:  err.Error(),
			Status: http.StatusUnprocessableEntity,
		})
		return
	}
	if err := validate.Struct(body); err != nil {
		render.JSON(w, r, ErrorResponse{
			Error:  err.Error(),
			Status: http.StatusUnprocessableEntity,
		})
		return
	}
	if isExists, _ := dbi.IsUserExists(body.Email); isExists {
		render.JSON(w, r, ErrorResponse{
			Error:  "Email already exists",
			Status: http.StatusConflict,
		})
		return
	}
	res, err := dbi.CreateUser(svc, body.Email, body.Password)
	if err != nil {
		render.JSON(w, r, ErrorResponse{
			Error:  err.Error(),
			Status: http.StatusInternalServerError,
		})
		return
	}
	render.JSON(w, r, res)
}
