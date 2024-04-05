package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stedigate/core/cmd/solana"
)

var solanaCmd = &cobra.Command{
	Use:   "solana",
	Args:  cobra.MinimumNArgs(1),
	Short: "Command to interact with the Solana blockchain.",
	Long:  `Command to interact with the Solana blockchain. It can be used to fetch transactions, wallet balances, and other information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("solana called")
	},
}

func init() {
	tronCmd.AddCommand(solana.GetTransactionCmd)
	rootCmd.AddCommand(solanaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// solanaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// solanaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
