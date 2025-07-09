package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/config"
	dto "github.com/Igrok95Ronin/todolist-v1.git/internal/dto/request"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/httperror"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/models"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/repository"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/utils"
	"html/template"
	"strings"
)

// UserService - интерфейс для работы с бизнес-логикой пользователей
type UserService interface {
	UserExists(ctx context.Context, users dto.RegisterRequest) error
}

type userService struct {
	repo repository.UserRepository
	cfg  *config.Config
}

func NewUserService(repo repository.UserRepository, cfg *config.Config) UserService {
	return &userService{
		repo: repo,
		cfg:  cfg,
	}
}

// UserExists проверяем есть ли пользователь регистрируем нового пользователя
func (s *userService) UserExists(ctx context.Context, dto dto.RegisterRequest) error {
	userName := strings.TrimSpace(dto.UserName)
	email := strings.TrimSpace(dto.Email)
	password := strings.TrimSpace(dto.Password)

	if userName == "" || email == "" || password == "" {
		return fmt.Errorf("s <- %w", httperror.ErrMissingFields)
	}

	userName = template.HTMLEscapeString(userName)
	email = template.HTMLEscapeString(email)
	password = template.HTMLEscapeString(password)

	// Проверка валидности email
	if err := utils.ValidateEmail(email); err != nil {
		return fmt.Errorf("s <- %w", err)
	}

	// UserExists проверяем есть ли пользователь в бд
	err := s.repo.UserExists(ctx, userName, email)
	if err == nil { // Если ошибки нет, значит пользователь найден
		return fmt.Errorf("s <- %w: r <- %w", httperror.ErrRegistrationDenied, httperror.ErrUserExists)
	}

	if err != sql.ErrNoRows { // Если ошибка не sql.ErrNoRows, значит это другая проблема
		return fmt.Errorf("%w: %w", httperror.ErrRegistrationInternal, err)
	}

	// Хешируем пароль
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return fmt.Errorf("%w: %w", httperror.ErrPasswordHashing, err)

	}

	// преобразуем DTO → модель
	// Создаём объект нового пользователя
	newUser := models.NewUserFromDTO(dto, hashedPassword)

	// Сохраняем пользователя
	if err = s.repo.Register(ctx, newUser); err != nil {
		return fmt.Errorf("%w: %s", httperror.ErrUserSaveFailed, err)
	}

	return nil
}
