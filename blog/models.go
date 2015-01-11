package blog

import (
	"fmt"
	"strings"
	"time"

	"github.com/kennygrant/sanitize"
	"github.com/russross/blackfriday"
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
	var b BlogPost = BlogPost{}
	b.Title = title
	b.Content = content
	b.DateCreated = time.Now()
	b.CreatedString = time.Now().Format("Jan 2, 2006 at 3:04pm")
	allPosts := GetAllPosts()
	if len(allPosts) == 0 {
		b.Id = 0
	} else {
		b.Id = allPosts[0].Id + 1
	}
	return b
}

func (b BlogPost) Print() {
	fmt.Println("Created at:", b.DateCreated)
	fmt.Println("Title:", b.Title)
	fmt.Println("Content:", b.Content)
}

func (b *BlogPost) Markdown() {
	b.Content = string(blackfriday.MarkdownBasic([]byte(b.Content)))
}

func (b *BlogPost) Sanitize() {
	b.Markdown()
	b.Content = sanitize.HTML(b.Content)
}

func (b *BlogPost) Slugify() {
	// FIXME: This is okay, but not great -
	// does not leave spaces after newlines
	// ... got to think about this
	b.Sanitize()
	s := strings.Split(b.Content, " ")
	if len(s) > 20 {
		b.Content = strings.Join(s[0:19], " ")
	}
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
