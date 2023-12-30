package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/imjowend/cursos-go/curso-go-rest-web-sockets/models"
	_ "github.com/lib/pq"
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

func (repository *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	rows, err := repository.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *PostgresRepository) Close() error {
	return repository.db.Close()
}
