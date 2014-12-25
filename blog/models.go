package blog

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
)

// Objects
type BlogPost struct {
	DateCreated time.Time
	Title       string
	Content     string
}

func NewBlogPost(title, content string) BlogPost {
	return BlogPost{DateCreated: time.Now(), Title: title, Content: content}
}

func (b BlogPost) Print() {
	fmt.Println("Created at:", b.DateCreated)
	fmt.Println("Title:", b.Title)
	fmt.Println("Content:", b.Content)
}

func (b BlogPost) InsertToDb(d mgo.Database) {
	// Insert this into the database using mgo
}
