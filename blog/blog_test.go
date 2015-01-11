package blog

import "testing"

func TestBlogPostObject(t *testing.T) {
	//	b := NewBlogPost("This is a post", "Hello world!")
	//	AddPost(b)
}

func TestGetBlogPostById(t *testing.T) {
	EditPost(1, "Hello this is updated", "updated")
}
