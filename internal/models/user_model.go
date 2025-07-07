package models

import (
	dto "github.com/Igrok95Ronin/todolist-v1.git/internal/dto/request"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// Users Структура для таблицы users
type Users struct {
	id           int64     // Первичный ключ
	userName     string    // Уникальное имя пользователя
	email        string    // Уникальный email
	passwordHash string    // Хеш пароля
	refreshToken string    // Токен обновления (может быть NULL)
	createdAt    time.Time // Дата создания
}

// MyClaims - своя структура для claim'ов JWT, включающая стандартные поля jwt.RegisteredClaims
// и ID пользователя (UserID), чтобы знать, кому принадлежит токен.
type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

// NewUser Конструктор
func NewUserFull(id int64, userName, email, passwordHash, refreshToken string, createdAt time.Time) *Users {
	u := &Users{}
	u.SetID(id)
	u.SetUserName(userName)
	u.SetEmail(email)
	u.SetPasswordHash(passwordHash)
	u.SetRefreshToken(refreshToken)
	u.SetCreatedAt(createdAt)
	return u
}

// NOTE: Set Сеттеры с логикой
func (u *Users) SetID(id int64) {
	u.id = id
}

func (u *Users) SetUserName(userName string) {
	u.userName = userName
}

func (u *Users) SetEmail(email string) {
	u.email = email
}

func (u *Users) SetPasswordHash(passwordHash string) {
	u.passwordHash = passwordHash
}

func (u *Users) SetRefreshToken(refreshToken string) {
	u.refreshToken = refreshToken
}

func (u *Users) SetCreatedAt(createdAt time.Time) {
	u.createdAt = createdAt
}

// NOTE: Get Геттеры
func (u *Users) ID() int64 {
	return u.id
}

func (u *Users) UserName() string {
	return u.userName
}

func (u *Users) Email() string {
	return u.email
}

func (u *Users) PasswordHash() string {
	return u.passwordHash
}

func (u *Users) RefreshToken() string {
	return u.refreshToken
}

func (u *Users) CreatedAt() time.Time {
	return u.createdAt
}

// NewUserFromDTO функция-конвертер
func NewUserFromDTO(dto dto.RegisterRequest, hashedPassword string) *Users {
	u := &Users{}

	u.SetUserName(dto.UserName)
	u.SetEmail(dto.Email)
	u.SetPasswordHash(hashedPassword)
	u.SetCreatedAt(time.Now())

	return u
}
