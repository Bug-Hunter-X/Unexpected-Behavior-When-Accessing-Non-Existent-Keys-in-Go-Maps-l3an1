```go
package main

import "fmt"

func main() {
    var m map[string]int
    value, ok := m["a"]
    if ok {
        fmt.Println("Value of a:", value)
    } else {
        fmt.Println("Key 'a' not found in the map")
    }
}
```