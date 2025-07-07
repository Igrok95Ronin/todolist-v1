package repository

import (
	"context"
	"database/sql"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/models"
)

// UserRepository - интерфейс для работы с пользователями
type UserRepository interface {
	UserExists(ctx context.Context, userName, email string) error
	Register(ctx context.Context, users *models.Users) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// UserExists проверяем есть ли пользователь в бд
func (r *userRepository) UserExists(ctx context.Context, userName, email string) error {
	query := "SELECT 1 FROM users WHERE user_name = $1 OR email = $2 LIMIT 1"
	var exists int
	return r.db.QueryRowContext(ctx, query, userName, email).Scan(&exists)
}

// Register Сохраняем пользователя в бд
func (r *userRepository) Register(ctx context.Context, users *models.Users) error {
	query := "INSERT INTO users (user_name, email, password_hash, created_at) VALUES ($1, $2, $3, $4)"
	_, err := r.db.ExecContext(ctx, query,
		users.UserName(),
		users.Email(),
		users.PasswordHash(),
		users.CreatedAt(),
	)
	return err
}
