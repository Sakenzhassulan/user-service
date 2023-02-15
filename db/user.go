package db

import (
	"context"
	"github.com/Sakenzhassulan/user-service/internal/pb"
	"github.com/Sakenzhassulan/user-service/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Email    string
	Salt     string
	Password string
}

func (db *DB) CreateUser(svc pb.SaltServiceClient, email string, password string) (*User, error) {
	res, err := svc.GenerateSalt(context.Background(), &pb.GenerateSaltRequest{})
	if err != nil {
		return nil, err
	}
	hashedPassword := utils.HashPassword(password, res.Salt)
	user := &User{
		Email:    email,
		Salt:     res.Salt,
		Password: hashedPassword,
	}
	_, err = db.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DB) IsUserExists(email string) (bool, error) {
	filter := bson.M{"email": email}
	count, _ := db.Collection.CountDocuments(context.Background(), filter)
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (db *DB) GetUserByEmail(email string) (*User, error) {
	user := User{}
	filter := bson.M{"email": email}
	err := db.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
