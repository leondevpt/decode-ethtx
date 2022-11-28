package pkg

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/tj/assert"
)

func TestDecodeTx(t *testing.T) {
	hexTx := "f8b4088502540be4008307a1209451c202fb78b3ae2bed8be41155aa37dc730b1d82884563918244f40000b8442e59905400000000000000000000000089ed2f18baaca8e178d8a872e69f27868f17565c2757218faa9c12a2940cc8ab856661c26ac8d47d93a1d672961a49bdafc854b5821792a0236c9aa9b59c698131e69df9619d3e7cb2af48d695b3ee3ffd3274a02f131e9da01ce2f63821d9035b8e102b9c5816109ab4c87fb941f712ba59e0c0f441cb28fc"

	r, _ := new(big.Int).SetString("0x236c9aa9b59c698131e69df9619d3e7cb2af48d695b3ee3ffd3274a02f131e9d", 0)
	s, _ := new(big.Int).SetString("0x1ce2f63821d9035b8e102b9c5816109ab4c87fb941f712ba59e0c0f441cb28fc", 0)
	v, _ := new(big.Int).SetString("1792", 0)

	type args struct {
		hexTx string
	}
	tests := []struct {
		name    string
		args    args
		want    *EthTx
		wantErr bool
	}{
		{name: "test1", args: args{hexTx: hexTx}, want: &EthTx{
			Nonce:    uint64(8),
			ChainId:  big.NewInt(2999),
			GasLimit: uint64(500000),
			GasPrice: big.NewInt(10000000000),
			Value:    big.NewInt(5000000000000000000),
			Data:     common.FromHex("2e59905400000000000000000000000089ed2f18baaca8e178d8a872e69f27868f17565c2757218faa9c12a2940cc8ab856661c26ac8d47d93a1d672961a49bdafc854b5"),
			From:     "0x89ed2f18BaAca8E178d8a872e69F27868f17565c",
			To:       "0x51c202Fb78B3AE2BEd8be41155Aa37dC730B1d82",
			R:        r,
			V:        v,
			S:        s,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeTx(tt.args.hexTx)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeTx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintDecodeTx(t *testing.T) {
	hexTx := "f8b4088502540be4008307a1209451c202fb78b3ae2bed8be41155aa37dc730b1d82884563918244f40000b8442e59905400000000000000000000000089ed2f18baaca8e178d8a872e69f27868f17565c2757218faa9c12a2940cc8ab856661c26ac8d47d93a1d672961a49bdafc854b5821792a0236c9aa9b59c698131e69df9619d3e7cb2af48d695b3ee3ffd3274a02f131e9da01ce2f63821d9035b8e102b9c5816109ab4c87fb941f712ba59e0c0f441cb28fc"

	rawTx, err := hex.DecodeString(hexTx)
	assert.NoError(t, err)
	var tx = new(types.Transaction)
	err = rlp.DecodeBytes(rawTx, tx)
	assert.NoError(t, err)
	msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), nil)
	assert.NoError(t, err)
	fmt.Printf("Hash: %s\n", tx.Hash().Hex())
	fmt.Printf("ChainId: %d\n", tx.ChainId())
	fmt.Printf("Value: %s\n", tx.Value().String())
	fmt.Printf("From: %s\n", msg.From().Hex()) // from field is not inside of transation
	fmt.Printf("To: %s\n", tx.To().Hex())
	fmt.Printf("Gas: %d\n", tx.Gas())
	fmt.Printf("Gas Price: %d\n", tx.GasPrice().Uint64())
	fmt.Printf("Nonce: %d\n", tx.Nonce())
	fmt.Printf("Transaction Data in hex: %s\n", hex.EncodeToString(tx.Data()))
	fmt.Print("\n")
}
