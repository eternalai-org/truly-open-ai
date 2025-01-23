package utils_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/pkg/utils"
)

func TestGetBalanceOnSolanaChain(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		address string
		want    *big.Int
		wantErr bool
	}{
		{name: "case 1", address: "4cvnq5d8ctv9zjqkc1rf4j1tdxxc7edt6sm1ocu8qpmc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := utils.GetBalanceOnSolanaChain(context.Background(), tt.address)
			spew.Dump(got)
			spew.Dump(gotErr)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetBalanceOnSolanaChain() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetBalanceOnSolanaChain() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetBalanceOnSolanaChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
