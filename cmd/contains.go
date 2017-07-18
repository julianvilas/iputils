package cmd

import (
	"fmt"

	"github.com/julianvilas/iputils"
	"github.com/spf13/cobra"
)

// containsCmd represents the scan command
var containsCmd = &cobra.Command{
	Use:   "contains <ip> <network>...",
	Short: "Checks if the IP address is contained in one of the networks provided",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("incorrect number of args, want 2, got %v", len(args))
		}

		ip := args[0]
		networks := args[1:]

		ok, network, err := iputils.ContainsIP(ip, networks...)
		if err != nil {
			return err
		}

		if ok {
			fmt.Println(network)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(containsCmd)
}
