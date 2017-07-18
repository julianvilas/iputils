package cmd

import (
	"fmt"

	"github.com/julianvilas/iputils"
	"github.com/spf13/cobra"
)

var removeNetAndBroadcast bool

// expandCmd represents the scan command
var expandCmd = &cobra.Command{
	Use:   "expand <network>",
	Short: "Prints all the IPs contained in the network provided",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("incorrect number of args, want 1, got %v", len(args))
		}

		network := args[0]

		ips, err := iputils.ExpandCIDR(network, removeNetAndBroadcast)
		if err != nil {
			return err
		}

		for _, ip := range ips {
			fmt.Println(ip)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(expandCmd)

	expandCmd.Flags().BoolVarP(&removeNetAndBroadcast, "rm", "r", false, "do not print the network and broadcast addresses")
}
