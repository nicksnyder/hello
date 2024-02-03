package hello_test

import (
	"testing"

	"github.com/nicksnyder/hello"
)

func TestHelloGet(t *testing.T) {
	expected := "Hello!"
	if h := hello.Get(); h != expected {
		t.Fatalf(`expected %q; got %q`, expected, h)
	}
}
