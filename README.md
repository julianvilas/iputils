# iputils

Is a Go library to check whether an IP address pertains to a list of IP networks. 

### Installation

```
go get github.com/julianvilas/iputils 
```

### Example 

```
package main

import (
    "fmt"

    "github.com/julianvilas/iputils"
)

func main() {
    networks := []string{"192.168.0.0/24"}
    fmt.Println(iputils.ContainsIP("192.168.0.1", networks...))
}
```
