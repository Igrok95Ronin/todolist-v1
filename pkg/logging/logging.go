package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
	"sync"
)

///////////////////////////////////////////////////////
// 🔧 Кастомный хук для записи логов в несколько мест
///////////////////////////////////////////////////////

type writeHook struct {
	Writer    []io.Writer    // Список writer'ов, куда писать лог
	LogLevels []logrus.Level // Уровни логирования, на которые реагировать
}

// Fire вызывается для каждой записи логов — записывает строку в каждый writer
func (hook *writeHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		_, err = w.Write([]byte(line))
		if err != nil {
			return err
		}
	}
	return nil
}

// Levels возвращает список уровней, для которых активен хук
func (hook *writeHook) Levels() []logrus.Level {
	return hook.LogLevels
}

///////////////////////////////////////////////////////
// 📦 Singleton-логгер
///////////////////////////////////////////////////////

// Logger — обёртка над logrus.Entry
type Logger struct {
	*logrus.Entry
}

var (
	instance *Logger   // Единственный экземпляр логгера
	once     sync.Once // Гарантия однократной инициализации
)

// GetLogger возвращает Singleton-логгер, инициализируя его один раз
func GetLogger() *Logger {
	once.Do(func() {
		// 1. Создаём базовый logrus.Logger
		l := logrus.New()
		l.SetReportCaller(true) // Показывать файл и строку в логах

		// 2. Устанавливаем формат вывода
		l.Formatter = &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			DisableColors:   false,
			ForceColors:     true,
			PadLevelText:    true,
			DisableQuote:    true,
			CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
				filename := path.Base(frame.File)
				return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
			},
		}

		// 3. Создаём директорию для логов (если её нет)
		if err := os.MkdirAll("logs", 0644); err != nil {
			panic(fmt.Sprintf("не удалось создать директорию logs: %v", err))
		}

		// 4. Открываем файл для записи логов
		allFile, err := os.OpenFile("logs/all.logs", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
		if err != nil {
			panic(fmt.Sprintf("не удалось открыть файл логов: %v", err))
		}

		// 5. Отключаем стандартный вывод logrus
		l.SetOutput(io.Discard)

		// 6. Добавляем хук, записывающий в файл и stdout
		l.AddHook(&writeHook{
			Writer:    []io.Writer{allFile, os.Stdout},
			LogLevels: logrus.AllLevels,
		})

		// 7. Устанавливаем уровень логирования
		l.SetLevel(logrus.TraceLevel)

		// 8. Создаём обёртку Logger
		instance = &Logger{logrus.NewEntry(l)}
	})

	return instance
}

// GetLoggerWithField возвращает логгер с добавленным полем контекста
func (l *Logger) GetLoggerWithField(key string, value interface{}) *Logger {
	return &Logger{l.WithField(key, value)}
}
