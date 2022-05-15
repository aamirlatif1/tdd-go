package blogpost_test

import (
	"tdd-go/blogpost"
	"testing"
	"testing/fstest"
)

func TestBlogPost(t *testing.T) {
	// Given
	fs := fstest.MapFS{
		"hello-workd.md":  {Data: []byte("hello, workd")},
		"hello-twitch.md": {Data: []byte("hello, twitch")},
	}
	// When
	posts := blogpost.PostsFromFS(fs)

	// Then
	if len(posts) != len(fs) {
		t.Errorf("expected %d posts, got %d posts", len(posts), len(fs))
	}
}
