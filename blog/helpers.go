package blog

func PaginateBlogPosts(b []BlogPost, postsPerPage int) [][]BlogPost {
	l_b := len(b)
	numPages := l_b / postsPerPage
	pages := make([][]BlogPost, numPages)

	for i := 0; i < numPages; i++ {
		start := i * postsPerPage
		end := (i + 1) * postsPerPage
		if end > l_b {
			end = l_b
		}
		pages[i] = b[start:end]
	}

	return pages
}
