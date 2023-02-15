package handlers

import (
	"context"
	"fmt"
	"github.com/Sakenzhassulan/user-service/internal/pb"
	"github.com/go-chi/render"
	"net/http"
)

func GenerateSalt(w http.ResponseWriter, r *http.Request, svc pb.SaltServiceClient) {
	res, err := svc.GenerateSalt(context.Background(), &pb.GenerateSaltRequest{})
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
	}
	render.JSON(w, r, res)
}
