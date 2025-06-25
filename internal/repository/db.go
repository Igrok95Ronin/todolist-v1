package repository

import (
	"database/sql"
	"fmt"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/config"
	"github.com/Igrok95Ronin/todolist-v1.git/pkg/logging"
	_ "modernc.org/sqlite"
	"sync"
)

type DB struct {
	instance *sql.DB
	once     sync.Once
	err      error
	cfg      *config.Config
	logger   *logging.Logger
}

type DBOption func(*DB)

func NewDB(option ...DBOption) *DB {
	db := &DB{}

	for _, opt := range option {
		opt(db)
	}

	return db
}

// With Функции-опции
func WithConfig(cfg *config.Config) DBOption {
	return func(d *DB) {
		d.cfg = cfg
	}
}

func WithLogger(logger *logging.Logger) DBOption {
	return func(d *DB) {
		d.logger = logger
	}
}

// Connect возвращает Singleton-подключение к SQLite
func (d *DB) Connect() (*sql.DB, error) {
	d.once.Do(func() {
		// В SQLite просто указываем путь к файлу базы данных
		// TODO: вынести ошиби
		d.instance, d.err = sql.Open("sqlite", d.cfg.DBPath)
		if d.err != nil {
			d.err = fmt.Errorf("ошибка открытия SQLite: %w", d.err)
			return
		}

		// Проверим соединение
		if err := d.instance.Ping(); err != nil {
			d.instance.Close()
			d.err = fmt.Errorf("не удалось подключиться к SQLite: %w", err)
			return
		}

		d.logger.Info("✅ Подключение к SQLite установлено")
	})

	return d.instance, d.err
}

// CloseDB закрывает SQLite соединение
func (d *DB) Close() {
	if d.instance != nil {
		d.instance.Close()
		d.logger.Info("🔌 SQLite подключение закрыто")
	}
}
