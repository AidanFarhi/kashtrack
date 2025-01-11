DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS expense;

CREATE TABLE user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE expense (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    timestamp TEXT DEFAULT (DATETIME('now')),
    amount DOUBLE,
    category TEXT,
    FOREIGN KEY (user_id) REFERENCES user(id)
);
