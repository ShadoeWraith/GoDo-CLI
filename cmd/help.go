package cmd

import "fmt"

func Help() {
	fmt.Println("--- Godo Help ---")
	fmt.Println("A list of commands for Godo CLI")

	fmt.Println("\n--- Offline Commands ---")
	fmt.Println("godo list")
	fmt.Println("godo create -title? <title> -description? <description>")
	fmt.Println("godo update")
	fmt.Println("godo delete")

	fmt.Println("\ngodo -help")
	fmt.Println("godo -version")

	fmt.Println("\n--- API Commands ---")
	fmt.Println("godo pull")
	fmt.Println("godo push")

	fmt.Println("\ngodo -setuser <userID>")
}
