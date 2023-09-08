package arcaflow_plugin_test_impl_go

import (
	"context"
	"fmt"
	"go.flow.arcalot.io/pluginsdk/plugin"
	"go.flow.arcalot.io/pluginsdk/schema"
	"time"
)

type Input struct {
	WaitTime int `json:"wait_time_ms"`
}

// We define a separate scope, so we can add subobjects later.
var inputSchema = schema.NewScopeSchema(
	// Struct-mapped object schemas are object definitions that are mapped to a specific struct (Input)
	schema.NewStructMappedObjectSchema[Input](
		// ID for the object:
		"input",
		// Properties of the object:
		map[string]*schema.PropertySchema{
			"wait_time_ms": schema.NewPropertySchema(
				// Type properties:
				schema.NewIntSchema(schema.PointerTo[int64](0), nil, nil),
				// Display metadata:
				schema.NewDisplayValue(
					schema.PointerTo("Wait Time"),
					schema.PointerTo("How long to wait."),
					nil,
				),
				// Required:
				true,
				// Required if:
				[]string{},
				// Required if not:
				[]string{},
				// Conflicts:
				[]string{},
				// Default value, JSON encoded:
				nil,
				//Examples:
				nil,
			),
		},
	),
)

type Output struct {
	Message string `json:"message"`
}

var outputSchema = schema.NewScopeSchema(
	schema.NewStructMappedObjectSchema[Output](
		"output",
		map[string]*schema.PropertySchema{
			"message": schema.NewPropertySchema(
				schema.NewStringSchema(nil, nil, nil),
				schema.NewDisplayValue(
					schema.PointerTo("Message"),
					schema.PointerTo("The resulting message."),
					nil,
				),
				true,
				nil,
				nil,
				nil,
				nil,
				nil,
			),
		},
	),
)

func cancelWait_(ctx context.Context, waitData *WaitStepData, input plugin.CancelInput) {
	waitData.cancellationChannel <- true
}

func wait_(ctx context.Context, waitData *WaitStepData, input Input) (string, any) {
	start := time.Now()
	select {
	case <-time.After(time.Duration(input.WaitTime) * time.Millisecond):
		return "success", Output{
			fmt.Sprintf("Plugin slept for %d ms.", input.WaitTime),
		}
	case <-ctx.Done(): // Terminated
		duration := time.Since(start)
		return "terminated_early", Output{
			fmt.Sprintf("Plugin cancelled early due to context done after %d ms after scheduled to sleep for %d ms.",
				duration.Milliseconds(), input.WaitTime),
		}
	case _ = <-waitData.cancellationChannel: // Cancelled
		duration := time.Since(start)
		return "cancelled_early", Output{
			fmt.Sprintf("Plugin cancelled early from signal after %d ms after scheduled to sleep for %d ms.",
				duration.Milliseconds(), input.WaitTime),
		}
	}

}

type WaitStepData struct {
	cancellationChannel chan bool
}

func waitInitializer_() *WaitStepData {
	return &WaitStepData{
		make(chan bool, 0),
	}
}

var WaitSchema = schema.NewCallableSchema(
	schema.NewCallableStepWithSignals[*WaitStepData, Input](
		// ID of the step:
		"wait",
		// Add the input schema:
		inputSchema,
		map[string]*schema.StepOutputSchema{
			// Define possible outputs:
			"success": schema.NewStepOutputSchema(
				// Add the output schema:
				outputSchema,
				schema.NewDisplayValue(
					schema.PointerTo("Success"),
					schema.PointerTo("Successfully waited"),
					nil,
				),
				false,
			),
			"cancelled_early": schema.NewStepOutputSchema(
				outputSchema,
				schema.NewDisplayValue(
					schema.PointerTo("Cancelled Early"),
					schema.PointerTo("Was cancelled before the expected wait period passed."),
					nil,
				),
				false,
			),
			"terminated_early": schema.NewStepOutputSchema(
				outputSchema,
				schema.NewDisplayValue(
					schema.PointerTo("Terminated Early"),
					schema.PointerTo("Was terminated before the expected wait period passed."),
					nil,
				),
				false,
			),
		},
		map[string]schema.CallableSignal{
			plugin.CancellationSignalSchema.ID(): schema.NewCallableSignalFromSchema(plugin.CancellationSignalSchema, cancelWait_),
		},
		// No emitted signals
		map[string]*schema.SignalSchema{},
		// Metadata for the function:
		schema.NewDisplayValue(
			schema.PointerTo("Wait"),
			schema.PointerTo("Wait for specified time."),
			nil,
		),
		waitInitializer_,
		// Reference the function
		wait_,
	),
)
