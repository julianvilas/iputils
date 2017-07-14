package iputils

import (
	"fmt"
	"net"
)

// ContainsIP checks if 'ip' is contained inside one of the 'networks' provided.
// It returns 'ok' when the the 'ip' is contained. Also it returns the 'network'
// where it's contained.
// Returns error if the ip or the networks can't be parsed.
func ContainsIP(ip string, networks ...string) (ok bool, network string, err error) {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false, "", fmt.Errorf("%v is not a valid IP address", ip)
	}

	for _, network = range networks {
		_, parsedNet, err := net.ParseCIDR(network)
		if err != nil {
			return false, "", fmt.Errorf("error parsing the network: %v", err)
		}

		if parsedNet.Contains(parsedIP) {
			return true, network, nil
		}
	}

	return false, "", nil
}
