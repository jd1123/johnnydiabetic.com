package server

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	UserId   string
	Password []byte
}

func (u *User) SetPassword(pw string) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = hpass
}
