package database

import (
	"context"
	"database/sql"

	"github.com/imjowend/cursos-go/curso-go-rest-web-sockets/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repository *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repository.db.ExecContext(ctx, "INSERT INTO users (email, password) VALUES  ($1, $2)", user.Email, user.Password)

	return err
}

func (repository *PostgresRepository) GetUserById(ctx context.Context, id int64) {

}
