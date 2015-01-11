package blog

import (
	"fmt"
	"time"
)

// Objects
type BlogPost struct {
	DateCreated   time.Time
	Title         string
	Content       string
	Id            int
	CreatedString string
}

func NewBlogPost(title, content string) BlogPost {
	// This is trying to recover from a runtime panic of
	// being out of bounds if no posts exist. I have not
	// figured out how to return something if it panics.
	// There is probably an easier way to do this:
	// check if collection.Count()==0 and if so make
	// id = 1 else check the id
	g := func(b BlogPost) {
		if r := recover(); r != nil {
		}
	}
	var b BlogPost = BlogPost{}
	b.Title = title
	b.Content = content
	b.DateCreated = time.Now()
	b.CreatedString = time.Now().Format("Jan 2, 2006 at 3:04pm")
	defer g(b)
	b.Id = GetAllPosts()[0].Id + 1
	return b
	//return BlogPost{DateCreated: time.Now(), Title: title, Content: content, Id: id, CreatedString: createdString}
}

func (b BlogPost) Print() {
	fmt.Println("Created at:", b.DateCreated)
	fmt.Println("Title:", b.Title)
	fmt.Println("Content:", b.Content)
}

type BlogPostCollection []BlogPost

func (b BlogPostCollection) Len() int {
	return len(b)
}

func (b BlogPostCollection) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b BlogPostCollection) Less(i, j int) bool {
	return b[j].DateCreated.Before(b[i].DateCreated)
}
