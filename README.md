# Testing plugin for Arcaflow

This plugin is used for standardized integration tests. It can be used directly in Go, or it can be deployed with any of the supported container deployers.

## Steps

### wait
This plugin has one step called `wait`, which waits for the duration specified by the input `wait_time_ms`

There are two possible outputs for step wait:
- success: When it waits the expected duration.
- cancelled_early: When it is cancelled at or before the end of its expected wait period.  

The output is predictable for testing purposes, with the success output just outputting the expected sleep time, and the cancelled_early output outputting the expected and actual wait time.