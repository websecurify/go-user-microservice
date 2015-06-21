package v1

// ---
// ---
// ---

import (
	"gopkg.in/mgo.v2/bson"
)

// ---
// ---
// ---

type Id string
type Name string
type Email string
type Password string
type PasswordSalt string
type PasswordHash string

// ---
// ---
// ---

type UserEntry struct {
	ObjectId bson.ObjectId `bson:"_id"`
	Id Id `bson:"id"`
	Name Name `bson:"name"`
	Email Email `bson:"email"`
	PasswordSalt PasswordSalt `bson:"passwordSalt"`
	PasswordHash PasswordHash `bson:"passwordHash"`
}

// ---
