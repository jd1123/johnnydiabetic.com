package server

import "golang.org/x/crypto/bcrypt"

type User struct {
	UserId   string
	password []byte
}

func (u *User) SetPassword(pw string) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.password = hpass
}
