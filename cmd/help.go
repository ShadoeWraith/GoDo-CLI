package cmd

import "fmt"

func Help() {
	fmt.Println("--- Godo Help ---")

	fmt.Println("\ngodo list")
	fmt.Println("godo create -title? <title> -description? <description>")
	fmt.Println("godo update -id? <id> -title? <title> -description? <description>")
	fmt.Println("godo delete -id? <id>")

	fmt.Println("\ngodo -help")
	fmt.Println("godo -version")
	fmt.Println("godo -setuser <userID>")
}
