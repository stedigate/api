package solana

import (
	"github.com/spf13/cobra"
	"github.com/stedigate/core/internal/config"
	"github.com/stedigate/core/pkg/blockchains/solana"
	"github.com/stedigate/core/pkg/redis"
	"log"
	"os"
	"os/signal"
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

	s, err := solana.New(cfg.Solana, r)
	if err != nil {
		panic(err)
	}

	s.GetContractEvents()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Perform any additional cleanup or logging as needed
	log.Println("Tracking solana shutting down...")
}
