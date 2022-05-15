package postgresql

import (
	"GoNews/pkg/storage"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres struct {
	db *pgxpool.Pool
}

func New(dbURL string) (*Postgres, error) {
	db, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}
	return &Postgres{db: db}, nil
}

func (s *Postgres) AddPost(p storage.Post) error {
	_, err := s.db.Exec(context.Background(),
		"INSERT INTO posts (author_id, title, content) "+
			"VALUES ($1,$2,$3) "+
			"RETURNING posts.id;",
		p.AuthorID, p.Title, p.Content)
	if err != nil {
		return err
	}
	return nil
}

func (s *Postgres) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(),
		"SELECT * FROM posts;")
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	for rows.Next() {
		var p storage.Post
		err = rows.Scan(&p.ID, &p.AuthorID, &p.Title, &p.Content, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}

func (s *Postgres) DeletePost(p storage.Post) error {
	_, err := s.db.Exec(context.Background(),
		"DELETE FROM posts WHERE id=$1;", p.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Postgres) UpdatePost(p storage.Post) error {
	_, err := s.db.Exec(context.Background(),
		"UPDATE posts "+
			"SET author_id = $1, "+
			"title = $2, "+
			"content = $3 "+
			"WHERE id = $4", p.AuthorID, p.Title, p.Content, p.ID)
	if err != nil {
		return err
	}
	return nil
}
