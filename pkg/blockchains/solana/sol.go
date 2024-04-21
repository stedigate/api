package solana

import (
	"context"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/text"
	"math/big"
	"os"
)

func (s *Solana) sendSol(src solana.PrivateKey, dest solana.PublicKey, amount uint64) (string, error) {
	recent, err := s.rpc.GetRecentBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return "", fmt.Errorf("unable to get recent blockhash: %w", err)
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(amount*solana.LAMPORTS_PER_SOL, src.PublicKey(), dest).Build(),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(src.PublicKey()),
	)
	if err != nil {
		return "", fmt.Errorf("unable to create transaction: %w", err)
	}

	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if src.PublicKey().Equals(key) {
				return &src
			}
			return nil
		},
	)
	if err != nil {
		return "", fmt.Errorf("unable to sign transaction: %w", err)
	}

	// Pretty print the transaction:
	_, err = tx.EncodeTree(text.NewTreeEncoder(os.Stdout, "Transfer SOL"))
	if err != nil {
		return "", err
	}

	// Send transaction, and wait for confirmation:
	sig, err := confirm.SendAndConfirmTransaction(context.Background(), s.rpc, s.ws, tx)
	if err != nil {
		return "", err
	}

	return sig.String(), nil
}

func (s *Solana) getSolBalance(address string) (*big.Float, error) {
	var addr solana.PublicKey
	err := addr.Set(address)
	if err != nil {
		return new(big.Float), err
	}

	balance, err := s.rpc.GetBalance(context.Background(), addr, rpc.CommitmentFinalized)
	if err != nil {
		return new(big.Float), err
	}
	var lamports = new(big.Float).SetUint64(uint64(balance.Value))
	var solBalance = new(big.Float).Quo(lamports, new(big.Float).SetUint64(solana.LAMPORTS_PER_SOL))
	return solBalance, nil
}
