# Unexpected Behavior When Accessing Non-Existent Keys in Go Maps

This example demonstrates a potential pitfall in Go when working with maps: accessing a non-existent key does not trigger a runtime panic, but returns the zero value of the map's value type.  This can lead to subtle bugs that are difficult to track down.

The `bug.go` file shows the issue. The `bugSolution.go` file provides a solution using the comma ok idiom.
