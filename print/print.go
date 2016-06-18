package print

import (
	"fmt"
)

func Println(content string) error {
	_, err := fmt.Printf("%s", content)
	return err
}
