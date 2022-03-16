package main

import (
	"fmt"
	"time"

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

	ticker := time.NewTicker(2500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		var i int
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				if i%2 == 0 {
					fmt.Println(i, "tick")
				} else {
					fmt.Println(i, "tock")
				}
				i++
			}
		}
	}()

	time.Sleep(1 * time.Minute)
	ticker.Stop()
	done <- true
	fmt.Println("Tick-tock stopped")

	return nil
}
