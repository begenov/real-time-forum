package repository

import (
	"context"
	"database/sql"

	"github.com/begenov/real-time-forum/internal/domain"
)

type SessionRepo struct {
	db *sql.DB
}

func NewSession(db *sql.DB) *SessionRepo {
	return &SessionRepo{
		db: db,
	}
}

func (r *SessionRepo) Create(ctx context.Context, session domain.Session) error {

	stmt := `INSERT INTO session(user_id, token, expiration_time) VALUES($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, stmt, session.UserID, session.Token, session.ExpiresAt)

	if err != nil {
		return err
	}
	return nil
}

func (r *SessionRepo) GetSessionByUserID(ctx context.Context, userID int) (domain.Session, error) {
	var session domain.Session
	stmt := `SELECT user_id, token, expiration_time FROM session WHERE user_id = $1`

	if err := r.db.QueryRowContext(ctx, stmt).Scan(&session.UserID, &session.Token, &session.ExpiresAt); err != nil {
		return session, err
	}

	return session, nil
}

func (r *SessionRepo) Update(ctx context.Context, session domain.Session) error {
	stmt := `UPDATE session SET token = $1, expiration_time = $2 WHERE user_id = $3`
	_, err := r.db.ExecContext(ctx, stmt, session.Token, session.ExpiresAt, session.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *SessionRepo) Delete(ctx context.Context, id int) error {
	stmt := `DELETE * FROM session WHERE user_id = $1`
	_, err := r.db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}
