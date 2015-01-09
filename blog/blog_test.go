package blog

import (
	"fmt"
	"testing"
)

func TestBlogPostObject(t *testing.T) {
	//	b := NewBlogPost("This is a post", "Hello world!")
	//	AddPost(b)
}

func TestGetBlogPostById(t *testing.T) {
	b := GetPostById(1)
	fmt.Println(b)
}
