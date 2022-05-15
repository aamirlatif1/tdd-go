package blogpost

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

func newPost(blogFile io.Reader) Post {

	scanner := bufio.NewScanner(blogFile)

	readLine := func(prefix string) string {
		scanner.Scan()
		title := strings.TrimPrefix(strings.Trim(scanner.Text(), "\t"), prefix)
		return title
	}

	return Post{
		Title:       readLine("Title: "),
		Description: readLine("Description: "),
		Tags:        strings.Split(readLine("Tags: "), ", "),
	}
}
