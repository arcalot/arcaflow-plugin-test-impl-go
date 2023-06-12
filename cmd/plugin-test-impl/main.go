package main

import (
	testimplplugin "go.flow.arcalot.io/plugin-test-impl"

	"go.flow.arcalot.io/pluginsdk/plugin"
)

func main() {
	plugin.Run(testimplplugin.WaitSchema)
}
