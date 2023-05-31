package hash

import "golang.org/x/crypto/bcrypt"

type PasswordHasher interface {
	GenerateFromPassword(password string) (string, error)
	CompareHashAndPassword(hashPassword string, password string) error
}

type HashPassword struct {
	cost int
}

func NewHash(cost int) *HashPassword {
	return &HashPassword{
		cost: cost,
	}
}

func (h *HashPassword) GenerateFromPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (h *HashPassword) CompareHashAndPassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
