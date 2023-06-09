CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nick_name VARCHAR(32) NOT NULL,
    age INTEGER NOT NULL,
    gender VARCHAR(32) NOT NULL, 
    first_name VARCHAR(32) NOT NULL,
    last_name VARCHAR(32) NOT NULL,
    email VARCHAR(32) NOT NULL, 
    password_hash VARCHAR(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS session (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL, 
    token VARCHAR(64) NOT NULL,
    expiration_time TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE 
);

CREATE TABLE IF NOT EXISTS post (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    create_at TIMESTAMP NOT NULL, 
    update_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS category (
    id INTEGER PRIMARY KEY AUTOINCREMENT, 
    title TEXT
);

INSERT INTO category(title)
SELECT 'Golang'
WHERE NOT EXISTS (SELECT * FROM category WHERE title = 'Golang');

INSERT INTO category(title)
SELECT 'Python'
WHERE NOT EXISTS (SELECT * FROM category WHERE title = 'Python');

INSERT INTO category(title)
SELECT 'Java'
WHERE NOT EXISTS (SELECT * FROM category WHERE title = 'Java');

INSERT INTO category(title)
SELECT 'Js'
WHERE NOT EXISTS (SELECT * FROM category WHERE title = 'Js');

INSERT INTO category(title)
SELECT 'Php'
WHERE NOT EXISTS (SELECT * FROM category WHERE title = 'Php');

CREATE TABLE IF NOT EXISTS post_category (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER NOT NULL, 
    category_id INTEGER NOT NULL, 
    FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS comment (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER NOT NULL, 
    user_id INTEGER NOT NULL,
    text TEXT NOT NULL,
    create_at TIMESTAMP NOT NULL, 
    update_at TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    from_user_id INTEGER NOT NULL,
    to_user_id INTEGER NOT NULL, 
    message TEXT,
    create_at TIMESTAMP NOT NULL,
    FOREIGN KEY (from_user_id) REFERENCES user (id) ON DELETE CASCADE,
    FOREIGN KEY (to_user_id) REFERENCES user (id) ON DELETE CASCADE
)