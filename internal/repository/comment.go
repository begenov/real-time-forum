package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/begenov/real-time-forum/internal/domain"
)

type CommentRepo struct {
	db *sql.DB
}

func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{
		db: db,
	}
}

func (r *CommentRepo) Create(ctx context.Context, comment domain.Comment) error {
	stmt := `INSERT INTO comment (post_id, user_id, text, create_at, update_at) VALUES($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, stmt, comment.PostID, comment.UserID, comment.Text, comment.CreateAt, comment.UpdateAt)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *CommentRepo) Update(ctx context.Context, comment domain.Comment) error {
	stmt := `UPDATE comment SET text=$1, update_at=$2 WHERE id = $3`
	res, err := r.db.ExecContext(ctx, stmt, comment.Text, comment.UpdateAt, comment.Id)
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

func (r *CommentRepo) GetAllComment(ctx context.Context) ([]domain.Comment, error) {
	stmt := `SELECT * FROM comment`
	row, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	var comments []domain.Comment
	for row.Next() {
		var comment domain.Comment
		if err = row.Scan(&comment.Id, &comment.PostID, &comment.UserID, &comment.Text, &comment.CreateAt, &comment.UpdateAt); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (r *CommentRepo) GetCommentByID(ctx context.Context, id int) (domain.Comment, error) {
	var comment domain.Comment
	stmt := `SELECT * FROM comment WHERE id = $1`
	err := r.db.QueryRowContext(ctx, stmt, id).Scan(&comment.Id, &comment.PostID, &comment.UserID, &comment.Text, &comment.CreateAt, &comment.UpdateAt)
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *CommentRepo) Delete(ctx context.Context, id int) error {
	stmt := `DELETE FROM comment WHERE id = $1`
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
