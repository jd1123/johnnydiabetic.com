package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type Us struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	UserId      string
	Password    []byte
	Permissions map[string]bool
}

func (u *Us) SetPassword(pw string) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = hpass
}
