package v1

// ---
// ---
// ---

import (
	"time"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	
	// ---
	
	"github.com/dgrijalva/jwt-go"
)

// ---
// ---
// ---

func randomBytes(size int) ([]byte, error) {
	b := make([]byte, size)
	
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
	b, e := randomBytes(64)
	
	if e != nil {
		return PasswordSalt(""), e
	}
	
	// ---
	
	return PasswordSalt(hash512(b, nil)), nil
}

func passwordHash(password Password, passwordSalt PasswordSalt) (PasswordHash) {
	return PasswordHash(hash512([]byte(password), []byte(passwordSalt)))
}

// ---
// ---
// ---

func makeSimpleJwt(key string, val string, exp time.Time) (string, error) {
	tokenObj := jwt.New(jwt.SigningMethodHS256)
	
	tokenObj.Claims["val"] = val
	tokenObj.Claims["exp"] = exp.Unix()
	
	return tokenObj.SignedString([]byte(key))
}

func makeTokenJwt(key string, val string, tkn string, exp time.Time) (string, error) {
	tokenObj := jwt.New(jwt.SigningMethodHS256)
	
	tokenObj.Claims["val"] = val
	tokenObj.Claims["tkn"] = tkn
	tokenObj.Claims["exp"] = exp.Unix()
	
	return tokenObj.SignedString([]byte(key))
}

func parseJwt(key string, token string) (map[string]interface{}, error) {
	tokenObj, tokenObjErr := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	
	if tokenObjErr != nil {
		return nil, tokenObjErr
	}
	
	if !tokenObj.Valid {
		return nil, ErrInvalidToken
	}
	
	return tokenObj.Claims, nil
}

// ---
