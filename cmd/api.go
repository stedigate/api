package cmd

import (
	"github.com/pushgate/core/cmd/api"
	"github.com/spf13/cobra"
)

var Port int
var Env string

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "API to interact with the stedicoin gateway.",
	Long:  `API to interact with the stedicoin gateway. It can be used to fetch transactions, wallet balances, and other information.`,
	Run: func(cmd *cobra.Command, args []string) {

		api.Run(Port, Env)
	},
}

func init() {
	apiCmd.PersistentFlags().IntVarP(&Port, "port", "p", 4000, "api server port")
	apiCmd.PersistentFlags().StringVarP(&Env, "environment", "e", "development", "api server environment (development|staging|production)")
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
