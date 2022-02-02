package auth

import "golang.org/x/crypto/bcrypt"

type Hasher struct {
}

func NewHashManager() *Hasher {
	return &Hasher{}
}

func (h *Hasher) Hash(str string) (string, error) {
	hashedStr, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedStr), nil
}
func (h *Hasher) Validate(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
