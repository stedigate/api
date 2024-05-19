package solana

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/text"
	"math/big"
	"os"
	"strconv"
)

func (s *Solana) GetContractEvents() {
	go func() {
		err := s.subscribeToProgram()
		if err != nil {
			panic(err)
		}
	}()
}

func (s *Solana) subscribeToProgram() error {
	// splToken := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	recent, err := s.rpc.GetRecentBlockhash(
		context.Background(),
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}

	lastSlot := uint64(recent.Context.Slot)
	blocks, err := s.rpc.GetBlocks(context.Background(), uint64(recent.Context.Slot-3), &lastSlot, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	for _, block := range blocks {
		go s.getBlock(block)

	}

	return nil
}

func (s *Solana) getBlock(slot uint64) {
	includeRewards := false
	maxSupportedTransactionVersion := uint64(0)
	b, err := s.rpc.GetBlockWithOpts(
		context.Background(),
		slot,
		&rpc.GetBlockOpts{
			Encoding:                       solana.EncodingBase64,
			Commitment:                     rpc.CommitmentFinalized,
			TransactionDetails:             rpc.TransactionDetailsFull,
			Rewards:                        &includeRewards,
			MaxSupportedTransactionVersion: &maxSupportedTransactionVersion,
		},
	)
	if err != nil {
		panic(err)
	}

	for _, tx := range b.Transactions {
		trx, err := tx.GetTransaction()
		if err != nil {
			panic(err)
		}

		isSplToken, isUsdc := false, false
		for _, acc := range trx.Message.AccountKeys {
			if acc.Equals(solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")) {
				isSplToken = true
			}
			if acc.Equals(solana.MustPublicKeyFromBase58("EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")) {
				isUsdc = true
			}
		}

		if isSplToken && isUsdc {
			var startAmount, endAmount *big.Float
			for _, pre := range tx.Meta.PreTokenBalances {
				if pre.Mint.Equals(solana.MustPublicKeyFromBase58("EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")) {
					amount, err := strconv.ParseFloat(pre.UiTokenAmount.Amount, 64)
					if err != nil {
						panic(err)
					}
					startAmount = new(big.Float).SetFloat64(amount / 1e6)
				}
			}

			for _, post := range tx.Meta.PostTokenBalances {
				if post.Mint.Equals(solana.MustPublicKeyFromBase58("EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")) {
					amount, err := strconv.ParseFloat(post.UiTokenAmount.Amount, 64)
					if err != nil {
						panic(err)
					}
					endAmount = new(big.Float).SetFloat64(amount / 1e6)
				}
			}

			fmt.Println("Start amount: ", startAmount)
			fmt.Println("End amount: ", endAmount)
			fmt.Println("Total amount: ", new(big.Float).Sub(endAmount, startAmount))
			spew.Dump(tx.Transaction)
			spew.Dump(trx)
		}

	}
}

func (s *Solana) sendToken(src solana.PrivateKey, dest, t solana.PublicKey, amount uint64) (string, error) {

	recent, err := s.rpc.GetRecentBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return "", fmt.Errorf("unable to get recent blockhash: %w", err)
	}
	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			token.NewTransferInstruction(amount*solana.LAMPORTS_PER_SOL, t, dest, src.PublicKey(), []solana.PublicKey{}).Build(),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(src.PublicKey()),
	)
	if err != nil {
		panic(err)
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

func (s *Solana) getContractBalance(address, contract string) (*big.Float, error) {
	addr := solana.MustPublicKeyFromBase58(address)
	contAddr := solana.MustPublicKeyFromBase58(contract)
	out, err := s.rpc.GetTokenAccountsByOwner(
		context.Background(),
		addr,
		&rpc.GetTokenAccountsConfig{
			Mint: contAddr.ToPointer(),
		},
		&rpc.GetTokenAccountsOpts{
			Commitment: rpc.CommitmentFinalized,
		},
	)
	if err != nil {
		panic(err)
	}

	if len(out.Value) == 0 {
		return new(big.Float), nil
	}

	var tokAcc token.Account
	data := out.Value[0].Account.Data.GetBinary()
	dec := bin.NewBinDecoder(data)
	err = dec.Decode(&tokAcc)
	if err != nil {
		panic(err)
	}

	var lamports = new(big.Float).SetUint64(tokAcc.Amount)
	var balance = new(big.Float).Quo(lamports, new(big.Float).SetInt(solana.DecimalsInBigInt(6)))
	return balance, nil
}
