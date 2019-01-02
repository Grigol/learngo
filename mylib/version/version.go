package version

import "runtime"

// Version is a function that will print the current istalled Go Version
// in the running PC.
func Version() string {
	return runtime.Version()
}
