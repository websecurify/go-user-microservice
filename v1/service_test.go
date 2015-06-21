package v1

// ---
// ---
// ---

import (
	"testing"
	
	// ---
	
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	
	// ---
	
	"github.com/websecurify/go-dockertest"
)

// ---
// ---
// ---

func findById(id Id) (UserEntry, error) {
	u := UserEntry{
	}
	
	e := MongoCollection.Find(bson.M{"id": id}).One(&u)
	
	return u, e
}

// ---
// ---
// ---

func create(name Name, email Email, password Password) (CreateReply, error) {
	s := UserService{}
	
	a := CreateArgs{
		Name: name,
		Email: email,
		Password: password,
	}
	
	r := CreateReply{
	}
	
	e := s.Create(nil, &a, &r)
	
	return r, e
}

func destroy(id Id) (DestroyReply, error) {
	s := UserService{}
	
	a := DestroyArgs{
		Id: id,
	}
	
	r := DestroyReply{
	}
	
	e := s.Destroy(nil, &a, &r)
	
	return r, e
}

func query(id Id) (QueryReply, error) {
	s := UserService{}
	
	a := QueryArgs{
		Id: id,
	}
	
	r := QueryReply{
	}
	
	e := s.Query(nil, &a, &r)
	
	return r, e
}

func queryByEmail(email Email) (QueryByEmailReply, error) {
	s := UserService{}
	
	a := QueryByEmailArgs{
		Email: email,
	}
	
	r := QueryByEmailReply{
	}
	
	e := s.QueryByEmail(nil, &a, &r)
	
	return r, e
}

func login(id Id, password Password) (LoginReply, error) {
	s := UserService{}
	
	a := LoginArgs{
		Id: id,
		Password: password,
	}
	
	r := LoginReply{
	}
	
	e := s.Login(nil, &a, &r)
	
	return r, e
}

func loginByEmail(email Email, password Password) (LoginByEmailReply, error) {
	s := UserService{}
	
	a := LoginByEmailArgs{
		Email: email,
		Password: password,
	}
	
	r := LoginByEmailReply{
	}
	
	e := s.LoginByEmail(nil, &a, &r)
	
	return r, e
}

func updateName(id Id, name Name) (UpdateNameReply, error) {
	s := UserService{}
	
	a := UpdateNameArgs{
		Id: id,
		Name: name,
	}
	
	r := UpdateNameReply{
	}
	
	e := s.UpdateName(nil, &a, &r)
	
	return r, e
}

func updatePassword(id Id, password Password) (UpdatePasswordReply, error) {
	s := UserService{}
	
	a := UpdatePasswordArgs{
		Id: id,
		Password: password,
	}
	
	r := UpdatePasswordReply{
	}
	
	e := s.UpdatePassword(nil, &a, &r)
	
	return r, e
}

// ---
// ---
// ---

func TestEndToEnd(t *testing.T) {
	cid, cip := dockertest.SetupMongoContainer(t)
	
	defer cid.KillRemove(t)
	
	// ---
	
	MongoServers = cip
	MongoDatabase = "testing"
	
	InitMongo()
	
	// ---
	
	name := Name("Test")
	email := Email("test@test")
	password := Password("TestTest")
	
	// ---
	
	cr, ce := create(name, email, password)
	
	if ce != nil {
		t.Error(ce)
	}
	
	if cr.Id == "" {
		t.Error("no reply id")
	}
	
	// ---
	
	fr, fe := findById(cr.Id)
	
	if fe != nil {
		t.Error(fe)
	}
	
	if fr.Name != name {
		t.Error("name mismatch")
	}
	
	if fr.Email != email {
		t.Error("email mismatch")
	}
	
	if fr.PasswordSalt == "" {
		t.Error("password salt mismatch")
	}
	
	if fr.PasswordHash != passwordHash(password, fr.PasswordSalt) {
		t.Error("password hash mismatch")
	}
	
	// ---
	
	qr, qe := query(cr.Id)
	
	if qe != nil {
		t.Error(qe)
	}
	
	if qr.Name != name {
		t.Error("name mismatch")
	}
	
	if qr.Email != email {
		t.Error("email mismatch")
	}
	
	// ---
	
	qber, qbee := queryByEmail(email)
	
	if qbee != nil {
		t.Error(qbee)
	}
	
	if qber.Id != cr.Id {
		t.Error("id mismatch")
	}
	
	if qber.Name != name {
		t.Error("name mismatch")
	}
	
	// ---
	
	lr, le := login(cr.Id, password)
	
	if le != nil {
		t.Error(le)
	}
	
	if lr.Name != name {
		t.Error("name mismatch")
	}
	
	if lr.Email != email {
		t.Error("email mismatch")
	}
	
	// ---
	
	lber, lbee := loginByEmail(email, password)
	
	if lbee != nil {
		t.Error(lbee)
	}
	
	if lber.Id != cr.Id {
		t.Error("id mismatch")
	}
	
	if lber.Name != name {
		t.Error("name mismatch")
	}
	
	// ---
	
	name = Name("Dummy")
	password = Password("DummyDummy")
	
	// ---
	
	_, une := updateName(cr.Id, name)
	
	if une != nil {
		t.Error(une)
	}
	
	// ---
	
	fr2, fe2 := findById(cr.Id)
	
	if fe2 != nil {
		t.Error(fe2)
	}
	
	if fr2.Name != name {
		t.Error("name mismatch")
	}
	
	// ---
	
	_, upe := updatePassword(cr.Id, password)
	
	if upe != nil {
		t.Error(upe)
	}
	
	// ---
	
	fr3, fe3 := findById(cr.Id)
	
	if fe3 != nil {
		t.Error(fe3)
	}
	
	if fr3.PasswordHash != passwordHash(password, fr3.PasswordSalt) {
		t.Error("password hash mismatch")
	}
	
	// ---
	
	_, le2 := login(cr.Id, password)
	
	if le2 != nil {
		t.Error(le2)
	}
	
	// ---
	
	_, lbee2 := loginByEmail(email, password)
	
	if lbee2 != nil {
		t.Error(lbee2)
	}
	
	// ---
	
	_, de := destroy(cr.Id)
	
	if de != nil {
		t.Error(de)
	}
	
	// ---
	
	_, fe4 := findById(cr.Id)
	
	if fe4 != mgo.ErrNotFound {
		if fe4 != nil {
			t.Error(fe4)
		} else {
			t.Error("entry found")
		}
	}
}

// ---
