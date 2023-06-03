package repository

import (
	"context"
	"database/sql"

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
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO post (user_id, title, description, create_at, update_at) VALUES ($1, $2, $3, $4, $5)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, post.UserID, post.Title, post.Description, post.CreateAt, post.UpdateAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	st := "INSERT INTO post_category (post_id, category_id) VALUES ($1, $2)"

	for _, categoryID := range post.CategoryID {
		_, err := tx.ExecContext(ctx, st, id, categoryID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *PostRepo) Update(ctx context.Context, post domain.Post) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	stmt := `UPDATE post SET title = $1, description = $2, update_at = $3 WHERE id = $4`
	_, err = tx.ExecContext(ctx, stmt, post.Title, post.Description, post.UpdateAt, post.ID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if len(post.CategoryID) != 0 {
		stmt = `DELETE FROM post_category WHERE post_id = $1`
		_, err = tx.ExecContext(ctx, stmt, post.ID)
		if err != nil {
			_ = tx.Rollback()
			return err
		}

		stmt = `INSERT INTO post_category (post_id, category_id) VALUES ($1, $2)`
		for _, category := range post.CategoryID {
			_, err = tx.ExecContext(ctx, stmt, post.ID, category)
			if err != nil {
				_ = tx.Rollback()
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepo) GetPostByID(ctx context.Context, id int) (domain.Post, error) {
	var post domain.Post
	stmt := `SELECT p.id, p.user_id, p.title, p.description, p.create_at, p.update_at, c.title, c.id, u.nick_name
			FROM post AS p
			LEFT JOIN post_category AS pc ON p.id = pc.post_id
			LEFT JOIN category AS c ON pc.category_id = c.id
			LEFT JOIN user AS u ON p.user_id = u.id
			WHERE p.id = $1`

	rows, err := r.db.QueryContext(ctx, stmt, id)
	if err != nil {
		return post, err
	}
	defer rows.Close()

	for rows.Next() {
		var categoryTitle sql.NullString
		var categoryID sql.NullInt64
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Description, &post.CreateAt, &post.UpdateAt, &categoryTitle, &categoryID, &post.Author)
		if err != nil {
			return post, err
		}

		if categoryTitle.Valid {
			post.Category = append(post.Category, categoryTitle.String)
			post.CategoryID = append(post.CategoryID, int(categoryID.Int64))
		}
	}

	if err = rows.Err(); err != nil {
		return post, err
	}

	return post, nil
}

func (r *PostRepo) GetAllPosts(ctx context.Context) ([]domain.Post, error) {
	stmt := `SELECT p.id, p.user_id, p.title, p.description, p.create_at, p.update_at, c.title, c.id, u.nick_name
				FROM post AS p
				LEFT JOIN post_category AS pc ON p.id = pc.post_id
				LEFT JOIN category AS c ON pc.category_id = c.id
				LEFT JOIN user AS u ON p.user_id = u.id
			`

	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []domain.Post
	postIDs := make(map[int]bool)

	for rows.Next() {
		var post domain.Post
		var categoryTitle sql.NullString
		var categoryID sql.NullInt64

		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Description, &post.CreateAt, &post.UpdateAt, &categoryTitle, &categoryID, &post.Author)
		if err != nil {
			return nil, err
		}

		if _, ok := postIDs[post.ID]; !ok {

			posts = append(posts, post)
			postIDs[post.ID] = true
		}

		if categoryTitle.Valid {
			// Category exists, add it to the post's Category and CategoryID fields
			posts[len(posts)-1].Category = append(posts[len(posts)-1].Category, categoryTitle.String)
			posts[len(posts)-1].CategoryID = append(posts[len(posts)-1].CategoryID, int(categoryID.Int64))
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepo) DeletePost(ctx context.Context, id int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	stmt := `DELETE FROM post_category WHERE post_id = $1`
	_, err = tx.ExecContext(ctx, stmt, id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	stmt = `DELETE FROM post WHERE id = $1`
	res, err := tx.ExecContext(ctx, stmt, id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if rowsAffected == 0 {
		_ = tx.Rollback()
		return sql.ErrNoRows
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
