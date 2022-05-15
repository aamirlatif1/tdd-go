package blogpost_test

import (
	"errors"
	"io/fs"
	"tdd-go/blogpost"
	"testing"
	"testing/fstest"
)

func TestBlogPost(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Given
		fs := fstest.MapFS{
			"hello-workd.md": {Data: []byte("Title: Hello, TDD world!")},
		}
		// When
		posts, err := blogpost.PostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		// Then
		if len(posts) != len(fs) {
			t.Errorf("expected %d posts, got %d posts", len(posts), len(fs))
		}

		expectedFirstPOst := blogpost.Post{Title: "Hello, TDD world!"}
		if posts[0] != expectedFirstPOst {
			t.Errorf("got %#v, want %#v", posts[0], expectedFirstPOst)
		}
	})

	t.Run("failing fs", func(t *testing.T) {
		_, err := blogpost.PostsFromFS(FailingFS{})
		if err == nil {
			t.Error("error is expcted")
		}
	})

}

type FailingFS struct {
}

func (f FailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("file not found")
}
