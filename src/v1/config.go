package v1

// ---
// ---
// ---

import (
	"os"
	"log"
)

// ---
// ---
// ---

var MongoServers string
var MongoDatabase string
var VerificationKey string

// ---
// ---
// ---

func init() {
	MongoServers = os.Getenv("MONGO_SERVERS")
	MongoDatabase = os.Getenv("MONGO_DATABASE")
	VerificationKey = os.Getenv("VERIFICATION_KEY")
	
	// ---
	
	if VerificationKey == "" {
		log.Println("generating verification key")
		
		vkb, vke := randomBytes(64)
		
		if vke != nil {
			log.Fatal(vke)
		}
		
		VerificationKey = hash512(vkb, nil)
	}
}

// ---
