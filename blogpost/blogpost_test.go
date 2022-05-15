package blogpost_test

import (
	"tdd-go/blogpost"
	"testing"
	"testing/fstest"
)

func TestBlogPost(t *testing.T) {
	// Given
	fs := fstest.MapFS{
		"hello-workd.md": {Data: []byte("Title: Hello, TDD world!")},
	}
	// When
	posts := blogpost.PostsFromFS(fs)

	// Then
	if len(posts) != len(fs) {
		t.Errorf("expected %d posts, got %d posts", len(posts), len(fs))
	}

	expectedFirstPOst := blogpost.Post{Title: "Hello, TDD world!"}
	if posts[0] != expectedFirstPOst {
		t.Errorf("got %#v, want %#v", posts[0], expectedFirstPOst)
	}

}
