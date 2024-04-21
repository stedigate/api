package tron

import (
	"context"
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/redis/rueidis"
	"github.com/spf13/cobra"
	"github.com/stedigate/core/internal/config"
	"github.com/stedigate/core/pkg/blockchains/tron"
	"github.com/stedigate/core/pkg/redis"
)

// Trc20EventsCmd represents the tronTrc20Events command
var Trc20EventsCmd = &cobra.Command{
	Use:   "track",
	Short: "Fetches TRC20 events from the Tron blockchain.",
	Long:  `Fetches TRC20 events from the Tron blockchain. It can be used to fetch all events from a specific contract address.`,
	Run: func(cmd *cobra.Command, args []string) {
		getEvents()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tronTrc20EventsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tronTrc20EventsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Run Event holds the configuration for the event service.
func getEvents() {
	cfg := config.Load(false)

	r, err := redis.New(cfg.Redis)
	if err != nil {
		panic(err)
	}

	t, err := tron.New(cfg.Tron, r)
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

	spew.Dump(events)
}
