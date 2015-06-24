package v1

// ---
// ---
// ---

import (
	"os"
)

// ---
// ---
// ---

var MongoServers string
var MongoDatabase string
var ValidationKey string

// ---
// ---
// ---

func init() {
	MongoServers = os.Getenv("MONGO_SERVERS")
	MongoDatabase = os.Getenv("MONGO_DATABASE")
	ValidationKey = os.Getenv("VALIDATION_KEY")
}

// ---
