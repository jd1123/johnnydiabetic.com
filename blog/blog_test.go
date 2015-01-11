package blog

import (
	"fmt"
	"testing"

	"github.com/kennygrant/sanitize"
)

func TestBlogPostObject(t *testing.T) {
	//	b := NewBlogPost("This is a post", "Hello world!")
	//	AddPost(b)
}

func TestGetBlogPostById(t *testing.T) {
}

func TestSanitize(t *testing.T) {
	htmlString := "<html><head></head><body>This is html<p>This is a new Paragraph</p></body></html>"
	plainString := sanitize.HTML(htmlString)
	fmt.Println("HTML:", htmlString)
	fmt.Println("Sanitized:", plainString)
}
