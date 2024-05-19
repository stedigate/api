package solana

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/redis/go-redis/v9"
	"github.com/stedigate/core/pkg/blockchains"
	"golang.org/x/time/rate"
	"log"
	"log/slog"
	"math/big"
	"strconv"
	"time"
)

type Solana struct {
	log    *slog.Logger
	redis  *redis.Client
	config *Config
	rpc    *rpc.Client
	ws     *ws.Client
}

type UsdtTransaction struct {
	ID string
}

func New(cfg *Config, r *redis.Client, l *slog.Logger) (*Solana, error) {
	rpcClient := rpc.NewWithCustomRPCClient(rpc.NewWithLimiter(cfg.RpcUrl, rate.Every(time.Second), 5))
	wsClient, err := ws.Connect(context.Background(), cfg.WssUrl)
	if err != nil {
		panic(err)
	}

	s := &Solana{
		log:    l,
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

func (s *Solana) GetContractTransactions(address string, lastScannedSignature string) ([]map[string]interface{}, error) {
	for {
		trxs, err := s.rpc.GetSignaturesForAddressWithOpts(
			context.Background(),
			solana.MustPublicKeyFromBase58(address),
			&rpc.GetSignaturesForAddressOpts{
				Until:      solana.MustSignatureFromBase58(lastScannedSignature),
				Commitment: rpc.CommitmentFinalized,
			},
		)
		if err != nil {
			return nil, err
		}

		var version uint64 = 0
		for _, tx := range trxs {
			tr, err := s.rpc.GetTransaction(
				context.Background(),
				tx.Signature,
				&rpc.GetTransactionOpts{
					Commitment:                     rpc.CommitmentFinalized,
					MaxSupportedTransactionVersion: &version,
				},
			)
			if err != nil {
				return nil, err
			}

			s.log.Warn("detect transaction destination address")
			spew.Dump(tr)
			preAmount := 0.0
			if len(tr.Meta.PreTokenBalances) >= 2 {
				preAmount, err = strconv.ParseFloat(tr.Meta.PreTokenBalances[1].UiTokenAmount.Amount, 64)
				if err != nil {
					return nil, err
				}
			}

			postAmount := 0.0
			if len(tr.Meta.PostTokenBalances) >= 2 {
				postAmount, err = strconv.ParseFloat(tr.Meta.PostTokenBalances[1].UiTokenAmount.Amount, 64)
				if err != nil {
					return nil, err
				}
			}

			amount := (postAmount - preAmount) / float64(tr.Meta.PostTokenBalances[1].UiTokenAmount.Decimals)
			te := Transaction{
				TxID:            tx.Signature.String(),
				BlockNumber:     tx.Slot,
				From:            Wallet{AddressBase58: tr.Meta.PreTokenBalances[0].Owner.String()},
				To:              Wallet{AddressBase58: tr.Meta.PreTokenBalances[1].Owner.String()},
				Amount:          *big.NewFloat(amount),
				Blockchain:      blockchains.Blockchain("Solana"),
				Status:          blockchains.TransactionStatus("confirmed"),
				Timestamp:       tr.BlockTime.Time(),
				FeeLimit:        tr.Meta.Fee / 10e9,
				ContractAddress: lastScannedSignature,
				Symbol:          blockchains.TokenSymbol("USDC"),
			}
			spew.Dump(te)
		}

		return nil, nil
	}
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
