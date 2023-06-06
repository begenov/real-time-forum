package service

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/begenov/real-time-forum/internal/config"
	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/begenov/real-time-forum/internal/repository"
	"github.com/begenov/real-time-forum/pkg/auth"
	"github.com/begenov/real-time-forum/pkg/hash"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	auth    repository.Authorization
	session repository.Session
	hash    hash.PasswordHasher
	cfg     config.Token
	manager auth.TokenManager
}

func NewUserService(auth repository.Authorization, session repository.Session, hash hash.PasswordHasher, manager auth.TokenManager, cfg config.Token) *UserService {
	return &UserService{
		auth:    auth,
		session: session,
		hash:    hash,
		manager: manager,
		cfg:     cfg,
	}
}

func (s *UserService) SignUp(ctx context.Context, user domain.User) error {
	var err error
	user.Password, err = s.hash.GenerateFromPassword(user.Password)
	if err != nil {
		return err
	}
	if err = s.auth.Create(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) SignIn(ctx context.Context, email string, password string) (domain.Session, error) {
	var user domain.User
	var err error
	if strings.ContainsAny(email, "@") {
		user, err = s.auth.GetByEmail(ctx, email)

		if err != nil {
			return domain.Session{}, err
		}
	} else {
		user, err = s.auth.GetByNickname(ctx, email)
		if err != nil {
			return domain.Session{}, err
		}
	}

	if err = s.hash.CompareHashAndPassword(user.Password, password); err != nil {
		return domain.Session{}, err
	}

	token, err := s.manager.GenerateToken()
	if err != nil {
		return domain.Session{}, err
	}
	session, err := s.session.GetSessionByUserID(ctx, user.Id)
	if err == sql.ErrNoRows {

		session = domain.Session{
			UserID:    user.Id,
			Token:     token,
			ExpiresAt: time.Now().Add(time.Duration(s.cfg.Ttl) * time.Hour),
		}

		if err := s.session.Create(ctx, session); err != nil {
			return session, err
		}
		return session, nil
	}

	if err != nil {
		return session, err
	}

	session = domain.Session{
		UserID:    user.Id,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Duration(s.cfg.Ttl) * time.Hour),
	}

	if err := s.session.Update(ctx, session); err != nil {
		return session, err
	}

	fmt.Println(session)

	return session, nil
}

func (s *UserService) UpdatePassword(ctx context.Context, password string, id int) error {
	user, err := s.auth.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = s.hash.CompareHashAndPassword(user.Password, password)
	if err == bcrypt.ErrMismatchedHashAndPassword {
		password, err = s.hash.GenerateFromPassword(password)
		if err != nil {
			return err
		}
		return s.UpdatePassword(ctx, password, id)
	}

	return err
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	return s.auth.GetByID(ctx, id)
}

func (s *UserService) GetUserByToken(ctx context.Context, value string) (domain.User, error) {
	session, err := s.session.GetUserIDByToken(ctx, value)
	if err != nil {
		return domain.User{}, err
	}

	user, err := s.auth.GetByID(ctx, session.UserID)
	if err != nil {
		return user, err
	}
	user.ExpiresAt = session.ExpiresAt
	return user, nil
}

func (s *UserService) DeleteSession(ctx context.Context, value string) error {
	return s.session.Delete(ctx, value)
}

func (s *UserService) AllUsers(ctx context.Context) ([]domain.Users, error) {
	return s.auth.AllUsers(ctx)
}
