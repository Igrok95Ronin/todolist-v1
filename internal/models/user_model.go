package models

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// Структура для таблицы users
type Users struct {
	ID           int64     `json:"ID"`            // Первичный ключ
	UserName     string    `json:"username"`      // Уникальное имя пользователя
	Email        string    `json:"email"`         // Уникальный email
	PasswordHash string    `json:"password" `     // Хеш пароля
	RefreshToken string    `json:"refreshToken" ` // Токен обновления (может быть NULL)
	CreatedAt    time.Time // Дата создания
}

// MyClaims - своя структура для claim'ов JWT, включающая стандартные поля jwt.RegisteredClaims
// и ID пользователя (UserID), чтобы знать, кому принадлежит токен.
type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}
