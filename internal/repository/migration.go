package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// InitSchema выполняет все SQL-миграции из папки ./migrations
func InitSchema(db *sql.DB) error {
	migrationsPath := "./migrations"

	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		return fmt.Errorf("не удалось прочитать папку миграций: %w", err)
	}

	// Сортируем по имени файла (001, 002, ...)
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		fullPath := filepath.Join(migrationsPath, file.Name())
		sqlBytes, err := os.ReadFile(fullPath)
		if err != nil {
			return fmt.Errorf("ошибка чтения файла %s: %w", file.Name(), err)
		}

		queries := string(sqlBytes)
		if strings.TrimSpace(queries) == "" {
			continue
		}

		_, err = db.Exec(queries)
		if err != nil {
			return fmt.Errorf("ошибка выполнения миграции %s: %w", file.Name(), err)
		}

		log.Printf("✅ Выполнена миграция: %s", file.Name())
	}

	return nil
}
