package auth

import "github.com/gofrs/uuid"

type TokenManager interface {
	GenerateToken() (string, error)
}

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) GenerateToken() (string, error) {
	token, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return token.String(), nil
}
