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
	txSig := solana.MustSignatureFromBase58("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	sub, err := s.ws.SignatureSubscribe(
		txSig,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	for {
		got, err := sub.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Printf("got: %+v\n", got)
		// s.config.USDTProgramAddress
		spew.Dump(got)
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
