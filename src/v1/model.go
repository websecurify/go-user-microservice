package v1

// ---
// ---
// ---

import (
	"log"
	
	// ---
	
	"gopkg.in/mgo.v2"
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
	ObjectId bson.ObjectId `bson:"_id,omitempty"`
	Id Id `bson:"id"`
	Name Name `bson:"name"`
	Email Email `bson:"email"`
	PasswordSalt PasswordSalt `bson:"passwordSalt"`
	PasswordHash PasswordHash `bson:"passwordHash"`
}

// ---
// ---
// ---

func initModel() {
	idIndex := mgo.Index{
		Key: []string{"id"},
		Unique: true,
		DropDups: false,
		Background: true,
		Sparse: false,
	}
	
	idEnsureErr := MongoCollection.EnsureIndex(idIndex)
	
	if idEnsureErr != nil {
		log.Fatal(idEnsureErr)
	}
	
	// ---
	
	emailIndex := mgo.Index{
		Key: []string{"email"},
	}
	
	emailEnsureErr := MongoCollection.EnsureIndex(emailIndex)
	
	if emailEnsureErr != nil {
		log.Fatal(emailEnsureErr)
	}
}

// ---
