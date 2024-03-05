// Package goodbye is for testing purposes. You should not depend on this package.
package goodbye

import "fmt"

func Get(name string) string {
	return fmt.Sprintf("Goodbye %s!", name)
}
