package solana

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/rueidis"
	"github.com/spf13/cobra"
	"github.com/stedigate/core/internal/config"
	"github.com/stedigate/core/pkg/redis"
	"github.com/stedigate/core/pkg/solana"
)

// GetTransactionCmd represents the tronTrc20Events command
var GetTransactionCmd = &cobra.Command{
	Use:   "track",
	Short: "Fetches TRC20 events from the Tron blockchain.",
	Long:  `Fetches TRC20 events from the Tron blockchain. It can be used to fetch all events from a specific contract address.`,
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

	t, err := solana.New(cfg.Solana, r)
	if err != nil {
		panic(err)
	}

	cmd := r.B().Get().Key("tron:events:trc20:last_transaction_id").Build()
	lastScannedTrxID, err := r.Do(context.Background(), cmd).ToString()
	if err != nil {
		if !errors.Is(err, rueidis.Nil) {
			panic(err)
		}
	}
	events, err := t.GetContractEvents(lastScannedTrxID)
	if err != nil {
		panic(err)
	}

	if len(events) != 0 {
		latestTrxID := events[0].TransactionID
		cmd = r.B().Set().Key("tron:events:trc20:last_transaction_id").Value(latestTrxID).Build()
		r.Do(context.Background(), cmd)
	}

	fmt.Printf("%v\n", events)
}
