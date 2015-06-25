package v1

// ---
// ---
// ---

import (
	"net/http"
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
	var err error
	
	reply.Id, err = Create(args.Name, args.Email, args.Verified, args.Password)
	
	return err
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
	var err error
	
	err = Destroy(args.Id)
	
	return err
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
	var err error
	
	reply.Name, reply.Email, reply.Verified, err = Query(args.Id)
	
	return err
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
	var err error
	
	reply.Id, reply.Name, reply.Verified, err = QueryByEmail(args.Email)
	
	return err
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
	var err error
	
	reply.Name, reply.Email, reply.Verified, err = Login(args.Id, args.Password)
	
	return err
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
	var err error
	
	reply.Id, reply.Name, reply.Verified, err = LoginByEmail(args.Email, args.Password)
	
	return err
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
	var err error
	
	err = UpdateName(args.Id, args.Name)
	
	return err
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
	var err error
	
	err = UpdatePassword(args.Id, args.Password)
	
	return err
}

// ---
// ---
// ---

type StartVerifyArgs struct {
	Id Id `json:"id"`
}

type StartVerifyReply struct {
	Token Token `json:"token"`
}

func (s *UserMicroservice) StartVerify(r *http.Request, args *StartVerifyArgs, reply *StartVerifyReply) (error) {
	var err error
	
	reply.Token, err = StartVerify(args.Id)
	
	return err
}

// ---
// ---
// ---

type VerifyArgs struct {
	Token Token `json:"token"`
}

type VerifyReply struct {
}

func (s *UserMicroservice) Verify(r *http.Request, args *VerifyArgs, reply *VerifyReply) (error) {
	var err error
	
	err = Verify(args.Token)
	
	return err
}

// ---
// ---
// ---

type StartResetArgs struct {
	Id Id `json:"id"`
}

type StartResetReply struct {
	Token Token `json:"token"`
}

func (s *UserMicroservice) StartReset(r *http.Request, args *StartResetArgs, reply *StartResetReply) (error) {
	var err error
	
	reply.Token, err = StartReset(args.Id)
	
	return err
}

// ---
// ---
// ---

type ResetArgs struct {
	Token Token `json:"token"`
	Password Password `json:"password"`
}

type ResetReply struct {
}

func (s *UserMicroservice) Reset(r *http.Request, args *ResetArgs, reply *ResetReply) (error) {
	var err error
	
	err = Reset(args.Token, args.Password)
	
	return err
}

// ---
// ---
// ---

func Start() {
	InitMongo()
}

// ---
