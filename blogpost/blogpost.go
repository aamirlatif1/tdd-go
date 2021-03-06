package blogpost

import (
	"io/fs"
)

func PostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post := makePostFromFile(fileSystem, f.Name())
		posts = append(posts, post)
	}
	return posts, nil
}

func makePostFromFile(fileSystem fs.FS, fileName string) Post {
	blogFile, _ := fileSystem.Open(fileName)
	return newPost(blogFile)
}
