package main

import (
	arcaflow_plugin_test_impl_go "arcaflow-plugin-test-impl-go"

	"go.flow.arcalot.io/pluginsdk/plugin"
)

func main() {
	plugin.Run(arcaflow_plugin_test_impl_go.WaitSchema)
}
