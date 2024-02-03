package hello

import "fmt"

func Get(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
