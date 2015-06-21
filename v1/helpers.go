package v1

// ---
// ---
// ---

import (
	"crypto/rand"
	"crypto/sha1"
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

func passwordSalt() (PasswordSalt, error) {
	b, e := randomBytes()
	
	if e != nil {
		return PasswordSalt(""), e
	}
	
	// ---
	
	h := sha1.New()
	
	// ---
	
	h.Write(b)
	
	// ---
	
	return PasswordSalt(base64.URLEncoding.EncodeToString(h.Sum(nil))), nil
}

// ---
// ---
// ---

func passwordHash(password Password, passwordSalt PasswordSalt) (PasswordHash) {
	h := sha1.New()
	
	// ---
	
	h.Write([]byte(password))
	h.Write([]byte(passwordSalt))
	
	// ---
	
	return PasswordHash(base64.URLEncoding.EncodeToString(h.Sum(nil)))
}

// ---
