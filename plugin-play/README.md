# plugin-play

Welcome, lets get started with your new Tanzu plugin repository.

Note: this README.md and the boostrapping for creating a plugin was made using,

```bash
tanzu builder tanzu builder init plugin-basic
```

More instructions are in the [tanzu-framework documentation](https://github.com/vmware-tanzu/tanzu-framework/blob/main/docs/cli/plugin_implementation_guide.md).

I removed the auto-generated `.git` and `.github` directories within `plugin-basic` because I'd like to have the plugin as a sub-project in this `tanzu-playground` repo.

## Getting Started

First, run `make init` to ensure your plugins go.sum is generated.

## Generate a plugin

Add a plugin with the tanzu cli: `tanzu builder cli add-plugin myplugin`

## Directory Structure

artifacts/: Where you plugins build output will be located

cmd/plugin/<plugin>: Path where you plugins main directory will live

cmd/plugin/<plugin>/test: Plugins are required to have a test command defined

## Commands

`make build`: builds your plugin artifacts
`make lint`: run the golangci linter on your plugin code
