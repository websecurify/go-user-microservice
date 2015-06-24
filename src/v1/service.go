package v1

// ---
// ---
// ---

import (
	"time"
	"errors"
	"net/http"
	
	// ---
	
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	
	// ---
	
	"github.com/dgrijalva/jwt-go"
	
	// ---
	
	"code.google.com/p/go-uuid/uuid"
)

// ---
// ---
// ---

var (
	ErrNotFound = errors.New("not found")
	ErrInvalidToken = errors.New("invalid token")
	ErrPasswordMismatch = errors.New("password mismatch")
)

// ---
// ---
// ---

type UserMicroservice struct {
}

// ---
// ---
// ---

type CreateArgs struct {
	Name Name `json:"name"`
	Email Email `json:"email"`
	Verified Verified `json:"verified"`
	Password Password `json:"password"`
}

type CreateReply struct {
	Id Id `json:"id"`
}

// ---

func (s *UserMicroservice) Create(r *http.Request, args *CreateArgs, reply *CreateReply) (error) {
	id := Id(uuid.NewRandom().String())
	
	// ---
	
	passwordSalt, passwordSaltErr := passwordSalt()
	
	if passwordSaltErr != nil {
		return passwordSaltErr
	}
	
	// ---
	
	passwordHash := passwordHash(args.Password, passwordSalt)
	
	// ---
	
	entry := UserEntry{
		ObjectId: bson.NewObjectId(),
		Id: id,
		Name: args.Name,
		Email: args.Email,
		Verified: args.Verified,
		PasswordSalt: passwordSalt,
		PasswordHash: passwordHash,
	}
	
	// ---
	
	insertErr := MongoCollection.Insert(entry)
	
	if insertErr != nil {
		return insertErr
	}
	
	// ---
	
	reply.Id = id
	
	// ---
	
	return nil
}

// ---
// ---
// ---

type DestroyArgs struct {
	Id Id `json:"id"`
}

type DestroyReply struct {
}

// ---

func (s *UserMicroservice) Destroy(r *http.Request, args *DestroyArgs, reply *DestroyReply) (error) {
	removeErr := MongoCollection.Remove(bson.M{"id": args.Id})
	
	if removeErr != nil {
		return removeErr
	}
	
	// ---
	
	return nil
}

// ---
// ---
// ---

type QueryArgs struct {
	Id Id `json:"id"`
}

type QueryReply struct {
	Name Name `json:"name"`
	Email Email `json:"email"`
	Verified Verified `json:"verified"`
}

// ---

func (s *UserMicroservice) Query(r *http.Request, args *QueryArgs, reply *QueryReply) (error) {
	result := UserEntry{}
	
	findErr := MongoCollection.Find(bson.M{"id": args.Id}).One(&result)
	
	if findErr != nil {
		if findErr == mgo.ErrNotFound {
			return ErrNotFound
		} else {
			return findErr
		}
	}
	
	// ---
	
	reply.Name = result.Name
	reply.Email = result.Email
	reply.Verified = result.Verified
	
	// ---
	
	return nil
}

// ---
// ---
// ---

type QueryByEmailArgs struct {
	Email Email `json:"email"`
}

type QueryByEmailReply struct {
	Id Id `json:"id"`
	Name Name `json:"name"`
	Verified Verified `json:"verified"`
}

// ---

func (s *UserMicroservice) QueryByEmail(r *http.Request, args *QueryByEmailArgs, reply *QueryByEmailReply) (error) {
	result := UserEntry{}
	
	findErr := MongoCollection.Find(bson.M{"email": args.Email}).One(&result)
	
	if findErr != nil {
		if findErr == mgo.ErrNotFound {
			return ErrNotFound
		} else {
			return findErr
		}
	}
	
	// ---
	
	reply.Id = result.Id
	reply.Name = result.Name
	reply.Verified = result.Verified
	
	// ---
	
	return nil
}

// ---
// ---
// ---

type LoginArgs struct {
	Id Id `json:"id"`
	Password Password `json:"password"`
}

type LoginReply struct {
	Name Name `json:"name"`
	Email Email `json:"email"`
	Verified Verified `json:"verified"`
}

// ---

func (s *UserMicroservice) Login(r *http.Request, args *LoginArgs, reply *LoginReply) (error) {
	result := UserEntry{}
	
	findErr := MongoCollection.Find(bson.M{"id": args.Id}).One(&result)
	
	if findErr != nil {
		if findErr == mgo.ErrNotFound {
			return ErrNotFound
		} else {
			return findErr
		}
	}
	
	// ---
	
	if result.PasswordHash != passwordHash(args.Password, result.PasswordSalt) {
		return ErrPasswordMismatch
	}
	
	// ---
	
	reply.Name = result.Name
	reply.Email = result.Email
	reply.Verified = result.Verified
	
	// ---
	
	return nil
}

// ---
// ---
// ---

type LoginByEmailArgs struct {
	Email Email `json:"email"`
	Password Password `json:"password"`
}

type LoginByEmailReply struct {
	Id Id `json:"id"`
	Name Name `json:"name"`
	Verified Verified `json:"verified"`
}

// ---

func (s *UserMicroservice) LoginByEmail(r *http.Request, args *LoginByEmailArgs, reply *LoginByEmailReply) (error) {
	result := UserEntry{}
	
	findErr := MongoCollection.Find(bson.M{"email": args.Email}).One(&result)
	
	if findErr != nil {
		if findErr == mgo.ErrNotFound {
			return ErrNotFound
		} else {
			return findErr
		}
	}
	
	// ---
	
	if result.PasswordHash != passwordHash(args.Password, result.PasswordSalt) {
		return ErrPasswordMismatch
	}
	
	// ---
	
	reply.Id = result.Id
	reply.Name = result.Name
	reply.Verified = result.Verified
	
	// ---
	
	return nil
}

// ---
// ---
// ---

type UpdateNameArgs struct {
	Id Id `json:"id"`
	Name Name `json:"name"`
}

type UpdateNameReply struct {
}

// ---

func (s *UserMicroservice) UpdateName(r *http.Request, args *UpdateNameArgs, reply *UpdateNameReply) (error) {
	updateErr := MongoCollection.Update(bson.M{"id": args.Id}, bson.M{"$set": bson.M{"name": args.Name}})
	
	if updateErr != nil {
		return updateErr
	}
	
	// ---
	
	return nil
}

// ---
// ---
// ---

type UpdatePasswordArgs struct {
	Id Id `json:"id"`
	Password Password `json:"password"`
}

type UpdatePasswordReply struct {
}

// ---

func (s *UserMicroservice) UpdatePassword(r *http.Request, args *UpdatePasswordArgs, reply *UpdatePasswordReply) (error) {
	result := UserEntry{}
	
	findErr := MongoCollection.Find(bson.M{"id": args.Id}).One(&result)
	
	if findErr != nil {
		if findErr == mgo.ErrNotFound {
			return ErrNotFound
		} else {
			return findErr
		}
	}
	
	// ---
	
	updateErr := MongoCollection.Update(bson.M{"id": args.Id}, bson.M{"$set": bson.M{"passwordHash": passwordHash(args.Password, result.PasswordSalt)}})
	
	if updateErr != nil {
		return updateErr
	}
	
	// ---
	
	return nil
}

// ---
// ---
// ---

type StartVerificationArgs struct {
	Id Id `json:"id"`
}

type StartVerificationReply struct {
	Token string `json:"token"`
}

func (s *UserMicroservice) StartVerification(r *http.Request, args *StartVerificationArgs, reply *StartVerificationReply) (error) {
	result := UserEntry{}
	
	findErr := MongoCollection.Find(bson.M{"id": args.Id}).One(&result)
	
	if findErr != nil {
		if findErr == mgo.ErrNotFound {
			return ErrNotFound
		} else {
			return findErr
		}
	}
	
	// ---
	
	token := jwt.New(jwt.SigningMethodHS256)
	
	token.Claims["uid"] = string(args.Id)
	token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	
	tokenString, tokenStringErr := token.SignedString([]byte(VerificationKey))
	
	if tokenStringErr != nil {
		return tokenStringErr
	}
	
	// ---
	
	reply.Token = tokenString
	
	// ---
	
	return nil
}

// ---
// ---
// ---

type VerifyArgs struct {
	Token string `json:"token"`
}

type VerifyReply struct {
}

func (s *UserMicroservice) Verify(r *http.Request, args *VerifyArgs, reply *VerifyReply) (error) {
	token, tokenErr := jwt.Parse(args.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(VerificationKey), nil
	})
	
	if tokenErr != nil {
		return tokenErr
	}
	
	if !token.Valid {
		return ErrInvalidToken
	}
	
	// ---
	
	updateErr := MongoCollection.Update(bson.M{"id": token.Claims["uid"]}, bson.M{"$set": bson.M{"verified": true}})
	
	if updateErr != nil {
		return updateErr
	}
	
	// ---
	
	return nil
}

// ---
// ---
// ---

func Start() {
	InitMongo()
}

// ---
