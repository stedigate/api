package solana

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/redis/rueidis"
	"golang.org/x/time/rate"
	"time"
)

type Solana struct {
	redis  rueidis.Client
	config *Config
	rpc    *rpc.Client
	ws     *ws.Client
}

type UsdtTransaction struct {
	ID string
}

func New(cfg *Config, r rueidis.Client) (*Solana, error) {
	rpcClient := rpc.NewWithCustomRPCClient(rpc.NewWithLimiter(
		rpc.MainNetBeta_RPC,
		rate.Every(time.Second), // time frame
		5,                       // limit of requests per time frame
	))

	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	if err != nil {
		panic(err)
	}

	t := &Solana{
		redis:  r,
		config: cfg,
		rpc:    rpcClient,
		ws:     wsClient,
	}

	return t, nil
}

func (s *Solana) GetContractEvents(lastScannedTrxID string) ([]TransferEvent, error) {
	go func() {
		err := s.subscribeToProgram(s.config.USDCTokenAddress)
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		err := s.subscribeToProgram(s.config.USDTTokenAddress)
		if err != nil {
			panic(err)
		}
	}()

	return nil, nil
}

func (s *Solana) subscribeToProgram(program string) error {
	usdcPubKey := solana.MustPublicKeyFromBase58(program)
	usdc, err := s.ws.ProgramSubscribeWithOpts(usdcPubKey, rpc.CommitmentConfirmed, solana.EncodingBase64Zstd, nil)
	if err != nil {
		return fmt.Errorf("unable to get token account: %w", err)
	}
	defer usdc.Unsubscribe()
	for {
		got, err := usdc.Recv()
		if err != nil {
			return fmt.Errorf("unable to get token account: %w", err)
		}
		spew.Dump(got)
	}
}

func (s *Solana) GetTransactionInfo(txId string) ([]map[string]interface{}, error) {

	return nil, nil
}

func (s *Solana) GetTrxBalance(address string) ([]map[string]interface{}, error) {

	return nil, nil
}

func (s *Solana) GetContractBalance(address string) ([]map[string]interface{}, error) {

	return nil, nil
}

func (s *Solana) GetTransactions(address string) ([]map[string]interface{}, error) {

	return nil, nil
}

func (s *Solana) GetContractTransactions(address string) ([]map[string]interface{}, error) {

	return nil, nil
}

func (s *Solana) SendTrx(src Wallet, dest string, amount float64) ([]map[string]interface{}, error) {

	return nil, nil
}

func (s *Solana) SendUsdt(src Wallet, dest string, amount float64) ([]map[string]interface{}, error) {

	return nil, nil
}
