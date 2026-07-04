package password_helper

import "golang.org/x/crypto/bcrypt"

// Hash returns a bcrypt hash for the given plain-text password.
func Hash(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(bytes), err
}

// Compare reports whether the plain-text password matches the stored hash.
func Compare(hash, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}
