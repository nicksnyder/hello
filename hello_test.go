package hello_test

import (
	"testing"

	"github.com/nicksnyder/hello/v2"
)

func TestHelloGet(t *testing.T) {
	expected := "Hello Nick!"
	if h := hello.Get("Nick"); h != expected {
		t.Fatalf(`expected %q; got %q`, expected, h)
	}
}
