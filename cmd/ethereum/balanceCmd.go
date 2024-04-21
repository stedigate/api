package ethereum

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stedigate/core/internal/config"
	"github.com/stedigate/core/pkg/redis"
)

// BalanceCmd represents the tronTrc20Events command
var BalanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Get the ETH/USDT/USDC balance of a wallet address",
	Long:  `Get the ETH/USDT/USDC balance of a wallet address.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Load(false)
		r, err := redis.New(cfg.Redis)
		if err != nil {
			panic(err)
		}

		e, err := ethereum.New(cfg.Ethereum, r)
		if err != nil {
			panic(err)
		}
		balance, err := e.GetBalance(publicKey, currency)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Balance: %f%s\n", balance, currency)
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	BalanceCmd.Flags().StringVarP(&currency, "currency", "c", "TRX", "currency TRX/USDT/USDC")
	BalanceCmd.Flags().StringVarP(&publicKey, "wallet", "w", "", "wallet address")
	err := BalanceCmd.MarkFlagRequired("wallet")
	if err != nil {
		panic(err)
	}
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tronTrc20EventsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
