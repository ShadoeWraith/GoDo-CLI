package cmd

import "fmt"

func Delete(args []string) {
	fmt.Println("Deleting...")

	if len(args) < 2 {
		fmt.Println("Usage: godo delete -id <id>")
		return
	}
}
