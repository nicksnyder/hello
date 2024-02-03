package hello_test

import (
	"testing"

	"github.com/nicksnyder/hello"
)

func TestHelloGet(t *testing.T) {
	if h := hello.Get(); h != "hello" {
		t.Fatalf(`expected "hello"; got %q`, h)
	}
}
