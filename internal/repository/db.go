package repository

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
	"sync"
)

var (
	dbInstance *sql.DB
	once       sync.Once
	dbErr      error
)

// GetDB возвращает Singleton-подключение к SQLite
func GetDB() (*sql.DB, error) {
	once.Do(func() {
		// В SQLite просто указываем путь к файлу базы данных
		// TODO: вынести в конфиг
		dbPath := "./db.sqlite"

		// TODO: вынести ошиби
		dbInstance, dbErr = sql.Open("sqlite", dbPath)
		if dbErr != nil {
			dbErr = fmt.Errorf("ошибка открытия SQLite: %w", dbErr)
			return
		}

		// Проверим соединение
		if err := dbInstance.Ping(); err != nil {
			dbInstance.Close()
			dbErr = fmt.Errorf("не удалось подключиться к SQLite: %w", err)
			return
		}

		// TODO: сделать логирование
		log.Println("✅ Подключение к SQLite установлено")
	})

	return dbInstance, dbErr
}

// CloseDB закрывает SQLite соединение
func CloseDB() {
	if dbInstance != nil {
		dbInstance.Close()
		log.Println("🔌 SQLite подключение закрыто")
	}
}
