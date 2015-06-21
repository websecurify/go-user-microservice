package v1

// ---
// ---
// ---

import (
	"errors"
	"net/http"
	
	// ---
	
	"gopkg.in/mgo.v2/bson"
	
	// ---
	
	"code.google.com/p/go-uuid/uuid"
)

// ---
// ---
// ---

type UserService struct {
}

// ---
// ---
// ---

type CreateArgs struct {
	Name Name `json:"name"`
	Email Email `json:"email"`
	Password Password `json:"password"`
}

type CreateReply struct {
	Id Id `json:"id"`
}

// ---

func (s *UserService) Create(r *http.Request, args *CreateArgs, reply *CreateReply) (error) {
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

func (s *UserService) Destroy(r *http.Request, args *DestroyArgs, reply *DestroyReply) (error) {
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
}

// ---

func (s *UserService) Query(r *http.Request, args *QueryArgs, reply *QueryReply) (error) {
	result := UserEntry{}
	
	// ---
	
	findErr := MongoCollection.Find(bson.M{"id": args.Id}).One(&result)
	
	if findErr != nil {
		return findErr
	}
	
	// ---
	
	reply.Name = result.Name
	reply.Email = result.Email
	
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
}

// ---

func (s *UserService) QueryByEmail(r *http.Request, args *QueryByEmailArgs, reply *QueryByEmailReply) (error) {
	result := UserEntry{}
	
	// ---
	
	findErr := MongoCollection.Find(bson.M{"email": args.Email}).One(&result)
	
	if findErr != nil {
		return findErr
	}
	
	// ---
	
	reply.Id = result.Id
	reply.Name = result.Name
	
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
}

// ---

func (s *UserService) Login(r *http.Request, args *LoginArgs, reply *LoginReply) (error) {
	result := UserEntry{}
	
	// ---
	
	findErr := MongoCollection.Find(bson.M{"id": args.Id}).One(&result)
	
	if findErr != nil {
		return findErr
	}
	
	// ---
	
	if result.PasswordHash != passwordHash(args.Password, result.PasswordSalt) {
		return errors.New("password mismatch")
	}
	
	// ---
	
	reply.Name = result.Name
	reply.Email = result.Email
	
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
}

// ---

func (s *UserService) LoginByEmail(r *http.Request, args *LoginByEmailArgs, reply *LoginByEmailReply) (error) {
	result := UserEntry{}
	
	// ---
	
	findErr := MongoCollection.Find(bson.M{"email": args.Email}).One(&result)
	
	if findErr != nil {
		return findErr
	}
	
	// ---
	
	if result.PasswordHash != passwordHash(args.Password, result.PasswordSalt) {
		return errors.New("password mismatch")
	}
	
	// ---
	
	reply.Id = result.Id
	reply.Name = result.Name
	
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

func (s *UserService) UpdateName(r *http.Request, args *UpdateNameArgs, reply *UpdateNameReply) (error) {
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

func (s *UserService) UpdatePassword(r *http.Request, args *UpdatePasswordArgs, reply *UpdatePasswordReply) (error) {
	result := UserEntry{}
	
	// ---
	
	findErr := MongoCollection.Find(bson.M{"id": args.Id}).One(&result)
	
	if findErr != nil {
		return findErr
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

func Start() {
	InitMongo()
}

// ---
