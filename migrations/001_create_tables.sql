CREATE TABLE IF NOT EXISTS users (
                                     id INTEGER PRIMARY KEY AUTOINCREMENT,
                                     user_name TEXT NOT NULL UNIQUE,
                                     email TEXT NOT NULL UNIQUE,
                                     password_hash TEXT NOT NULL,
                                     refresh_token TEXT DEFAULT '',
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS all_notes (
                                         id INTEGER PRIMARY KEY AUTOINCREMENT,
                                         note TEXT NOT NULL,
                                         completed BOOLEAN DEFAULT FALSE,
                                         user_id INTEGER NOT NULL,
                                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                         FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );
