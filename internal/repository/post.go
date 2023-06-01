package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/begenov/real-time-forum/internal/domain"
)

type PostRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (r *PostRepo) Create(ctx context.Context, post domain.Post) error {
	stmt := `INSERT INTO post (user_id, title, description, category, create_at, update_at) VALUES ($1, $2, $3, $4, $5, $6)`
	category := strings.Join(post.Category, " ")
	_, err := r.db.ExecContext(ctx, stmt, post.UserID, post.Title, post.Description, category, post.CreateAt, post.UpdateAt)

	if err != nil {
		fmt.Println(category, err)
		return err
	}

	return nil
}

func (r *PostRepo) Update(ctx context.Context, post domain.Post) error {
	category := strings.Join(post.Category, " ")
	stmt := `UPDATE post SET title = $1, description = $2, category = $3, update_at = $4 WHERE id = $5`
	res, err := r.db.ExecContext(ctx, stmt, post.Title, post.Description, category, post.UpdateAt, post.ID)
	if err != nil {
		return err
	}
	i, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if i == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *PostRepo) GetPostByID(ctx context.Context, id int) (domain.Post, error) {
	var post domain.Post
	stmt := `SELECT * FROM post WHERE id = $1`
	category := ""
	if err := r.db.QueryRowContext(ctx, stmt, id).Scan(&post.ID, &post.UserID, &post.Title, &post.Description, &category, &post.CreateAt, &post.UpdateAt); err != nil {
		return post, err
	}

	post.Category = strings.Split(category, " ")

	return post, nil
}

func (r *PostRepo) GetAllPosts(ctx context.Context) ([]domain.Post, error) {
	var posts []domain.Post
	stmt := `SELECT * FROM post`
	row, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		return posts, err
	}
	for row.Next() {
		fmt.Println("1")
		var post domain.Post
		var category string
		err := row.Scan(&post.ID, &post.UserID, &post.Title, &post.Description, &category, &post.CreateAt, &post.UpdateAt)
		if err != nil {
			return posts, err
		}
		post.Category = strings.Split(category, " ")
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepo) DeletePost(ctx context.Context, id int) error {
	stmt := `DELETE FROM post WHERE id = $1`
	res, err := r.db.ExecContext(ctx, stmt, id)

	if err != nil {
		return err
	}

	i, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if i == 0 {
		return sql.ErrNoRows
	}
	return nil
}
