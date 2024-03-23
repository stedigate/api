package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stedigate/core/cmd/tron"
)

// trc20EventsCmd represents the tronTrc20Events command
var trc20EventsCmd = &cobra.Command{
	Use:   "trc20Events",
	Short: "Fetches TRC20 events from the Tron blockchain.",
	Long:  `Fetches TRC20 events from the Tron blockchain. It can be used to fetch all events from a specific contract address.`,
	Run: func(cmd *cobra.Command, args []string) {
		tron.Run()
	},
}

func init() {
	tronCmd.AddCommand(trc20EventsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tronTrc20EventsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tronTrc20EventsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
