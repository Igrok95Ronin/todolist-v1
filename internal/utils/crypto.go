package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword - хеширует пароль с помощью bcrypt (с cost = bcrypt.DefaultCost).
func HashPassword(password string) (string, error) {
	// bcrypt.GenerateFromPassword вернёт хеш пароля.
	// bcrypt.DefaultCost по умолчанию равен 10 (можно увеличить, чтобы усложнить подбор).
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
