package tron

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stedigate/core/internal/config"
	"github.com/stedigate/core/pkg/redis"
	"github.com/stedigate/core/pkg/tron"
)

// TransferCmd represents the tronTrc20Events command
var TransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Sent TRX/USDT/USDC of a wallet address to another wallet address",
	Long:  `Sent TRX/USDT/USDC of a wallet address to another wallet address.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Load(false)
		r, err := redis.New(cfg.Redis)
		if err != nil {
			panic(err)
		}

		t, err := tron.New(cfg.Tron, r)
		if err != nil {
			panic(err)
		}
		txId, err := t.Send(srcPrivateKey, destPublicKey, currency, amount)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Transaction: %s\n", txId)
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	TransferCmd.Flags().StringVarP(&currency, "currency", "c", "", "currency TRX/USDT/USDC")
	TransferCmd.Flags().StringVarP(&srcPrivateKey, "source", "s", "", "source wallet address")
	TransferCmd.Flags().StringVarP(&destPublicKey, "destination", "d", "", "destination wallet address")
	TransferCmd.Flags().Uint64VarP(&amount, "amount", "a", 0, "amount to transfer")
	err := TransferCmd.MarkFlagRequired("source")
	if err != nil {
		panic(err)
	}
	err = TransferCmd.MarkFlagRequired("destination")
	if err != nil {
		panic(err)
	}
	err = TransferCmd.MarkFlagRequired("amount")
	if err != nil {
		panic(err)
	}
	err = TransferCmd.MarkFlagRequired("currency")
	if err != nil {
		panic(err)
	}
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tronTrc20EventsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}