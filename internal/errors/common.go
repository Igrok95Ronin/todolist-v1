package errors

import "errors"

var (
	ErrDBOpen = errors.New("не удалось открыть базу данных")
	ErrDBPing = errors.New("не удалось подключиться к базе данных")
	//ErrConfigEmpty = errors.New("конфигурация отсутствует")
)
