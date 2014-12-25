package blog

import "testing"

func TestBlogPostObject(t *testing.T) {
	b := NewBlogPost("This is a post", "Hello world!")
	AddPost(b)
}
