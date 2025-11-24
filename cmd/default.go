package cmd

import (
	"fmt"
	"godo-cli/pkg/common"
	"os"
)

func Default(command string) {
	fmt.Fprintf(os.Stderr, "%sError%s: Unknown command '%s'\n\n", common.ColorError, common.ColorReset, command)
	fmt.Printf("For help use\n   go run main.go -help\n\n")
	os.Exit(1)
}
