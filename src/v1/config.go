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

// ---
// ---
// ---

func init() {
	MongoServers = os.Getenv("MONGO_SERVERS")
	MongoDatabase = os.Getenv("MONGO_DATABASE")
}

// ---
