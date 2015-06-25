package v1

// ---
// ---
// ---

import (
	"time"
	"errors"
	
	// ---
	
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	
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

func Create(name Name, email Email, verified Verified, password Password) (Id, error) {
	id := Id(uuid.NewRandom().String())
	
	// ---
	
	passwordSalt, passwordSaltErr := passwordSalt()
	
	if passwordSaltErr != nil {
		return "", passwordSaltErr
	}
	
	// ---
	
	passwordHash := passwordHash(password, passwordSalt)
	
	// ---
	
	entry := UserEntry{
		ObjectId: bson.NewObjectId(),
		Id: id,
		Name: name,
		Email: email,
		Verified: verified,
		PasswordSalt: passwordSalt,
		PasswordHash: passwordHash,
	}
	
	// ---
	
	insertErr := MongoCollection.Insert(entry)
	
	if insertErr != nil {
		return "", insertErr
	}
	
	// ---
	
	return id, nil
}

func Destroy(id Id) (error) {
	removeErr := MongoCollection.Remove(bson.M{"id": id})
	
	if removeErr != nil {
		return removeErr
	}
	
	// ---
	
	return nil
}

func Query(id Id) (Name, Email, Verified, error) {
	result := UserEntry{}
	
	findErr := MongoCollection.Find(bson.M{"id": id}).One(&result)
	
	if findErr != nil {
		if findErr == mgo.ErrNotFound {
			return "", "", false, ErrNotFound
		} else {
			return "", "", false, findErr
		}
	}
	
	// ---
	
	return result.Name, result.Email, result.Verified, nil
}

func QueryByEmail(email Email) (Id, Name, Verified, error) {
	result := UserEntry{}
	
	findErr := MongoCollection.Find(bson.M{"email": email}).One(&result)
	
	if findErr != nil {
		if findErr == mgo.ErrNotFound {
			return "", "", false, ErrNotFound
		} else {
			return "", "", false, findErr
		}
	}
	
	// ---
	
	return result.Id, result.Name, result.Verified, nil
}

func Login(id Id, password Password) (Name, Email, Verified, error) {
	result := UserEntry{}
	
	findErr := MongoCollection.Find(bson.M{"id": id}).One(&result)
	
	if findErr != nil {
		if findErr == mgo.ErrNotFound {
			return "", "", false, ErrNotFound
		} else {
			return "", "", false, findErr
		}
	}
	
	// ---
	
	if result.PasswordHash != passwordHash(password, result.PasswordSalt) {
		return "", "", false, ErrPasswordMismatch
	}
	
	// ---
	
	return result.Name, result.Email, result.Verified, nil
}

func LoginByEmail(email Email, password Password) (Id, Name, Verified, error) {
	result := UserEntry{}
	
	findErr := MongoCollection.Find(bson.M{"email": email}).One(&result)
	
	if findErr != nil {
		if findErr == mgo.ErrNotFound {
			return "", "", false, ErrNotFound
		} else {
			return "", "", false, findErr
		}
	}
	
	// ---
	
	if result.PasswordHash != passwordHash(password, result.PasswordSalt) {
		return "", "", false, ErrPasswordMismatch
	}
	
	// ---
	
	return result.Id, result.Name, result.Verified, nil
}

func UpdateName(id Id, name Name) (error) {
	updateErr := MongoCollection.Update(bson.M{"id": id}, bson.M{"$set": bson.M{"name": name}})
	
	if updateErr != nil {
		return updateErr
	}
	
	// ---
	
	return nil
}

func UpdatePassword(id Id, password Password) (error) {
	result := UserEntry{}
	
	findErr := MongoCollection.Find(bson.M{"id": id}).One(&result)
	
	if findErr != nil {
		if findErr == mgo.ErrNotFound {
			return ErrNotFound
		} else {
			return findErr
		}
	}
	
	// ---
	
	updateErr := MongoCollection.Update(bson.M{"id": id}, bson.M{"$set": bson.M{"passwordHash": passwordHash(password, result.PasswordSalt)}})
	
	if updateErr != nil {
		return updateErr
	}
	
	// ---
	
	return nil
}

func StartVerify(id Id) (Token, error) {
	result := UserEntry{}
	
	findErr := MongoCollection.Find(bson.M{"id": id}).One(&result)
	
	if findErr != nil {
		if findErr == mgo.ErrNotFound {
			return "", ErrNotFound
		} else {
			return "", findErr
		}
	}
	
	// ---
	
	exp := time.Now().Add(time.Hour * 1)
	
	token, tokenErr := makeSimpleJwt(VerifyKey, string(id), exp)
	
	if tokenErr != nil {
		return "", tokenErr
	}
	
	// ---
	
	return Token(token), nil
}

func Verify(token Token) (error) {
	claims, claimsErr := parseJwt(VerifyKey, string(token))
	
	if claimsErr != nil {
		return claimsErr
	}
	
	// ---
	
	updateErr := MongoCollection.Update(bson.M{"id": claims["val"]}, bson.M{"$set": bson.M{"verified": true}})
	
	if updateErr != nil {
		return updateErr
	}
	
	// ---
	
	return nil
}

func StartReset(id Id) (Token, error) {
	result := UserEntry{}
	
	findErr := MongoCollection.Find(bson.M{"id": id}).One(&result)
	
	if findErr != nil {
		if findErr == mgo.ErrNotFound {
			return "", ErrNotFound
		} else {
			return "", findErr
		}
	}
	
	// ---
	
	exp := time.Now().Add(time.Minute * 5)
	
	// ---
	
	subKey := hash512([]byte(result.PasswordHash), nil)
	
	// ---
	
	subToken, subTokenErr := makeSimpleJwt(subKey, "lav", exp)
	
	if subTokenErr != nil {
		return "", subTokenErr
	}
	
	// ---
	
	token, tokenErr := makeTokenJwt(ResetKey, string(id), subToken, exp)
	
	if tokenErr != nil {
		return "", tokenErr
	}
	
	// ---
	
	return Token(token), nil
}

func Reset(token Token, password Password) (error) {
	claims, claimsErr := parseJwt(ResetKey, string(token))
	
	if claimsErr != nil {
		return claimsErr
	}
	
	// ---
	
	result := UserEntry{}
	
	findErr := MongoCollection.Find(bson.M{"id": claims["val"]}).One(&result)
	
	if findErr != nil {
		if findErr == mgo.ErrNotFound {
			return ErrNotFound
		} else {
			return findErr
		}
	}
	
	// ---
	
	subKey := hash512([]byte(result.PasswordHash), nil)
	
	// ---
	
	_, subClaimsErr := parseJwt(subKey, claims["tkn"].(string))
	
	if subClaimsErr != nil {
		return subClaimsErr
	}
	
	// ---
	
	updateErr := MongoCollection.Update(bson.M{"id": claims["val"]}, bson.M{"$set": bson.M{"passwordHash": passwordHash(password, result.PasswordSalt)}})
	
	if updateErr != nil {
		return updateErr
	}
	
	// ---
	
	return nil
}

// ---
