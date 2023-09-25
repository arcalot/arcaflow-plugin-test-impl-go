package main

import (
	testimplplugin "go.flow.arcalot.io/testplugin"

	"go.flow.arcalot.io/pluginsdk/plugin"
)

func main() {
	plugin.Run(testimplplugin.TestStepsSchema)
}
