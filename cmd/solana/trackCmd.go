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

// TrackCmd represents the tronTrc20Events command
var TrackCmd = &cobra.Command{
	Use:   "track",
	Short: "Fetches USDT/C events from the Solana blockchain.",
	Long:  `Fetches USDT/C events from the Solana blockchain. It can be used to fetch all events from a specific contract address.`,
	Run: func(cmd *cobra.Command, args []string) {
		getTransaction()
	},
}

func init() {

	TrackCmd.Flags().Int64VarP(&block, "block", "b", 1, "start block number")

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
