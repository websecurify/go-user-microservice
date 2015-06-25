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

var ResetKey string
var VerifyKey string
var MongoServers string
var MongoDatabase string

// ---
// ---
// ---

func init() {
	ResetKey = os.Getenv("RESET_KEY")
	VerifyKey = os.Getenv("VERIFY_KEY")
	MongoServers = os.Getenv("MONGO_SERVERS")
	MongoDatabase = os.Getenv("MONGO_DATABASE")
	
	// ---
	
	if ResetKey == "" {
		log.Println("generating verify key")
		
		rkb, rke := randomBytes(64)
		
		if rke != nil {
			log.Fatal(rke)
		}
		
		ResetKey = hash512(rkb, nil)
	}
	
	// ---
	
	if VerifyKey == "" {
		log.Println("generating reset key")
		
		vkb, vke := randomBytes(64)
		
		if vke != nil {
			log.Fatal(vke)
		}
		
		VerifyKey = hash512(vkb, nil)
	}
}

// ---
