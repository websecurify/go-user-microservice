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
	ObjectId bson.ObjectId `bson:"_id"`
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
	index := mgo.Index{
		Key: []string{"id"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}
	
	// ---
	
	ensureErr := MongoCollection.EnsureIndex(index)
	
	if ensureErr != nil {
		log.Fatal(ensureErr)
	}
	
	// ---
	
	index2 := mgo.Index{
		Key: []string{"email"},
	}
	
	// ---
	
	ensureErr2 := MongoCollection.EnsureIndex(index2)
	
	if ensureErr2 != nil {
		log.Fatal(ensureErr2)
	}
}

// ---
