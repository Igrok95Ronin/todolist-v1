INSERT OR IGNORE INTO users (user_name, email, password_hash) VALUES
('User1', 'user1@example.com', 'hash1'),
('User2', 'user2@example.com', 'hash2'),
('User3', 'user3@example.com', 'hash3');

INSERT OR IGNORE INTO all_notes (note, completed, user_id) VALUES
('Note 1', false, 1),
('Note 2', true, 2),
('Note 3', false, 3),
('Note 4', true, 1),
('Note 5', false, 2);
