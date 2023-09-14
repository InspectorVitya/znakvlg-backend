package hasher

import "golang.org/x/crypto/bcrypt"

//type PasswordHasher interface {
//	HashPassword(password string) (string, error)
//	CheckPasswordHash(password, hash string) bool
//}
//
//type BcryptHasher struct{}
//
//func New() *BcryptHasher {
//	return &BcryptHasher{}
//}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
