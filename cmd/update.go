package cmd

import "fmt"

func Update(args []string) {
	fmt.Println("Updating...")

	if len(args) < 2 {
		fmt.Println("Usage: godo update -id <id> -title <title> -description <description>")
		return
	}
}
