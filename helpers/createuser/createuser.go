package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
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

func NewUser(username, pw string) User {
	u := User{UserId: username}
	u.SetPassword(pw)
	return u
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("User creator")
	fmt.Println("Please enter username:")
	user, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input Error:", err)
		os.Exit(1)
	}
	user = user[:len(user)-1]

	s, err := mgo.Dial("localhost")
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
	c := s.DB("test").C("users")

	result := User{}
	err = c.Find(bson.M{"userid": user}).One(&result)
	if err == nil {
		fmt.Println("User already exists. Exiting...")
		os.Exit(1)
	}
	fmt.Println(result)

	fmt.Println("Please enter password:")
	pw, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input Error:", err)
		os.Exit(1)
	}
	pw = pw[:len(pw)-1]
	u := NewUser(user, pw)

	err = c.Insert(u)
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("User", user, "saved successfully")

}
