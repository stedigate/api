package tron

import (
	"context"
	"errors"
	"fmt"
	"github.com/pushgate/core/internal/config"
	"github.com/pushgate/core/pkg/redis"
	"github.com/pushgate/core/pkg/tron"
	"github.com/redis/rueidis"
)

// Run Event holds the configuration for the event service.
func Run() {
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

	fmt.Println("%v", events)
}
