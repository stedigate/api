package solana

import (
	"context"
	"github.com/gagliardetto/solana-go/rpc"
)

func (s *Solana) GetCurrentBlock() (uint64, error) {
	recent, err := s.rpc.GetSlot(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return 0, err
	}

	return recent, nil
}
