package salt

import (
	"fmt"
	"github.com/Sakenzhassulan/user-service/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SaltServiceClient struct {
	Client pb.SaltServiceClient
}

func NewSaltServiceClient(url string) SaltServiceClient {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect: ", err)
	}
	client := SaltServiceClient{
		Client: pb.NewSaltServiceClient(cc),
	}
	return client
}
