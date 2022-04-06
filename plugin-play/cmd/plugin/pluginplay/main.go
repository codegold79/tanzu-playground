package main

import (
	"log"
	"sync"

	cliv1alpha1 "github.com/vmware-tanzu/tanzu-framework/apis/cli/v1alpha1"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/cli/command/plugin"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/cli/signals"
)

var descriptor = cliv1alpha1.PluginDescriptor{
	Name:        "pluginplay",
	Description: "create a very simple plugin to test and play with",
	Version:     "v0.0.1",
	Group:       cliv1alpha1.ManageCmdGroup, // set group
}

func main() {
	p, err := plugin.NewPlugin(&descriptor)
	if err != nil {
		log.Fatal(err)
	}
	p.AddCommands(
		HelloWorldCmd,
	)

	ctx := signals.SetupSignalHandler()
	p.AddContext(ctx)

	var wg sync.WaitGroup
	errs := make(chan error)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range errs {
			log.Printf("%s: error: %s\n", descriptor.Name, err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(errs)
		if err = p.Execute(); err != nil {
			errs <- err
		}
	}()

	wg.Wait()
}
