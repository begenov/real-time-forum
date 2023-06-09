package repository

import (
	"context"
	"database/sql"

	"github.com/begenov/real-time-forum/internal/domain"
)

type AuthorizationRepo struct {
	db *sql.DB
}

func NewAuthorization(db *sql.DB) *AuthorizationRepo {
	return &AuthorizationRepo{
		db: db,
	}
}

func (r *AuthorizationRepo) Create(ctx context.Context, user domain.User) error {
	stmt := `INSERT INTO user(nick_name, age, gender, first_name, last_name, email, password_hash) VALUES($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, stmt, user.Nickname, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (r *AuthorizationRepo) GetByID(ctx context.Context, id int) (domain.User, error) {
	var user domain.User
	stmt := `SELECT id, nick_name, age, gender, first_name, last_name, email, password_hash FROM user WHERE id = $1`
	if err := r.db.QueryRowContext(ctx, stmt, id).Scan(&user.Id, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password); err != nil {
		return user, err
	}
	return user, nil
}

func (r *AuthorizationRepo) GetByNickname(ctx context.Context, nickname string) (domain.User, error) {
	var user domain.User
	stmt := `SELECT id, nick_name, age, gender, first_name, last_name, email, password_hash FROM user WHERE nick_name=$1`
	if err := r.db.QueryRowContext(ctx, stmt, nickname).Scan(&user.Id, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password); err != nil {
		return user, err
	}
	return user, nil
}

func (r *AuthorizationRepo) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	stmt := `SELECT id, nick_name, age, gender, first_name, last_name, email, password_hash FROM user WHERE email=$1`
	if err := r.db.QueryRowContext(ctx, stmt, email).Scan(&user.Id, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password); err != nil {
		return user, err
	}
	return user, nil
}

func (r *AuthorizationRepo) UpdatePassword(ctx context.Context, password string, id int) error {
	stmt := `UPDATE user SET password_hash = $1 WHERE id = $2`
	res, err := r.db.ExecContext(ctx, stmt, password, id)

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

func (r *AuthorizationRepo) AllUsers(ctx context.Context, userID int) ([]domain.Users, error) {
	var users []domain.Users
	stmt := `SELECT id, nick_name FROM user WHERE id <> ? `
	row, err := r.db.QueryContext(ctx, stmt, userID)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var user domain.Users
		if err := row.Scan(&user.ID, &user.Nickname); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
