package blogpost

import (
	"io/fs"
	"testing/fstest"
)

type Post struct {
}

func PostsFromFS(fileSystem fstest.MapFS) []Post {
	dir, _ := fs.ReadDir(fileSystem, ".")
	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}
	return posts
}
