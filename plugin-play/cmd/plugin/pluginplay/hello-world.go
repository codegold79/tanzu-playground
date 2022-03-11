package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var HelloWorldCmd = &cobra.Command{
	Use:   "hi",
	Short: "say Hello World",
	Args:  cobra.NoArgs,
	Example: `
	tanzu pluginplay hi`,
	RunE: helloWorld,
}

func helloWorld(cmd *cobra.Command, _ []string) error {
	fmt.Println("Hello, World.")
	return nil
}