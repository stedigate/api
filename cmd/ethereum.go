package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stedigate/core/cmd/ethereum"
)

var ethereumCmd = &cobra.Command{
	Use:   "ethereum",
	Args:  cobra.MinimumNArgs(1),
	Short: "Command to interact with the Solana blockchain.",
	Long:  `Command to interact with the Solana blockchain. It can be used to fetch transactions, wallet balances, and other information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ethereum called")
	},
}

func init() {
	ethereumCmd.AddCommand(ethereum.GetTransactionCmd)
	ethereumCmd.AddCommand(ethereum.BalanceCmd)
	rootCmd.AddCommand(ethereumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ethereumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ethereumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
