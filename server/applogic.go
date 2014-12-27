package server

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func login(username, pw string) (*User, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Println("Database error")
		return nil, err
	}
	c := session.DB("test").C("users")
	result := User{}
	err = c.Find(bson.M{"userid": username}).One(&result)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(result.Password, []byte(pw))
	if err != nil {
		return nil, err
	}
	return &result, nil
}
