package tron

import (
	"context"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	goredis "github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"github.com/stedigate/core/internal/config"
	"github.com/stedigate/core/pkg/blockchains/tron"
	"github.com/stedigate/core/pkg/logger"
	"github.com/stedigate/core/pkg/redis"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// TrackTrc20Cmd represents the tronTrc20Events command
var TrackTrc20Cmd = &cobra.Command{
	Use:   "track",
	Short: "Fetches TRC20 events from the Tron blockchain.",
	Long:  `Fetches TRC20 events from the Tron blockchain. It can be used to fetch all events from a specific contract address.`,
	Run: func(cmd *cobra.Command, args []string) {
		getEvents()
	},
}

func init() {

	TrackTrc20Cmd.Flags().Int64VarP(&block, "block", "b", 1, "start block number")

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

	l := logger.NewLogger(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: false}))

	t, err := tron.New(cfg.Tron, r, l)
	if err != nil {
		panic(err)
	}

	l.Info("Starting Tron network scanner")

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	cacheKey := "tron:events:trc20:last_scanned_block"

	r.SAdd(ctx, "tron:wallets", "0xadc12d51725df5e2efaa5f6ed86f9e65735e0b42")
	r.SAdd(ctx, "tron:wallets", "0xa7cfa5539d7bd9798cf88d555dbed75b865b863c")

	var lastScannedBlock int64
	if block == 1 {
		lastScannedBlock, err = r.Get(ctx, cacheKey).Int64()
		switch {
		case errors.Is(err, goredis.Nil):
			fmt.Println("key does not exist")
		case err != nil:
			fmt.Println("Get failed", err)
		case lastScannedBlock == 0:
			lastScannedBlock, err = t.GetCurrentBlock()
			fmt.Println("current block", lastScannedBlock)
			if err != nil {
				panic(err)
			}
			r.Set(ctx, cacheKey, lastScannedBlock, 0)
		}
	} else {
		lastScannedBlock = block
	}

	l.Info("Latest scanned block", slog.Int64("block_number", lastScannedBlock))
	go func() {
		for {
			events, err := t.GetLatestTokenTransactions(cfg.Tron.USDTAddress, lastScannedBlock)
			if err != nil {
				panic(err)
			}

			spew.Dump(events)
			lastScannedBlock, err = r.Get(ctx, cacheKey).Int64()
			if err != nil {
				panic(err)
			}
			l.Info("scanning blocks finished", slog.Int64("last_scanned_block", lastScannedBlock))
			lastScannedBlock++

			<-time.After(5 * time.Second)
		}
	}()

	<-c
	cancel()
}
