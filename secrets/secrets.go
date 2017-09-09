package secrets

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a password as a string and returns a hashed string
func HashPassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePassword checks to see if the stored hash and the plain text are the same password
func ComparePassword(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}
