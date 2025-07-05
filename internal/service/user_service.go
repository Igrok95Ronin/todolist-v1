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
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"regexp"
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
	if err := ValidateEmail(email); err != nil {
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
	hashedPassword, err := HashPassword(password)
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

// HashPassword - хеширует пароль с помощью bcrypt (с cost = bcrypt.DefaultCost).
func HashPassword(password string) (string, error) {
	// bcrypt.GenerateFromPassword вернёт хеш пароля.
	// bcrypt.DefaultCost по умолчанию равен 10 (можно увеличить, чтобы усложнить подбор).
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Проверка валидности email
func ValidateEmail(email string) error {
	// Проверка длины email
	if len(email) < 4 {
		return fmt.Errorf("%w", httperror.ErrEmailTooShort)
	}

	// Проверка наличия символа "@" в email
	if !strings.Contains(email, "@") {
		return fmt.Errorf("%w", httperror.ErrEmailMissingAt)
	}

	// Проверка позиции символа "@" (не должен быть первым или последним символом)
	if strings.HasPrefix(email, "@") || strings.HasSuffix(email, "@") {
		return fmt.Errorf("%w", httperror.ErrEmailInvalidAtPos)
	}

	// Дополнительно: базовая проверка формата email с помощью регулярного выражения
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return fmt.Errorf("%w: %w", httperror.ErrEmailRegexCheckFail, err)
	}
	if !matched {
		return fmt.Errorf("%w", httperror.ErrEmailRegexMismatch)
	}

	return nil
}
