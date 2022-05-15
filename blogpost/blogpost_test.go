package blogpost_test

import (
	"errors"
	"io/fs"
	"reflect"
	"tdd-go/blogpost"
	"testing"
	"testing/fstest"
)

func TestBlogPostFromFS(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Given
		fs := fstest.MapFS{
			"hello-workd.md": {Data: []byte(`Title: Hello, TDD world!
			Description: lol
			Tags: tdd, go`)},
		}
		// When
		posts, err := blogpost.PostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		// Then
		assertEqual(t, len(fs), len(posts), "post size")
		assertEqual(t, blogpost.Post{
			Title:       "Hello, TDD world!",
			Description: "lol",
			Tags:        []string{"tdd", "go"}},
			posts[0], "post")
	})

	t.Run("failing fs", func(t *testing.T) {
		_, err := blogpost.PostsFromFS(FailingFS{})
		if err == nil {
			t.Error("error is expcted")
		}
	})

}

func assertEqual(t *testing.T, expected, actual interface{}, descs ...string) {
	desc := ""
	if len(descs) > 0 {
		desc = descs[0] + " "
	}
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%sexpected: %#v, actual: %#v", desc, expected, actual)
	}
}

type FailingFS struct {
}

func (f FailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("file not found")
}
