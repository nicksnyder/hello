// Package hello is for testing purposes. You should not depend on this package.
package hello

import "fmt"

func Get(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
