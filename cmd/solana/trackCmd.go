package solana

import (
	"context"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	goredis "github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"github.com/stedigate/core/internal/config"
	"github.com/stedigate/core/pkg/blockchains/solana"
	"github.com/stedigate/core/pkg/logger"
	"github.com/stedigate/core/pkg/redis"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// TrackContractCmd represents the tronTrc20Events command
var TrackContractCmd = &cobra.Command{
	Use:   "track",
	Short: "Fetches USDT/C events from the Solana blockchain.",
	Long:  `Fetches USDT/C events from the Solana blockchain. It can be used to fetch all events from a specific contract address.`,
	Run: func(cmd *cobra.Command, args []string) {
		getTransaction()
	},
}

func init() {

	TrackContractCmd.Flags().Uint64VarP(&block, "block", "b", 1, "start block number")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tronTrc20EventsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tronTrc20EventsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// getTransaction Event holds the configuration for the event service.
func getTransaction() {
	cfg := config.Load(false)

	r, err := redis.New(cfg.Redis)
	if err != nil {
		panic(err)
	}

	l := logger.NewLogger(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: false}))

	s, err := solana.New(cfg.Solana, r, l)
	if err != nil {
		panic(err)
	}

	l.Info("Starting Solana network scanner")

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	cacheKey := "solana:contracts:last_scanned_block"

	r.SAdd(ctx, "solana:wallets", "GXgjFg6pDRHPSrJY4Y28mAKRirvxhD1NDg5pe7roF9CL")
	r.SAdd(ctx, "solana:wallets", "6jmPKEMVd3yJWCbpQf1iLgGBVKSKZaX4J6CBqXPTyRvQ")

	var lastSlot uint64
	if block == 1 {
		lastSlot, err = r.Get(ctx, cacheKey).Uint64()
		switch {
		case errors.Is(err, goredis.Nil):
			fmt.Println("key does not exist")
		case err != nil:
			fmt.Println("Get failed", err)
		}
	} else {
		lastSlot = block
	}

	if lastSlot == 0 {
		lastSlot, err = s.GetCurrentBlock()
		fmt.Println("current block", lastSlot)
		if err != nil {
			panic(err)
		}
		r.Set(ctx, cacheKey, lastSlot, 0)
	}
	l.Info("Latest scanned block", slog.Uint64("block_number", lastSlot))
	go func() {
		for {
			events, err := s.GetContractTransactions(cfg.Solana.USDCAddress, "2zuyjyTVbSaCRjSnpYupn2BZQ5r7WG8NCftLX3ThoZMQfk2XgjAk2wwMP59Pq7dKZpw8ruyzBu6pQ3YvNa2A1MQK")
			if err != nil {
				panic(err)
			}

			spew.Dump(events)
			lastSlot, err = r.Get(ctx, cacheKey).Uint64()
			if err != nil {
				panic(err)
			}
			l.Info("scanning blocks finished", slog.Uint64("last_scanned_block", lastSlot))
			lastSlot++

			<-time.After(5 * time.Second)
		}
	}()

	<-c
	cancel()
}
