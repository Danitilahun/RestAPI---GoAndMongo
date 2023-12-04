package repository

import (
	"context"
	"github.com/Danitilahun/GolangAndMongoDB-REST-API/model"
)

type Repository interface {
	GetUser(ctx context.Context, email string) (model.User, error)
	CreateUser(ctx context.Context, in model.User) (model.User, error)
	UpdateUser(ctx context.Context, in model.User) (model.User, error)
	DeleteUser(ctx context.Context, email string) error
}
