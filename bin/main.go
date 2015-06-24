package main // import "github.com/websecurify/go-user-microservice"

// ---
// ---
// ---

import (
	"os"
	"log"
	"net/http"
	
	// ---
	
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/gorilla/handlers"
	
	// ---
	
	"v1"
)

// ---
// ---
// ---

func main() {
	v1.Start()
	
	// ---
	
	s := rpc.NewServer()
	
	// ---
	
	s.RegisterService(new(v1.UserMicroservice), "v1")
	
	// ---
	
	s.RegisterCodec(json.NewCodec(), "application/json")
	
	// ---
	
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, s))
	
	// ---
	
	p := os.Getenv("PORT")
	
	if p == "" {
		p = "8080"
	}
	
	// ---
	
	log.Println("listening on", p)
	
	// ---
	
	log.Fatal(http.ListenAndServe(":" + p, nil))
}

// ---
