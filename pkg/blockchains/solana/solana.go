package solana

import (
	"context"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/redis/go-redis/v9"
	"golang.org/x/time/rate"
	"log"
	"math/big"
	"time"
)

type Solana struct {
	redis  *redis.Client
	config *Config
	rpc    *rpc.Client
	ws     *ws.Client
}

type UsdtTransaction struct {
	ID string
}

func New(cfg *Config, r *redis.Client) (*Solana, error) {
	rpcClient := rpc.NewWithCustomRPCClient(rpc.NewWithLimiter(cfg.RpcUrl, rate.Every(time.Second), 5))
	wsClient, err := ws.Connect(context.Background(), cfg.WssUrl)
	if err != nil {
		panic(err)
	}

	s := &Solana{
		redis:  r,
		config: cfg,
		rpc:    rpcClient,
		ws:     wsClient,
	}

	return s, nil
}

func (s *Solana) GetTransactionInfo(txId string) ([]map[string]interface{}, error) {

	return nil, nil
}

func (s *Solana) GetTransactions(address string) ([]map[string]interface{}, error) {

	return nil, nil
}

func (s *Solana) GetContractTransactions(address string) ([]map[string]interface{}, error) {

	return nil, nil
}

func (s *Solana) Send(from, to, currency string, amount uint64) (string, error) {
	src, err := solana.PrivateKeyFromBase58(from)
	if err != nil {
		log.Fatal(err)
	}
	dest := solana.MustPublicKeyFromBase58(to)

	switch currency {
	case "SOL":
		return s.sendSol(src, dest, amount)
	case "USDT":
		t := solana.MustPublicKeyFromBase58(s.config.USDTAddress)
		return s.sendToken(src, dest, t, amount)
	case "USDC":
		t := solana.MustPublicKeyFromBase58(s.config.USDCAddress)
		return s.sendToken(src, dest, t, amount)
	default:
		return "", fmt.Errorf("unsupported currency")
	}
}

func (s *Solana) GetBalance(address string, currency string) (*big.Float, error) {
	switch currency {
	case "SOL":
		return s.getSolBalance(address)
	case "USDT":
		return s.getContractBalance(address, s.config.USDTAddress)
	case "USDC":
		return s.getContractBalance(address, s.config.USDCAddress)
	default:
		return nil, fmt.Errorf("unsupported currency")
	}
}
