package v1

// ---
// ---
// ---

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

// ---
// ---
// ---

func randomBytes() ([]byte, error) {
	b := make([]byte, 20)
	
	// ---
	
	_, e := rand.Read(b)
	
	if e != nil {
		return nil, e
	}
	
	// ---
	
	return b, nil
}

// ---
// ---
// ---

func hash512(value []byte, salt []byte) (string) {
	h := sha512.New()
	
	// ---
	
	h.Write(value)
	h.Write(salt)
	
	// ---
	
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

// ---
// ---
// ---

func passwordSalt() (PasswordSalt, error) {
	b, e := randomBytes()
	
	if e != nil {
		return PasswordSalt(""), e
	}
	
	// ---
	
	return PasswordSalt(hash512(b, nil)), nil
}

// ---
// ---
// ---

func passwordHash(password Password, passwordSalt PasswordSalt) (PasswordHash) {
	return PasswordHash(hash512([]byte(password), []byte(passwordSalt)))
}

// ---
