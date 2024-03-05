package goodbye_test

import (
	"testing"

	"github.com/nicksnyder/hello/v2/goodbye"
)

func TestHelloGet(t *testing.T) {
	expected := "Goodbye Nick!"
	if h := goodbye.Get("Nick"); h != expected {
		t.Fatalf(`expected %q; got %q`, expected, h)
	}
}
