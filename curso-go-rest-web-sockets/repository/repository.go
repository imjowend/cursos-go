package repository

import (
	"context"

	"github.com/imjowend/cursos-go/curso-go-rest-web-sockets/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	Close() error
	InsertPost(ctx context.Context, post *models.Post) error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func Close() error {
	return implementation.Close()
}

func GetUserByEmail(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, id)
}

func InsertPost(ctx context.Context, post *models.Post) error {
	return implementation.InsertPost(ctx, post)
}
