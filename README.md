# iputils

A Go package to perform operations over IP addresses and network ranges, using strings. 

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

A naive command is also given that allows using the package from the command-line.
```bash
$ iputils
A command to perform operations over IP addresses and network ranges

Usage:
  iputils [command]

Available Commands:
  contains    Checks if the IP address is contained in one of the networks provided
  expand      Prints all the IPs contained in the network provided
  help        Help about any command

Flags:
  -h, --help   help for iputils

Use "iputils [command] --help" for more information about a command.

$ iputils contains 192.168.1.32 192.168.1.0/24
192.168.1.0/24

$ iputils expand 192.168.1.0/31
192.168.1.0
192.168.1.1
```
