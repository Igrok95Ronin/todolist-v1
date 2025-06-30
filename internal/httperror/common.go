package httperror

import "errors"

var (

	//////////////////////////////////////
	// 📦 Ошибки конфигурации и БД
	//////////////////////////////////////

	ErrDBOpen         = errors.New("не удалось открыть базу данных")
	ErrDBPing         = errors.New("не удалось подключиться к базе данных")
	ErrConfigEmpty    = errors.New("конфигурация отсутствует")
	ErrUserSaveFailed = errors.New("ошибка при сохранении пользователя")

	//////////////////////////////////////
	// 📡 HTTP-статусы и стандартные ошибки
	//////////////////////////////////////

	ErrBadRequest          = errors.New("некорректные данные")                      // 400
	ErrUnauthorized        = errors.New("необходима авторизация")                   // 401
	ErrForbidden           = errors.New("доступ запрещён")                          // 403
	ErrNotFound            = errors.New("пользователь не найден")                   // 404
	ErrUserExists          = errors.New("пользователь уже существует")              // 409
	ErrUnprocessableEntity = errors.New("некорректные данные в теле запроса")       // 422
	ErrTooManyRequests     = errors.New("слишком много запросов, попробуйте позже") // 429
	ErrServiceUnavailable  = errors.New("сервис временно недоступен")               // 503

	//////////////////////////////////////
	// 🔎 Валидация и парсинг запроса
	//////////////////////////////////////

	ErrMissingFields       = errors.New("все поля (username, email, password) обязательны")
	ErrInvalidEmailFormat  = errors.New("неверный формат email")
	ErrJSONDecode          = errors.New("некорректный JSON")
	ErrEmailTooShort       = errors.New("email должен быть не менее 4 символов")
	ErrEmailMissingAt      = errors.New("email должен содержать символ '@'")
	ErrEmailInvalidAtPos   = errors.New("email не может начинаться или заканчиваться на '@'")
	ErrEmailRegexMismatch  = errors.New("email не соответствует формату")
	ErrEmailRegexCheckFail = errors.New("ошибка при проверке формата email")

	//////////////////////////////////////
	// 👤 Регистрация / Авторизация
	//////////////////////////////////////

	ErrRegistrationDenied   = errors.New("регистрация невозможна")
	ErrRegistrationInternal = errors.New("внутренняя ошибка при регистрации")
	ErrPasswordHashing      = errors.New("ошибка при хешировании пароля")

	//////////////////////////////////////
	// ⚙️ Внутренние системные ошибки
	//////////////////////////////////////

	ErrInternalServer   = errors.New("внутренняя ошибка сервера")
	ErrResponseEncoding = errors.New("обработка ошибки ответа")
)
