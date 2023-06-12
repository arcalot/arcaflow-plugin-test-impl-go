package main

import (
	testimplplugin "go.flow.arcalot.io/plugintestimpl"

	"go.flow.arcalot.io/pluginsdk/plugin"
)

func main() {
	plugin.Run(testimplplugin.WaitSchema)
}
