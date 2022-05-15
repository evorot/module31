DROP TABLE IF EXISTS posts, authors;

CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    author_id INTEGER REFERENCES authors(id) NOT NULL,
    title TEXT  NOT NULL,
    content TEXT NOT NULL,
    created_at BIGINT NOT NULL DEFAULT extract (epoch from now())
);



--Создание тестовых авторов
CREATE OR REPLACE PROCEDURE populate()
-- язык, на котором написана процедура
    LANGUAGE plpgsql
AS $$
-- начало транзакции
BEGIN
    FOR i IN 1..6 LOOP
            INSERT INTO authors(name) VALUES ('Author ' || i);
        END LOOP;
-- завершение процедуры и транзакции
END
$$;
CALL populate();

--Создание тестовых постов
CREATE OR REPLACE PROCEDURE create_posts()
-- язык, на котором написана процедура
    LANGUAGE plpgsql
AS $$
-- начало транзакции
BEGIN
    FOR i IN 1..4 LOOP
            INSERT INTO posts(author_id, title, content) VALUES (1,'Заголовок '||i,'Создание текста поста №'||i);
        END LOOP;
    FOR i IN 5..7 LOOP
            INSERT INTO posts(author_id, title, content) VALUES (2,'Заголовок '||i,'Создание текста поста №'||i);
        END LOOP;
    FOR i IN 8..9 LOOP
            INSERT INTO posts(author_id, title, content) VALUES (3,'Заголовок '||i,'Создание текста поста №'||i);
        END LOOP;
    FOR i IN 9..10 LOOP
            INSERT INTO posts(author_id, title, content) VALUES (4,'Заголовок '||i,'Создание текста поста №'||i);
        END LOOP;
    FOR i IN 11..12 LOOP
            INSERT INTO posts(author_id, title, content) VALUES (5,'Заголовок '||i,'Создание текста поста №'||i);
        END LOOP;
-- завершение процедуры и транзакции
END
$$;

CALL create_posts();
