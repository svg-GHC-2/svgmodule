package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/svg-at-wss/secretmod"

	"svgmodule/svgpkg"
)

func main() {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Modules!")

			svgpackage.PrintHello()
		},
	}

	fmt.Println("Calling cmd.Execute()!")
	cmd.Execute()
	secretmod.PrintSecretHello()
}