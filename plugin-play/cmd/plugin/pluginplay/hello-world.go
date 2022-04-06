package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

const timerPeriod time.Duration = 5 * time.Second

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

	var i int
	ctx := cmd.Context()
	startTime := time.Now()

	for {
		select {
		case <-time.After(1 * time.Second):
			tickTock(i)
			i++

			if time.Since(startTime) >= timerPeriod {
				fmt.Println("Bye, World.")
				return nil
			}
		case <-ctx.Done():
			log.Println("early shutdown")
			return nil
		}
	}
}

func tickTock(i int) {
	if i%2 == 0 {
		fmt.Println(i, "tick")
	} else {
		fmt.Println(i, "tock")
	}
}
