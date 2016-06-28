package fmt

import (
	"fmt"
)

func Println(content string) error {
	_, err := fmt.Printf("%s\n", content)
	return err
}
