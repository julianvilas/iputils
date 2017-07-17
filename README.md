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

A naive command is also given that allows using the lib from the command-line.
```bash
$ iputils
Usage: iputils [flags] ip network1 [network2 ...]

$ iputils 192.168.1.1 192.168.1.0/24
192.168.1.0/24

$ iputils 192.168.1.1 192.168.2.0/24

$ iputils 192.168.1.1 192.168.3.0/24 192.168.1.0/24
192.168.1.0/24

$ cat /tmp/networks
192.168.1.0/24
192.168.2.0/24

$ iputils 192.168.1.0 $(cat /tmp/networks | xargs)
192.168.1.0/24
```
