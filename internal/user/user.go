package user

import (
	"github.com/Sakenzhassulan/user-service/config"
	"github.com/Sakenzhassulan/user-service/db"
	"github.com/Sakenzhassulan/user-service/internal/pb"
	"github.com/Sakenzhassulan/user-service/internal/salt"
	"github.com/go-playground/validator/v10"
)

type ServiceClient struct {
	SaltServiceClient pb.SaltServiceClient
	Validator         *validator.Validate
	Config            *config.Config
	DBI               *db.DB
}

func NewServiceClient(config *config.Config, dbi *db.DB) (*ServiceClient, error) {
	saltService := salt.NewSaltServiceClient(config.SaltUrl)
	return &ServiceClient{
		SaltServiceClient: saltService.Client,
		Validator:         validator.New(),
		Config:            config,
		DBI:               dbi,
	}, nil
}
