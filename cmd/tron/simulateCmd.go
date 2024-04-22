package tron

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stedigate/core/internal/config"
	"github.com/stedigate/core/pkg/blockchains/tron"
	"github.com/stedigate/core/pkg/logger"
	"github.com/stedigate/core/pkg/redis"
	"log/slog"
	"os"
)

// SimulateCmd represents the tronTrc20Events command
var SimulateCmd = &cobra.Command{
	Use:   "simulate",
	Short: "Sent TRX/USDT/USDC of a wallet address to another wallet address",
	Long:  `Sent TRX/USDT/USDC of a wallet address to another wallet address.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Load(false)
		r, err := redis.New(cfg.Redis)
		if err != nil {
			panic(err)
		}
		l := logger.NewLogger(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: false}))
		t, err := tron.New(cfg.Tron, r, l)
		if err != nil {
			panic(err)
		}
		fee, err := t.SimulateSend(srcPrivateKey, destPublicKey, currency, amount)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Fee: %f%s\n", fee, "TRX")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	SimulateCmd.Flags().StringVarP(&currency, "currency", "c", "", "currency TRX/USDT/USDC")
	SimulateCmd.Flags().StringVarP(&srcPrivateKey, "source", "s", "", "source wallet address")
	SimulateCmd.Flags().StringVarP(&destPublicKey, "destination", "d", "", "destination wallet address")
	SimulateCmd.Flags().Uint64VarP(&amount, "amount", "a", 0, "amount to transfer")
	err := SimulateCmd.MarkFlagRequired("source")
	if err != nil {
		panic(err)
	}
	err = SimulateCmd.MarkFlagRequired("destination")
	if err != nil {
		panic(err)
	}
	err = SimulateCmd.MarkFlagRequired("amount")
	if err != nil {
		panic(err)
	}
	err = SimulateCmd.MarkFlagRequired("currency")
	if err != nil {
		panic(err)
	}
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tronTrc20EventsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
