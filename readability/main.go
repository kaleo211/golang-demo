package main

import (
	"fmt"
)

func summary(user string) (bike string, aveSpeed int, err error) {
	if user == "" {
		return "", 0, fmt.Errorf("user name is empty\n")
	}

	return "Specialized", 10, nil
}

func main() {
	user := "Peter"

	bike, ave, err := summary(user)
	if err != nil {
		fmt.Printf("error getting summary: %v\n", err)
	}

	fmt.Printf("%s ride on %s at %dmph\n", user, bike, ave)
}
