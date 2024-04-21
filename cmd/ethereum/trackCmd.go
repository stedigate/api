package ethereum

import (
	"context"
	"errors"
	"github.com/redis/rueidis"
	"github.com/spf13/cobra"
	"github.com/stedigate/core/internal/config"
	"github.com/stedigate/core/pkg/ethereum"
	"github.com/stedigate/core/pkg/redis"
	"log"
)

// GetTransactionCmd represents the tronTrc20Events command
var GetTransactionCmd = &cobra.Command{
	Use:   "track",
	Short: "Fetches USDT/C events from the Solana blockchain.",
	Long:  `Fetches USDT/C events from the Solana blockchain. It can be used to fetch all events from a specific contract address.`,
	Run: func(cmd *cobra.Command, args []string) {
		getTransaction()
	},
}

// getTransaction Event holds the configuration for the event service.
func getTransaction() {
	cfg := config.Load(false)

	r, err := redis.New(cfg.Redis)
	if err != nil {
		panic(err)
	}

	e, err := ethereum.New(cfg.Ethereum, r)
	if err != nil {
		panic(err)
	}

	cmd := r.B().Get().Key("ethereum:events:erc20:usdt:last_transaction_id").Build()
	lastScannedTrxID, err := r.Do(context.Background(), cmd).ToString()
	if err != nil {
		if !errors.Is(err, rueidis.Nil) {
			panic(err)
		}
	}
	events, err := e.GetUsdtContractEvents(lastScannedTrxID)
	if err != nil {
		panic(err)
	}

	if len(events) != 0 {
		latestTrxID := events[0].Hash
		cmd = r.B().Set().Key("ethereum:events:erc20:usdt:last_transaction_id").Value(latestTrxID).Build()
		r.Do(context.Background(), cmd)
	}
	/*
		cmd = r.B().Get().Key("ethereum:events:erc20:usdc:last_transaction_id").Build()
		lastScannedTrxID, err = r.Do(context.Background(), cmd).ToString()
		if err != nil {
			if !errors.Is(err, rueidis.Nil) {
				panic(err)
			}
		}
		lastScannedTrxID = "0x439946736b31544b20d93ecebca024353ab4910810c13f866c505f2b518fd4f9"
		events, err = e.GetUsdcContractEvents(lastScannedTrxID)
		if err != nil {
			panic(err)
		}

		if len(events) != 0 {
			latestTrxID := events[0].Hash
			cmd = r.B().Set().Key("ethereum:events:erc20:usdc:last_transaction_id").Value(latestTrxID).Build()
			r.Do(context.Background(), cmd)
		}*/

	// Perform any additional cleanup or logging as needed
	log.Println("Tracking solana shutting down...")
}
