package blogpost

import (
	"bufio"
	"io"
	"strings"
)

const (
	titlePrefix       = "Title: "
	descriptionPrefix = "Description: "
	tagsPrefix        = "Tags: "
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
		Title:       readLine(titlePrefix),
		Description: readLine(descriptionPrefix),
		Tags:        strings.Split(readLine(tagsPrefix), ", "),
	}
}
