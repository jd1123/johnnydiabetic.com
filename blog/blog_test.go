package blog

import "testing"

func TestBlogPostObject(t *testing.T) {
	b := NewBlogPost("First Post!", "Merry Christmas Eve! I'm going to work :(")
	b.Print()
}
