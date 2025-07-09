package utils

import (
	"fmt"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/httperror"
	"regexp"
	"strings"
)

// ValidateEmail Проверка валидности email
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
