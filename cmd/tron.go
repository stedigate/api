package cmd

import (
	"fmt"
	"github.com/stedigate/core/cmd/tron"

	"github.com/spf13/cobra"
)

// tronCmd represents the tron command
var tronCmd = &cobra.Command{
	Use:   "tron",
	Args:  cobra.MinimumNArgs(1),
	Short: "Command to interact with the Tron blockchain.",
	Long:  `Command to interact with the Tron blockchain. It can be used to fetch transactions, wallet balances, and other information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tron called")
	},
}

func init() {
	tronCmd.AddCommand(tron.TrackTrc20Cmd)
	tronCmd.AddCommand(tron.BalanceCmd)
	tronCmd.AddCommand(tron.TransferCmd)
	tronCmd.AddCommand(tron.SimulateCmd)
	rootCmd.AddCommand(tronCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tronCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tronCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
